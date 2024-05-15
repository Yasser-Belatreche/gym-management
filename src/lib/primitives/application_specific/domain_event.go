package application_specific

import "time"
import "gym-management/src/lib/primitives/generic"

type DomainEvent[T interface{}] struct {
	EventId    string
	OccurredAt time.Time
	EventType  string
	Payload    T
}

func NewDomainEvent[T interface{}](eventType string, payload T) *DomainEvent[T] {
	return &DomainEvent[T]{
		EventId:    generic.GenerateULID(),
		OccurredAt: time.Now(),
		EventType:  eventType,
		Payload:    payload,
	}
}
