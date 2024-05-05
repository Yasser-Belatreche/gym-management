package domain

import (
	"gym-management/src/components/gyms/core/domain/events"
	"gym-management/src/lib/primitives/application_specific"
	"gym-management/src/lib/primitives/generic"
	"strings"
	"time"
)

type GymOwner struct {
	id          string
	name        string
	phoneNumber string
	email       application_specific.Email
	restricted  bool
	gyms        []*Gym
	createdBy   string
	updatedBy   string
	deletedAt   *time.Time
	deleteBy    *string
	events      []*events.GymEvent[interface{}]
}

type GymOwnerState struct {
	Id          string
	Name        string
	PhoneNumber string
	Email       string
	Restricted  bool
	Gyms        []GymState
	CreatedBy   string
	UpdatedBy   string
	DeletedBy   *string
	DeletedAt   *time.Time
}

func CreateGymOwner(name string, phoneNumber string, email application_specific.Email, password string, by string) (*GymOwner, *application_specific.ApplicationException) {
	name = strings.TrimSpace(name)
	phoneNumber = strings.TrimSpace(phoneNumber)
	password = strings.TrimSpace(password)

	if name == "" {
		return nil, application_specific.NewValidationException("GYMS.OWNER.NAME_REQUIRED", "Name is required", map[string]string{
			"name": name,
		})
	}

	if phoneNumber == "" {
		return nil, application_specific.NewValidationException("GYMS.OWNER.PHONE_NUMBER_REQUIRED", "Phone number is required", map[string]string{
			"phoneNumber": phoneNumber,
		})
	}

	if len(password) < 6 {
		return nil, application_specific.NewValidationException("GYMS.OWNER.PASSWORD_TOO_SHORT", "Password is too short", map[string]string{
			"password": password,
		})
	}

	owner := &GymOwner{
		id:          generic.GenerateUUID(),
		name:        name,
		phoneNumber: phoneNumber,
		email:       email,
		restricted:  false,
		createdBy:   by,
		updatedBy:   by,
		deletedAt:   nil,
		deleteBy:    nil,
		events:      make([]*events.GymEvent[interface{}], 0),
	}

	owner.events = append(
		owner.events,
		NewGymOwnerCreatedEvent(owner.State(), password),
	)

	return owner, nil
}

func GymOwnerFromState(state GymOwnerState) *GymOwner {
	email, _ := application_specific.NewEmail(state.Email)

	gyms := make([]*Gym, 0)

	for _, gymState := range state.Gyms {
		gyms = append(gyms, GymFromState(gymState))
	}

	return &GymOwner{
		id:          state.Id,
		name:        state.Name,
		phoneNumber: state.PhoneNumber,
		email:       email,
		restricted:  state.Restricted,
		gyms:        gyms,
		deleteBy:    state.DeletedBy,
		createdBy:   state.CreatedBy,
		updatedBy:   state.UpdatedBy,
		deletedAt:   state.DeletedAt,
		events:      make([]*events.GymEvent[any], 0),
	}
}

func (owner *GymOwner) CreateGym(name string, address string, by string) *application_specific.ApplicationException {
	if owner.restricted {
		return application_specific.NewValidationException("GYMS.OWNER.RESTRICTED", "Owner is restricted", map[string]string{
			"id": owner.id,
		})
	}

	if owner.deletedAt != nil {
		return application_specific.NewValidationException("GYMS.OWNER.DELETED", "Owner is deleted", map[string]string{
			"id": owner.id,
		})
	}

	gym, err := CreateGym(name, address, by)
	if err != nil {
		return err
	}

	owner.gyms = append(owner.gyms, gym)
	owner.updatedBy = gym.createdBy

	owner.events = append(
		owner.events,
		NewGymCreatedEvent(gym.State(), owner.State()),
	)
	owner.events = append(
		owner.events,
		NewGymOwnerUpdatedEvent(owner.State(), nil),
	)

	return nil
}

func (owner *GymOwner) DisableGym(id, by string) *application_specific.ApplicationException {
	if owner.restricted {
		return application_specific.NewValidationException("GYMS.OWNER.RESTRICTED", "Owner is restricted", map[string]string{
			"id": owner.id,
		})
	}

	if owner.deletedAt != nil {
		return application_specific.NewValidationException("GYMS.OWNER.DELETED", "Owner is deleted", map[string]string{
			"id": owner.id,
		})
	}

	gym := owner.findGymById(id)
	if gym == nil {
		return application_specific.NewNotFoundException("GYMS.GYM_NOT_FOUND", "Gym not found", map[string]string{
			"id": id,
		})
	}

	err := gym.Disable(by, "Owner disabled the gym")
	if err != nil {
		return err
	}

	owner.updatedBy = by

	owner.events = append(
		owner.events,
		NewGymDisabledEvent(gym.State(), owner.State()),
	)
	owner.events = append(
		owner.events,
		NewGymOwnerUpdatedEvent(owner.State(), nil),
	)

	return nil
}

