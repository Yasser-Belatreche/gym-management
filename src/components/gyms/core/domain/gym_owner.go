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

	return &GymOwner{
		id:          state.Id,
		name:        state.Name,
		phoneNumber: state.PhoneNumber,
		email:       email,
		restricted:  state.Restricted,
		createdBy:   state.CreatedBy,
		updatedBy:   state.UpdatedBy,
		deletedAt:   state.DeletedAt,
		events:      make([]*events.GymEvent[any], 0),
	}
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

	now := time.Now()
	owner.deletedAt = &now
	owner.deleteBy = &by

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
	return GymOwnerState{
		Id:          owner.id,
		Name:        owner.name,
		PhoneNumber: owner.phoneNumber,
		Email:       owner.email.Value,
		Restricted:  owner.restricted,
		DeletedBy:   owner.deleteBy,
		CreatedBy:   owner.createdBy,
		UpdatedBy:   owner.updatedBy,
		DeletedAt:   owner.deletedAt,
	}
}

func (owner *GymOwner) PullEvents() []*events.GymEvent[any] {
	return owner.events
}
