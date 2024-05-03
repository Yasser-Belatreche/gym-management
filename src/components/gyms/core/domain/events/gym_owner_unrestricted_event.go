package events

import (
	"gym-management/src/components/gyms/core/domain"
)

const GymOwnerUnrestrictedEventType = "Gyms.Owners.Unrestricted"

type GymOwnerUnrestrictedEventPayload struct {
	Id             string
	UnrestrictedBy string
}

func NewGymOwnerUnrestrictedEvent(state domain.GymOwnerState) *GymEvent[interface{}] {
	return NewGymEvent(GymOwnerUnrestrictedEventType, GymOwnerUnrestrictedEventPayload{
		Id:             state.Id,
		UnrestrictedBy: state.UpdatedBy,
	})
}
