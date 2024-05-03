package domain

import "gym-management/src/components/gyms/core/domain/events"

func NewGymOwnerUpdatedEvent(state GymOwnerState, newPassword *string) *events.GymEvent[interface{}] {
	return events.NewGymEvent(events.GymOwnerUpdatedEventType, events.GymOwnerUpdatedEventPayload{
		Id:          state.Id,
		Name:        state.Name,
		PhoneNumber: state.PhoneNumber,
		Email:       state.Email,
		Restricted:  state.Restricted,
		NewPassword: newPassword,
		UpdatedBy:   state.UpdatedBy,
	})
}

func NewGymOwnerUnrestrictedEvent(state GymOwnerState) *events.GymEvent[interface{}] {
	return events.NewGymEvent(events.GymOwnerUnrestrictedEventType, events.GymOwnerUnrestrictedEventPayload{
		Id:             state.Id,
		UnrestrictedBy: state.UpdatedBy,
	})
}

func NewGymOwnerRestrictedEvent(state GymOwnerState) *events.GymEvent[interface{}] {
	return events.NewGymEvent(events.GymOwnerRestrictedEventType, events.GymOwnerRestrictedEventPayload{
		Id:           state.Id,
		RestrictedBy: state.UpdatedBy,
	})
}

func NewGymOwnerDeletedEvent(state GymOwnerState) *events.GymEvent[interface{}] {
	return events.NewGymEvent(events.GymOwnerDeletedEventType, events.GymOwnerDeletedEventPayload{
		Id:        state.Id,
		DeletedBy: *state.DeletedBy,
	})
}

func NewGymOwnerCreatedEvent(state GymOwnerState, password string) *events.GymEvent[interface{}] {
	return events.NewGymEvent(events.GymOwnerCreatedEventType, events.GymOwnerCreatedEventPayload{
		Id:          state.Id,
		Name:        state.Name,
		PhoneNumber: state.PhoneNumber,
		Email:       state.Email,
		Password:    password,
		CreatedBy:   state.CreatedBy,
	})
}
