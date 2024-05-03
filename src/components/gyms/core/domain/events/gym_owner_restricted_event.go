package events

import (
	"gym-management/src/components/gyms/core/domain"
)

const GymOwnerRestrictedEventType = "Gyms.Owners.Restricted"

type GymOwnerRestrictedEventPayload struct {
	Id           string
	RestrictedBy string
}

func NewGymOwnerRestrictedEvent(state domain.GymOwnerState) *GymEvent[interface{}] {
	return NewGymEvent(GymOwnerRestrictedEventType, GymOwnerRestrictedEventPayload{
		Id:           state.Id,
		RestrictedBy: state.UpdatedBy,
	})
}