func (owner *GymOwner) EnableGym(id, by string) *application_specific.ApplicationException {
	if owner.restricted {
		return application_specific.NewValidationException("GYMS.OWNER.RESTRICTED", "Owner is restricted", map[string]string{
			"id": owner.id,
		})
	}

	if owner.deletedAt != nil {
		return application_specific.NewValidationException("GYMS.OWNER.DELETED", "Owner is deleted", map[string]string{
			"id": owner.id,
		})
	}

	gym := owner.findGymById(id)
	if gym == nil {
		return application_specific.NewNotFoundException("GYMS.GYM_NOT_FOUND", "Gym not found", map[string]string{
			"id": id,
		})
	}

	err := gym.Enable(by)
	if err != nil {
		return err
	}

	owner.updatedBy = by

	owner.events = append(
		owner.events,
		NewGymEnabledEvent(gym.State(), owner.State()),
	)
	owner.events = append(
		owner.events,
		NewGymOwnerUpdatedEvent(owner.State(), nil),
	)

	return nil
}

func (owner *GymOwner) DeleteGym(id, by string) *application_specific.ApplicationException {
	if owner.restricted {
		return application_specific.NewValidationException("GYMS.OWNER.RESTRICTED", "Owner is restricted", map[string]string{
			"id": owner.id,
		})
	}

	if owner.deletedAt != nil {
		return application_specific.NewValidationException("GYMS.OWNER.DELETED", "Owner is deleted", map[string]string{
			"id": owner.id,
		})
	}

	gym := owner.findGymById(id)
	if gym == nil {
		return application_specific.NewNotFoundException("GYMS.GYM_NOT_FOUND", "Gym not found", map[string]string{
			"id": id,
		})
	}

	err := gym.Disable(by, "Owner deleted the gym")
	if err != nil {
		return err
	}

	err = gym.Delete(by)
	if err != nil {
		return err
	}

	owner.updatedBy = by

	owner.events = append(
		owner.events,
		NewGymDisabledEvent(gym.State(), owner.State()),
	)
	owner.events = append(
		owner.events,
		NewGymDeletedEvent(gym.State(), owner.State()),
	)
	owner.events = append(
		owner.events,
		NewGymOwnerUpdatedEvent(owner.State(), nil),
	)

	return nil
}

func (owner *GymOwner) UpdateGym(id, name, address, by string) *application_specific.ApplicationException {
	if owner.restricted {
		return application_specific.NewValidationException("GYMS.OWNER.RESTRICTED", "Owner is restricted", map[string]string{
			"id": owner.id,
		})
	}

	if owner.deletedAt != nil {
		return application_specific.NewValidationException("GYMS.OWNER.DELETED", "Owner is deleted", map[string]string{
			"id": owner.id,
		})
	}

	gym := owner.findGymById(id)
	if gym == nil {
		return application_specific.NewNotFoundException("GYMS.GYM_NOT_FOUND", "Gym not found", map[string]string{
			"id": id,
		})
	}

	err := gym.Update(name, address, by)
	if err != nil {
		return err
	}

	owner.updatedBy = by

	owner.events = append(
		owner.events,
		NewGymUpdatedEvent(gym.State(), owner.State()),
	)
	owner.events = append(
		owner.events,
		NewGymOwnerUpdatedEvent(owner.State(), nil),
	)

	return nil
}

func (owner *GymOwner) findGymById(id string) *Gym {
	for _, gym := range owner.gyms {
		if gym.id == id {
			return gym
		}
	}

	return nil
}

