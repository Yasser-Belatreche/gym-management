package domain

import (
	"gym-management-gyms/src/lib/primitives/application_specific"
	"gym-management-gyms/src/lib/primitives/generic"
	"strings"
	"time"
)

type Gym struct {
	id          string
	name        string
	address     string
	enabled     bool
	disabledFor *string
	createdBy   string
	updatedBy   string
	deletedAt   *time.Time
	deletedBy   *string
}

type GymState struct {
	Id          string
	Name        string
	Address     string
	Enabled     bool
	DisabledFor *string
	CreatedBy   string
	UpdatedBy   string
	DeletedAt   *time.Time
	DeletedBy   *string
}

func CreateGym(name string, address string, by string) (*Gym, *application_specific.ApplicationException) {
	name = strings.TrimSpace(name)
	address = strings.TrimSpace(address)

	if name == "" {
		return nil, application_specific.NewValidationException("GYMS.NAME_REQUIRED", "Name is required", map[string]string{
			"name": name,
		})
	}

	if address == "" {
		return nil, application_specific.NewValidationException("GYMS.ADDRESS_REQUIRED", "Address is required", map[string]string{
			"address": address,
		})
	}

	gym := &Gym{
		id:          generic.GenerateULID(),
		name:        name,
		address:     address,
		enabled:     true,
		disabledFor: nil,
		createdBy:   by,
		updatedBy:   by,
		deletedAt:   nil,
		deletedBy:   nil,
	}

	return gym, nil
}

func GymFromState(state GymState) *Gym {
	return &Gym{
		id:          state.Id,
		name:        state.Name,
		address:     state.Address,
		enabled:     state.Enabled,
		createdBy:   state.CreatedBy,
		disabledFor: state.DisabledFor,
		updatedBy:   state.UpdatedBy,
		deletedAt:   state.DeletedAt,
		deletedBy:   state.DeletedBy,
	}
}

func (g *Gym) Disable(by string, reason string) *application_specific.ApplicationException {
	if g.deletedAt != nil {
		return application_specific.NewValidationException("GYMS.ALREADY_DELETED", "Gym is already deleted", map[string]string{
			"id": g.id,
		})
	}

	if !g.enabled {
		return application_specific.NewValidationException("GYMS.ALREADY_DISABLED", "Gym is already disabled", map[string]string{
			"id": g.id,
		})
	}

	g.disabledFor = &reason
	g.enabled = false
	g.updatedBy = by

	return nil
}

func (g *Gym) Enable(by string) *application_specific.ApplicationException {
	if g.deletedAt != nil {
		return application_specific.NewValidationException("GYMS.ALREADY_DELETED", "Gym is already deleted", map[string]string{
			"id": g.id,
		})
	}

	if g.enabled {
		return application_specific.NewValidationException("GYMS.ALREADY_ENABLED", "Gym is already enabled", map[string]string{
			"id": g.id,
		})
	}

	g.disabledFor = nil
	g.enabled = true
	g.updatedBy = by

	return nil
}

func (g *Gym) Update(name string, address string, by string) *application_specific.ApplicationException {
	if g.deletedAt != nil {
		return application_specific.NewValidationException("GYMS.ALREADY_DELETED", "Gym is already deleted", map[string]string{
			"id": g.id,
		})
	}

	if !g.enabled {
		return application_specific.NewValidationException("GYMS.DISABLED", "Gym is disabled", map[string]string{
			"id": g.id,
		})
	}

	name = strings.TrimSpace(name)
	address = strings.TrimSpace(address)

	if name == "" {
		return application_specific.NewValidationException("GYMS.NAME_REQUIRED", "Name is required", map[string]string{
			"name": name,
		})
	}

	if address == "" {
		return application_specific.NewValidationException("GYMS.ADDRESS_REQUIRED", "Address is required", map[string]string{
			"address": address,
		})
	}

	g.name = name
	g.address = address
	g.updatedBy = by

	return nil
}

func (g *Gym) Delete(by string) *application_specific.ApplicationException {
	if g.deletedAt != nil {
		return application_specific.NewValidationException("GYMS.ALREADY_DELETED", "Gym is already deleted", map[string]string{
			"id": g.id,
		})
	}

	now := time.Now()
	g.deletedAt = &now
	g.deletedBy = &by

	return nil
}

func (g *Gym) IsDisabledBecause(str string) bool {
	if g.disabledFor == nil {
		return false
	}

	return *g.disabledFor == str
}

func (g *Gym) State() GymState {
	return GymState{
		Id:          g.id,
		Name:        g.name,
		Address:     g.address,
		Enabled:     g.enabled,
		CreatedBy:   g.createdBy,
		DisabledFor: g.disabledFor,
		UpdatedBy:   g.updatedBy,
		DeletedAt:   g.deletedAt,
		DeletedBy:   g.deletedBy,
	}
}
