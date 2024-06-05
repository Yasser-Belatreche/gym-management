package domain

import "gym-management/src/components/gyms/core/domain/events"

func NewGymOwnerUpdatedEvent(state GymOwnerState, newPassword *string) *events.GymEvent[interface{}] {
	gyms := make([]string, 0)
	enabledGyms := make([]string, 0)

	for _, gym := range state.Gyms {
		gyms = append(gyms, gym.Id)

		if gym.Enabled {
			enabledGyms = append(enabledGyms, gym.Id)
		}
	}

	return events.NewGymEvent(events.GymOwnerUpdatedEventType, events.GymOwnerUpdatedEventPayload{
		Id:          state.Id,
		Name:        state.Name,
		PhoneNumber: state.PhoneNumber,
		Email:       state.Email,
		Restricted:  state.Restricted,
		NewPassword: newPassword,
		UpdatedBy:   state.UpdatedBy,
		Gyms:        gyms,
		EnabledGyms: enabledGyms,
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
	gyms := make([]string, 0)

	for _, gym := range state.Gyms {
		if gym.Enabled {
			gyms = append(gyms, gym.Id)
		}
	}

	return events.NewGymEvent(events.GymOwnerCreatedEventType, events.GymOwnerCreatedEventPayload{
		Id:          state.Id,
		Name:        state.Name,
		PhoneNumber: state.PhoneNumber,
		Email:       state.Email,
		Password:    password,
		CreatedBy:   state.CreatedBy,
		Gyms:        gyms,
	})
}

func NewGymCreatedEvent(state GymState, owner GymOwnerState) *events.GymEvent[interface{}] {
	return events.NewGymEvent(events.GymCreatedEventType, events.GymCreatedEventPayload{
		Id:        state.Id,
		Name:      state.Name,
		OwnerId:   owner.Id,
		Address:   state.Address,
		CreatedBy: state.CreatedBy,
	})
}

func NewGymUpdatedEvent(state GymState, owner GymOwnerState) *events.GymEvent[interface{}] {
	return events.NewGymEvent(events.GymUpdatedEventType, events.GymUpdatedEventPayload{
		Id:        state.Id,
		Name:      state.Name,
		Address:   state.Address,
		UpdatedBy: state.UpdatedBy,
		OwnerId:   owner.Id,
	})
}

func NewGymDeletedEvent(state GymState, owner GymOwnerState) *events.GymEvent[interface{}] {
	return events.NewGymEvent(events.GymDeletedEventType, events.GymDeletedEventPayload{
		Id:        state.Id,
		DeletedBy: *state.DeletedBy,
		OwnerId:   owner.Id,
	})
}

func NewGymEnabledEvent(state GymState, owner GymOwnerState) *events.GymEvent[interface{}] {
	return events.NewGymEvent(events.GymEnabledEventType, events.GymEnabledEventPayload{
		Id:        state.Id,
		EnabledBy: state.UpdatedBy,
		OwnerId:   owner.Id,
	})
}

func NewGymDisabledEvent(state GymState, owner GymOwnerState) *events.GymEvent[interface{}] {
	return events.NewGymEvent(events.GymDisabledEventType, events.GymDisabledEventPayload{
		Id:          state.Id,
		DisabledBy:  state.UpdatedBy,
		OwnerId:     owner.Id,
		DisabledFor: *state.DisabledFor,
	})
}
