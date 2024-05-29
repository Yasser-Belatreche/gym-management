package messages_broker

import (
	"gym-management/src/lib/primitives/application_specific"
)

type InMemoryMessagesBrokerConfig struct {
	Async bool
}

type InMemoryMessagesBroker struct {
	config        InMemoryMessagesBrokerConfig
	eventHandlers map[string][]func(event *application_specific.DomainEvent[interface{}], session *application_specific.Session) *application_specific.ApplicationException
	answers       map[string]func(params map[string]string, session *application_specific.Session) (map[string]string, *application_specific.ApplicationException)
}

func NewInMemoryMessagesBroker(config InMemoryMessagesBrokerConfig) *InMemoryMessagesBroker {
	return &InMemoryMessagesBroker{
		config:        config,
		eventHandlers: make(map[string][]func(event *application_specific.DomainEvent[interface{}], session *application_specific.Session) *application_specific.ApplicationException),
		answers:       make(map[string]func(params map[string]string, session *application_specific.Session) (map[string]string, *application_specific.ApplicationException)),
	}
}

func (b *InMemoryMessagesBroker) Publish(event *application_specific.DomainEvent[interface{}], session *application_specific.Session) *application_specific.ApplicationException {
	if b.config.Async {
		return b.asyncPublish(event, session)
	}

	return b.syncPublish(event, session)
}

func (b *InMemoryMessagesBroker) syncPublish(event *application_specific.DomainEvent[interface{}], session *application_specific.Session) *application_specific.ApplicationException {
	if b.eventHandlers[event.EventType] == nil {
		return nil
	}

	routine := make(chan *application_specific.ApplicationException)

	for _, handler := range b.eventHandlers[event.EventType] {
		go func(handler func(event *application_specific.DomainEvent[interface{}], session *application_specific.Session) *application_specific.ApplicationException) {
			defer func() {
				if r := recover(); r != nil {
					switch v := r.(type) {
					case error:
						err := application_specific.NewUnknownException("EVENT_HANDLER_PANIC", v.Error(), nil)
						routine <- err
					case string:
						err := application_specific.NewUnknownException("EVENT_HANDLER_PANIC", v, nil)
						routine <- err
					default:
						err := application_specific.NewUnknownException("EVENT_HANDLER_PANIC", "panic", nil)
						routine <- err
					}
				}
			}()

			err := handler(event, session)
			if err != nil {
				// Log error
			}

			routine <- err
		}(handler)
	}

	var lastErr *application_specific.ApplicationException = nil
	for range b.eventHandlers[event.EventType] {
		if err := <-routine; err != nil {
			lastErr = err
		}
	}

	return lastErr
}

func (b *InMemoryMessagesBroker) asyncPublish(event *application_specific.DomainEvent[interface{}], session *application_specific.Session) *application_specific.ApplicationException {
	if b.eventHandlers[event.EventType] == nil {
		return nil
	}

	for _, handler := range b.eventHandlers[event.EventType] {
		go func(handler func(event *application_specific.DomainEvent[interface{}], session *application_specific.Session) *application_specific.ApplicationException) {
			defer func() {
				if r := recover(); r != nil {
					// Log the panic
				}
			}()

			err := handler(event, session)
			if err != nil {
				// Log error
			}
		}(handler)
	}

	return nil
}

func (b *InMemoryMessagesBroker) Subscribe(subscribers ...*Subscriber) {
	for _, subscriber := range subscribers {
		if b.eventHandlers[subscriber.Event] == nil {
			b.eventHandlers[subscriber.Event] = make([]func(event *application_specific.DomainEvent[interface{}], session *application_specific.Session) *application_specific.ApplicationException, 0)
		}

		b.eventHandlers[subscriber.Event] = append(b.eventHandlers[subscriber.Event], subscriber.Handler)
	}
}

func (b *InMemoryMessagesBroker) Ask(question string, params map[string]string, session *application_specific.Session) (map[string]string, *application_specific.ApplicationException) {
	if b.answers[question] == nil {
		panic("Answer not registered")
	}

	return b.answers[question](params, session)
}

func (b *InMemoryMessagesBroker) RegisterAnswer(answers ...*Answer) {
	for _, answer := range answers {
		if b.answers[answer.Question] != nil {
			panic("Answer already registered")
		}

		b.answers[answer.Question] = answer.Answer
	}
}
