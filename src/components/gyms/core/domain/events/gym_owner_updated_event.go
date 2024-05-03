package events

import (
	"gym-management/src/components/gyms/core/domain"
)

const GymOwnerUpdatedEventType = "Gyms.Owners.Updated"

type GymOwnerUpdatedEventPayload struct {
	Id          string
	Name        string
	PhoneNumber string
	Email       string
	Restricted  bool
	NewPassword *string
	UpdatedBy   string
}

func NewGymOwnerUpdatedEvent(state domain.GymOwnerState, newPassword *string) *GymEvent[interface{}] {
	return NewGymEvent(GymOwnerUpdatedEventType, GymOwnerUpdatedEventPayload{
		Id:          state.Id,
		Name:        state.Name,
		PhoneNumber: state.PhoneNumber,
		Email:       state.Email,
		Restricted:  state.Restricted,
		NewPassword: newPassword,
		UpdatedBy:   state.UpdatedBy,
	})
}
