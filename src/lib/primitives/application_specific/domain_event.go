package application_specific

import "time"
import "gym-management/src/lib/primitives/generic"

type DomainEvent[T interface{}] struct {
	eventId    string
	occurredAt time.Time
	eventType  string
	payload    T
}

func NewDomainEvent[T any](eventType string, payload T) *DomainEvent[T] {
	return &DomainEvent[T]{
		eventId:    generic.GenerateUUID(),
		occurredAt: time.Now(),
		eventType:  eventType,
		payload:    payload,
	}
}
