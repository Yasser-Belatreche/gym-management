package events

import "gym-management/src/lib/primitives/application_specific"

type GymEvent[T interface{}] struct {
	*application_specific.DomainEvent[T]
}

func NewGymEvent(eventType string, payload interface{}) *GymEvent[interface{}] {
	return &GymEvent[interface{}]{
		DomainEvent: application_specific.NewDomainEvent(eventType, payload),
	}
}
