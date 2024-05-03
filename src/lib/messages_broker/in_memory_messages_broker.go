package messages_broker

import (
	"gym-management/src/lib/primitives/application_specific"
	"sync"
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
			err := handler(event, session)
			if err != nil {
				// Log error
			}

			routine <- err
		}(handler)
	}

	for range b.eventHandlers[event.EventType] {
		if err := <-routine; err != nil {
			return err
		}
	}

	return nil
}

func (b *InMemoryMessagesBroker) asyncPublish(event *application_specific.DomainEvent[interface{}], session *application_specific.Session) *application_specific.ApplicationException {
	if b.eventHandlers[event.EventType] == nil {
		return nil
	}

	wg := sync.WaitGroup{}

	for _, handler := range b.eventHandlers[event.EventType] {
		wg.Add(1)

		go func(handler func(event *application_specific.DomainEvent[interface{}], session *application_specific.Session) *application_specific.ApplicationException) {
			defer wg.Done()

			err := handler(event, session)
			if err != nil {
				// Log error
			}
		}(handler)
	}

	wg.Wait()

	return nil
}

func (b *InMemoryMessagesBroker) Subscribe(subscriber *Subscriber) {
	if b.eventHandlers[subscriber.event] == nil {
		b.eventHandlers[subscriber.event] = make([]func(event *application_specific.DomainEvent[interface{}], session *application_specific.Session) *application_specific.ApplicationException, 0)
	}

	b.eventHandlers[subscriber.event] = append(b.eventHandlers[subscriber.event], subscriber.handler)
}

func (b *InMemoryMessagesBroker) Ask(question *Question, session *application_specific.Session) (map[string]string, *application_specific.ApplicationException) {
	if b.answers[question.Question] == nil {
		panic("Answer not registered")
	}

	return b.answers[question.Question](question.Params, session)
}

func (b *InMemoryMessagesBroker) RegisterAnswer(answer *Answer) {
	if b.answers[answer.Question] != nil {
		panic("Answer already registered")
	}

	b.answers[answer.Question] = answer.Answer
}
