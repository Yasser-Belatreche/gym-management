package events

import "gym-management/src/lib/primitives/application_specific"

type MembershipEvent[T interface{}] struct {
	*application_specific.DomainEvent[T]
}

func NewMembershipEvent(eventType string, payload interface{}) *MembershipEvent[interface{}] {
	return &MembershipEvent[interface{}]{
		DomainEvent: application_specific.NewDomainEvent(eventType, payload),
	}
}
