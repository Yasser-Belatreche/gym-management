package events

import (
	"gym-management/src/components/gyms/core/domain"
)

const GymOwnerCreatedEventType = "Gyms.Owners.Created"

type GymOwnerCreatedEventPayload struct {
	Id          string
	Name        string
	PhoneNumber string
	Email       string
	Password    string
	CreatedBy   string
}

func NewGymOwnerCreatedEvent(state domain.GymOwnerState, password string) *GymEvent[interface{}] {
	return NewGymEvent(GymOwnerCreatedEventType, GymOwnerCreatedEventPayload{
		Id:          state.Id,
		Name:        state.Name,
		PhoneNumber: state.PhoneNumber,
		Email:       state.Email,
		Password:    password,
		CreatedBy:   state.CreatedBy,
	})
}
