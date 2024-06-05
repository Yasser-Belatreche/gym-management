package messages_broker

import (
	"context"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"gym-management-gyms/src/lib/primitives/application_specific"
	"time"
)

type RabbitMQMessagesBrokerConfig struct {
	Uri     string
	Durable bool
}

type RabbitMQMessagesBroker struct {
	config RabbitMQMessagesBrokerConfig

	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewRabbitMQMessagesBroker(config RabbitMQMessagesBrokerConfig) *RabbitMQMessagesBroker {
	return &RabbitMQMessagesBroker{
		config: config,
	}
}

func (r *RabbitMQMessagesBroker) Publish(event *application_specific.DomainEvent[interface{}], session *application_specific.Session) *application_specific.ApplicationException {
	err := r.assertConnected()
	if err != nil {
		return err
	}

	exchangeName := "Events::" + event.EventType
	e := r.ch.ExchangeDeclare(exchangeName, "fanout", r.config.Durable, false, false, false, nil)
	if e != nil {
		return application_specific.NewUnknownException("RABBITMQ_EXCHANGE_DECLARE_ERROR", e.Error(), nil)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	jsonData, e := json.Marshal(event)
	if e != nil {
		return application_specific.NewUnknownException("JSON_MARSHAL_ERROR", e.Error(), nil)
	}

	var deliveryMode uint8
	if r.config.Durable {
		deliveryMode = amqp.Persistent
	} else {
		deliveryMode = amqp.Transient
	}

	e = r.ch.PublishWithContext(ctx, exchangeName, "", false, false, amqp.Publishing{
		ContentType:   "application/json",
		Body:          jsonData,
		DeliveryMode:  deliveryMode,
		CorrelationId: session.CorrelationId,
	})
	if e != nil {
		return application_specific.NewUnknownException("RABBITMQ_PUBLISH_ERROR", e.Error(), nil)
	}

	return nil
}

func (r *RabbitMQMessagesBroker) Subscribe(subscribers ...*Subscriber) {
	err := r.assertConnected()
	if err != nil {
		panic(err)
	}

	for _, subscriber := range subscribers {
		exchangeName := "Events::" + subscriber.Event
		err := r.ch.ExchangeDeclare(exchangeName, "fanout", r.config.Durable, false, false, false, nil)
		if err != nil {
			panic(application_specific.NewUnknownException("RABBITMQ_EXCHANGE_DECLARE_ERROR", err.Error(), nil))
		}

		// eg: "EventConsumer::GymsManager::Gyms.Disabled"
		queueName := "EventConsumer::" + subscriber.Component + "::" + subscriber.Event

		queue, err := r.ch.QueueDeclare(queueName, r.config.Durable, false, false, false, nil)
		if err != nil {
			panic(application_specific.NewUnknownException("RABBITMQ_QUEUE_DECLARE_ERROR", err.Error(), nil))
		}

		err = r.ch.QueueBind(queue.Name, "", exchangeName, false, nil)
		if err != nil {
			panic(application_specific.NewUnknownException("RABBITMQ_QUEUE_BIND_ERROR", err.Error(), nil))
		}

		messages, err := r.ch.Consume(queue.Name, "", false, false, false, false, nil)
		if err != nil {
			panic(application_specific.NewUnknownException("RABBITMQ_CONSUME_ERROR", err.Error(), nil))
		}

		go func() {
			for d := range messages {
				go func() {
					defer func() {
						if r := recover(); r != nil {
							//TODO log error
							d.Nack(false, false)
						}
					}()

					var event application_specific.DomainEvent[interface{}]
					err := json.Unmarshal(d.Body, &event)
					if err != nil {
						//TODO log error
						d.Nack(false, false)
						return
					}

					er := subscriber.Handler(&event, application_specific.NewSessionWithCorrelationId(d.CorrelationId))
					if er != nil {
						//TODO log error
						fmt.Println(er)

						d.Nack(false, false)
						return
					}

					d.Ack(false)
				}()
			}
		}()
	}
}

func (r *RabbitMQMessagesBroker) GetReply(message string, params map[string]string, session *application_specific.Session) (map[string]string, *application_specific.ApplicationException) {
	err := r.assertConnected()
	if err != nil {
		return nil, err
	}

	exchangeName := "Replies::" + message

	queue, er := r.ch.QueueDeclare("", false, false, true, false, nil)
	if er != nil {
		return nil, application_specific.NewUnknownException("RABBITMQ_QUEUE_DECLARE_ERROR", er.Error(), nil)
	}

	er = r.ch.QueueBind(queue.Name, queue.Name, exchangeName, false, nil)
	if er != nil {
		return nil, application_specific.NewUnknownException("RABBITMQ_QUEUE_BIND_ERROR", er.Error(), nil)
	}

	messages, er := r.ch.Consume(queue.Name, "", false, false, false, false, nil)
	if er != nil {
		return nil, application_specific.NewUnknownException("RABBITMQ_CONSUME_ERROR", er.Error(), nil)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	jsonData, er := json.Marshal(params)
	if er != nil {
		return nil, application_specific.NewUnknownException("JSON_MARSHAL_ERROR", er.Error(), nil)
	}

	er = r.ch.PublishWithContext(ctx, exchangeName, message, false, false, amqp.Publishing{
		ContentType:   "application/json",
		Body:          jsonData,
		DeliveryMode:  amqp.Persistent,
		ReplyTo:       queue.Name,
		CorrelationId: session.CorrelationId,
	})
	if er != nil {
		return nil, application_specific.NewUnknownException("RABBITMQ_PUBLISH_ERROR", er.Error(), nil)
	}

	var response map[string]string

	for d := range messages {
		if session.CorrelationId == d.CorrelationId {
			err := json.Unmarshal(d.Body, &response)
			if err != nil {
				d.Nack(false, false)

				return nil, application_specific.NewUnknownException("JSON_UNMARSHAL_ERROR", err.Error(), nil)
			}

			d.Ack(false)

			break
		}
	}

	if response == nil {
		return nil, application_specific.NewUnknownException("RABBITMQ_REPLY_NOT_FOUND", "Reply not found", nil)
	}

	return response, nil
}

func (r *RabbitMQMessagesBroker) RegisterReply(replies ...*Reply) {
	err := r.assertConnected()
	if err != nil {
		panic(err)
	}

	for _, reply := range replies {
		exchangeName := "Replies::" + reply.Message
		err := r.ch.ExchangeDeclare(exchangeName, "direct", r.config.Durable, false, false, false, nil)
		if err != nil {
			panic(application_specific.NewUnknownException("RABBITMQ_EXCHANGE_DECLARE_ERROR", err.Error(), nil))
		}

		queueName := "ReplyHandler::" + reply.Message
		queue, err := r.ch.QueueDeclare(queueName, r.config.Durable, false, false, false, nil)
		if err != nil {
			panic(application_specific.NewUnknownException("RABBITMQ_QUEUE_DECLARE_ERROR", err.Error(), nil))
		}

		if queue.Consumers > 0 {
			panic(application_specific.NewUnknownException("RABBITMQ_QUEUE_CONSUMERS_ERROR", "Queue already has consumers", nil))
		}

		err = r.ch.QueueBind(queue.Name, reply.Message, exchangeName, false, nil)
		if err != nil {
			panic(application_specific.NewUnknownException("RABBITMQ_QUEUE_BIND_ERROR", err.Error(), nil))
		}

		r.ch.Qos(0, 0, false)
		messages, err := r.ch.Consume(queue.Name, "", false, true, false, false, nil)
		if err != nil {
			panic(application_specific.NewUnknownException("RABBITMQ_CONSUME_ERROR", err.Error(), nil))
		}

		go func() {
			for d := range messages {
				go func() {
					defer func() {
						if r := recover(); r != nil {
							//TODO log error
							d.Nack(false, false)
						}
					}()

					var params map[string]string
					err := json.Unmarshal(d.Body, &params)
					if err != nil {
						//TODO log error
						d.Nack(false, false)
						return
					}

					response, er := reply.Handler(params, application_specific.NewSessionWithCorrelationId(d.CorrelationId))
					if er != nil {
						//TODO log error
						d.Nack(false, false)
						return
					}

					jsonData, err := json.Marshal(response)
					if err != nil {
						//TODO log error
						d.Nack(false, false)
						return
					}

					err = r.ch.Publish(exchangeName, d.ReplyTo, false, false, amqp.Publishing{
						ContentType:   "application/json",
						Body:          jsonData,
						DeliveryMode:  amqp.Persistent,
						CorrelationId: d.CorrelationId,
					})
					if err != nil {
						//TODO log error
						d.Nack(false, false)
						return
					}

					d.Ack(false)
				}()
			}
		}()
	}
}

func (r *RabbitMQMessagesBroker) HealthCheck() *Health {
	err := r.assertConnected()
	if err != nil {
		return &Health{
			Provider: "RabbitMQ",
			Status:   "DOWN",
			Message:  err.Message,
		}
	}

	return &Health{
		Provider: "RabbitMQ",
		Status:   "UP",
		Message:  "Connection is active",
	}
}

func (r *RabbitMQMessagesBroker) assertConnected() *application_specific.ApplicationException {
	if r.conn == nil {
		err := r.connect()
		if err != nil {
			return err
		}

		return nil
	}

	if r.conn.IsClosed() {
		err := r.connect()
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (r *RabbitMQMessagesBroker) connect() *application_specific.ApplicationException {
	conn, err := amqp.Dial(r.config.Uri)
	if err != nil {
		return application_specific.NewUnknownException("RABBITMQ_CONNECTION_ERROR", err.Error(), nil)
	}

	r.conn = conn

	ch, err := conn.Channel()
	if err != nil {
		return application_specific.NewUnknownException("RABBITMQ_CHANNEL_ERROR", err.Error(), nil)
	}

	r.ch = ch
	err = r.ch.Qos(1, 0, true)
	if err != nil {
		return application_specific.NewUnknownException("RABBITMQ_QOS_ERROR", err.Error(), nil)
	}

	return nil
}

//"Replies::Emails.IsUsed"
