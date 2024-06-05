package domain

import (
	"gym-management/src/components/gyms/core/domain/events"
	"gym-management/src/lib/primitives/application_specific"
)

type EventsPublisher interface {
	Publish(events []*events.GymEvent[interface{}], session *application_specific.Session) *application_specific.ApplicationException
}
