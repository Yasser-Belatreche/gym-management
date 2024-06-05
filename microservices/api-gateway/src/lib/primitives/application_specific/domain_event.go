package application_specific

import (
	"github.com/mitchellh/mapstructure"
	"time"
)
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

func ParsePayload[T interface{}](event *DomainEvent[interface{}]) (*T, *ApplicationException) {
	switch payload := event.Payload.(type) {
	case T:
		return &payload, nil
	case *T:
		return payload, nil
	case map[string]interface{}:
		var decodedPayload T

		err := mapstructure.Decode(payload, &decodedPayload)
		if err != nil {
			return nil, NewDeveloperException("INVALID_EVENT_PAYLOAD_TYPE", event.EventType+" payload is not as expected")
		}

		return &decodedPayload, nil
	default:
		return nil, NewDeveloperException("INVALID_EVENT_PAYLOAD_TYPE", event.EventType+" payload is not as expected")
	}
}
