package messages_broker

import (
	"gym-management/src/lib/primitives/application_specific"
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
			name:   "Async In Memory Broker",
			broker: NewInMemoryMessagesBroker(InMemoryMessagesBrokerConfig{Async: true}),
		},
		{
			name:   "Sync In Memory Broker",
			broker: NewInMemoryMessagesBroker(InMemoryMessagesBrokerConfig{Async: false}),
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
		event, session, subscriber, countCalls := getMocks("test", nil)

		broker.Subscribe(subscriber)

		if err := broker.Publish(event, session); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		time.Sleep(50 * time.Millisecond)

		if countCalls() != 1 {
			t.Errorf("Expected 1 call, got %v", countCalls())
		}
	})

	t.Run("Should not publish Event to the subscriber", func(t *testing.T) {
		event, session, _, countCalls := getMocks("test", nil)

		if err := broker.Publish(event, session); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		time.Sleep(50 * time.Millisecond)

		if countCalls() != 0 {
			t.Errorf("Expected 0 calls, got %v", countCalls())
		}
	})

	t.Run("Should be call all subscribers of the same Event", func(t *testing.T) {
		event, session, subscriber, countCalls := getMocks("test", nil)
		_, _, subscriber2, countCalls2 := getMocks("test", nil)

		broker.Subscribe(subscriber)
		broker.Subscribe(subscriber2)

		if err := broker.Publish(event, session); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		time.Sleep(50 * time.Millisecond)

		if countCalls() != 1 {
			t.Errorf("Expected 0 calls, got %v", countCalls())
		}

		if countCalls2() != 1 {
			t.Errorf("Expected 1 call, got %v", countCalls2())
		}
	})

	t.Run("Should Not call subscribers of different events", func(t *testing.T) {
		event, session, subscriber, countCalls := getMocks("test", nil)
		_, _, subscriber2, countCalls2 := getMocks("test2", nil)

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
}

func getMocks(eventName string, payload interface{}) (*application_specific.DomainEvent[interface{}], *application_specific.Session, *Subscriber, func() int) {
	event := application_specific.NewDomainEvent(eventName, payload)
	session := application_specific.NewSession()

	callCount := 0

	subscriber := &Subscriber{
		Event: eventName,
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