func (owner *GymOwner) Update(name string, phoneNumber string, email application_specific.Email, password *string, updatedBy string) *application_specific.ApplicationException {
	if owner.restricted {
		return application_specific.NewValidationException("GYMS.OWNER.RESTRICTED", "Owner is restricted", map[string]string{
			"id": owner.id,
		})
	}

	if owner.deletedAt != nil {
		return application_specific.NewValidationException("GYMS.OWNER.DELETED", "Owner is deleted", map[string]string{
			"id": owner.id,
		})
	}

	name = strings.TrimSpace(name)
	phoneNumber = strings.TrimSpace(phoneNumber)

	if name == "" {
		return application_specific.NewValidationException("GYMS.OWNER.NAME_REQUIRED", "Name is required", map[string]string{
			"name": name,
		})
	}

	if phoneNumber == "" {
		return application_specific.NewValidationException("GYMS.OWNER.PHONE_NUMBER_REQUIRED", "Phone number is required", map[string]string{
			"phoneNumber": phoneNumber,
		})
	}

	owner.name = name
	owner.phoneNumber = phoneNumber
	owner.email = email
	owner.updatedBy = updatedBy

	owner.events = append(
		owner.events,
		NewGymOwnerUpdatedEvent(owner.State(), password),
	)

	return nil
}

func (owner *GymOwner) Restrict(by string) *application_specific.ApplicationException {
	if owner.deletedAt != nil {
		return application_specific.NewValidationException("GYMS.OWNER.DELETED", "Owner is deleted", map[string]string{
			"id": owner.id,
		})
	}

	if owner.restricted {
		return application_specific.NewValidationException("GYMS.OWNER.ALREADY_RESTRICTED", "Owner is already restricted", map[string]string{
			"id": owner.id,
		})
	}

	owner.restricted = true
	owner.updatedBy = by

	for _, gym := range owner.gyms {
		if gym.enabled {
			err := gym.Disable(by, "Owner is restricted")
			if err != nil {
				return err
			}

			owner.events = append(
				owner.events,
				NewGymDisabledEvent(gym.State(), owner.State()),
			)
		}
	}

	owner.events = append(
		owner.events,
		NewGymOwnerUpdatedEvent(owner.State(), nil),
	)
	owner.events = append(
		owner.events,
		NewGymOwnerRestrictedEvent(owner.State()),
	)

	return nil
}

func (owner *GymOwner) Unrestrict(by string) *application_specific.ApplicationException {
	if owner.deletedAt != nil {
		return application_specific.NewValidationException("GYMS.OWNER.DELETED", "Owner is deleted", map[string]string{
			"id": owner.id,
		})
	}

	if !owner.restricted {
		return application_specific.NewValidationException("GYMS.OWNER.NOT_RESTRICTED", "Owner is not restricted", map[string]string{
			"id": owner.id,
		})
	}

	owner.restricted = false
	owner.updatedBy = by

	for _, gym := range owner.gyms {
		if gym.IsDisabledBecause("Owner is restricted") {
			err := gym.Enable(by)
			if err != nil {
				return err
			}

			owner.events = append(
				owner.events,
				NewGymEnabledEvent(gym.State(), owner.State()),
			)
		}
	}

	owner.events = append(
		owner.events,
		NewGymOwnerUpdatedEvent(owner.State(), nil),
	)
	owner.events = append(
		owner.events,
		NewGymOwnerUnrestrictedEvent(owner.State()),
	)

	return nil
}

func (owner *GymOwner) Delete(by string) *application_specific.ApplicationException {
	if owner.deletedAt != nil {
		return application_specific.NewValidationException("GYMS.OWNER.ALREADY_DELETED", "Owner is already deleted", map[string]string{
			"id": owner.id,
		})
	}

	owner.restricted = true
	now := time.Now()
	owner.deletedAt = &now
	owner.deleteBy = &by

	for _, gym := range owner.gyms {
		if gym.enabled {
			err := gym.Disable(by, "Owner is deleted")
			if err != nil {
				return err
			}

			owner.events = append(
				owner.events,
				NewGymDisabledEvent(gym.State(), owner.State()),
			)
		}
	}

	owner.events = append(
		owner.events,
		NewGymOwnerDeletedEvent(owner.State()),
	)

	return nil
}

func (owner *GymOwner) EmailIs(email application_specific.Email) bool {
	return owner.email.Equals(email)
}

func (owner *GymOwner) State() GymOwnerState {
	gyms := make([]GymState, 0)

	for _, gym := range owner.gyms {
		gyms = append(gyms, gym.State())
	}

	return GymOwnerState{
		Id:          owner.id,
		Name:        owner.name,
		PhoneNumber: owner.phoneNumber,
		Email:       owner.email.Value,
		Restricted:  owner.restricted,
		Gyms:        gyms,
		CreatedBy:   owner.createdBy,
		UpdatedBy:   owner.updatedBy,
		DeletedBy:   owner.deleteBy,
		DeletedAt:   owner.deletedAt,
	}
}

func (owner *GymOwner) PullEvents() []*events.GymEvent[any] {
	return owner.events
}
