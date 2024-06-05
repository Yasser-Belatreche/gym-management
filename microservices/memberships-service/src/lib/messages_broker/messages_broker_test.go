package messages_broker

import (
	"gym-management-memberships/src/lib/primitives/application_specific"
	"math"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

type testBroker struct {
	name   string
	broker MessagesBroker
}

func TestMessagesBroker(t *testing.T) {
	brokers := []testBroker{
		{
			name:   "AsyncEvents In Memory Broker",
			broker: NewInMemoryMessagesBroker(InMemoryMessagesBrokerConfig{AsyncEvents: true}),
		},
		{
			name:   "Sync In Memory Broker",
			broker: NewInMemoryMessagesBroker(InMemoryMessagesBrokerConfig{AsyncEvents: false}),
		},
		{
			name:   "RabbitMQ Memory Broker",
			broker: NewRabbitMQMessagesBroker(RabbitMQMessagesBrokerConfig{Uri: "amqp://user:password@localhost:5672", Durable: false}),
		},
	}

	for _, elem := range brokers {
		t.Run(elem.name, func(t *testing.T) {
			runTestsOn(elem.broker, t)
		})
	}
}

func runTestsOn(broker MessagesBroker, t *testing.T) {
	t.Run("Should publish Event to the subscriber", func(t *testing.T) {
		event, session, subscriber, countCalls := getEventsMocks("test", nil)

		broker.Subscribe(subscriber)

		if err := broker.Publish(event, session); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		time.Sleep(50 * time.Millisecond)

		if countCalls() != 1 {
			t.Errorf("Expected 1 call, got %v", countCalls())
		}
	})

	t.Run("Should not publish Event to the subscriber when he is not subscribed", func(t *testing.T) {
		event, session, _, countCalls := getEventsMocks("test", nil)

		if err := broker.Publish(event, session); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		time.Sleep(50 * time.Millisecond)

		if countCalls() != 0 {
			t.Errorf("Expected 0 calls, got %v", countCalls())
		}
	})

	t.Run("Should publish the event to all the subscribers", func(t *testing.T) {
		event, session, subscriber, countCalls := getEventsMocks("test", nil)
		_, _, subscriber2, countCalls2 := getEventsMocks("test", nil)

		broker.Subscribe(subscriber)
		broker.Subscribe(subscriber2)

		if err := broker.Publish(event, session); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		time.Sleep(50 * time.Millisecond)

		if countCalls() != 1 {
			t.Errorf("Expected 1 calls, got %v", countCalls())
		}

		if countCalls2() != 1 {
			t.Errorf("Expected 1 call, got %v", countCalls2())
		}
	})

	t.Run("Should Not call subscribers of different events", func(t *testing.T) {
		event, session, subscriber, countCalls := getEventsMocks("test", nil)
		_, _, subscriber2, countCalls2 := getEventsMocks("test2", nil)

		broker.Subscribe(subscriber)
		broker.Subscribe(subscriber2)

		if err := broker.Publish(event, session); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		time.Sleep(50 * time.Millisecond)

		if countCalls() != 1 {
			t.Errorf("Expected 1 call, got %v", countCalls())
		}

		if countCalls2() != 0 {
			t.Errorf("Expected 0 calls, got %v", countCalls2())
		}
	})

	t.Run("Should be able to register a reply and get it", func(t *testing.T) {
		reply, session := getReplyMocks("test", func(params map[string]string, session *application_specific.Session) (map[string]string, *application_specific.ApplicationException) {
			return map[string]string{"test": "test"}, nil
		})

		broker.RegisterReply(reply)

		response, err := broker.GetReply("test", nil, session)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if response["test"] != "test" {
			t.Errorf("Expected test, got %v", response["test"])
		}
	})

	t.Run("Should not be able to register two replies with the same message", func(t *testing.T) {
		reply, _ := getReplyMocks("test1", func(params map[string]string, session *application_specific.Session) (map[string]string, *application_specific.ApplicationException) {
			return map[string]string{"test": "test"}, nil
		})

		reply2, _ := getReplyMocks("test1", func(params map[string]string, session *application_specific.Session) (map[string]string, *application_specific.ApplicationException) {
			return map[string]string{"test": "test"}, nil
		})

		broker.RegisterReply(reply)

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected a panic, got none")
			}
		}()

		broker.RegisterReply(reply2)
	})

	t.Run("Should pass the params to the reply handler", func(t *testing.T) {
		reply, session := getReplyMocks("test2", func(params map[string]string, session *application_specific.Session) (map[string]string, *application_specific.ApplicationException) {
			if params["test"] != "test" {
				t.Errorf("Expected test, got %v", params["test"])
			}
			return map[string]string{"test": "test"}, nil
		})

		broker.RegisterReply(reply)

		response, err := broker.GetReply("test2", map[string]string{"test": "test"}, session)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if response["test"] != "test" {
			t.Errorf("Expected test, got %v", response["test"])
		}
	})
}

func getEventsMocks(eventName string, payload interface{}) (*application_specific.DomainEvent[interface{}], *application_specific.Session, *Subscriber, func() int) {
	event := application_specific.NewDomainEvent(eventName, payload)
	session := application_specific.NewSession()

	callCount := 0

	subscriber := &Subscriber{
		Event:     eventName,
		Component: "Test" + strconv.Itoa(rand.Intn(math.MaxInt32)), // to randomize the component name
		Handler: func(event *application_specific.DomainEvent[interface{}], session *application_specific.Session) *application_specific.ApplicationException {
			callCount++
			return nil
		},
	}

	subscriberCallCount := func() int {
		return callCount
	}

	return event, session, subscriber, subscriberCallCount
}

func getReplyMocks(message string, handler func(params map[string]string, session *application_specific.Session) (map[string]string, *application_specific.ApplicationException)) (*Reply, *application_specific.Session) {
	reply := &Reply{
		Message: message,
		Handler: handler,
	}
	session := application_specific.NewSession()

	return reply, session
}
