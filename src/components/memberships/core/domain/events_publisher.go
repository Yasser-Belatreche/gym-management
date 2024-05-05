package domain

import (
	"gym-management/src/components/memberships/core/domain/events"
	"gym-management/src/lib/primitives/application_specific"
)

type EventsPublisher interface {
	Publish(events []*events.MembershipEvent[interface{}], session *application_specific.Session) *application_specific.ApplicationException
}
