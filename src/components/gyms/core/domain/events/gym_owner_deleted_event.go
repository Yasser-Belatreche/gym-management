package events

import (
	"gym-management/src/components/gyms/core/domain"
)

const GymOwnerDeletedEventType = "Gyms.Owners.Deleted"

type GymOwnerDeletedEventPayload struct {
	Id        string
	DeletedBy string
}

func NewGymOwnerDeletedEvent(state domain.GymOwnerState) *GymEvent[interface{}] {
	return NewGymEvent(GymOwnerDeletedEventType, GymOwnerDeletedEventPayload{
		Id:        state.Id,
		DeletedBy: *state.DeletedBy,
	})
}
