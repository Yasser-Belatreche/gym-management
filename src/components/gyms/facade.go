package gyms

import (
	"gym-management/src/components/gyms/core/domain"
	"gym-management/src/components/gyms/core/usecases/gym_owners/create_gym_owner"
	"gym-management/src/lib/primitives/application_specific"
)

type Facade struct {
	EmailService       domain.EmailService
	EventsPublisher    domain.EventsPublisher
	GymOwnerRepository domain.GymOwnerRepository
}

func (f *Facade) CreateGymOwner(command *create_gym_owner.CreateGymOwnerCommand) (*create_gym_owner.CreateGymOwnerCommandResponse, *application_specific.ApplicationException) {
	handler := &create_gym_owner.CreateGymOwnerCommandHandler{
		EmailService:       f.EmailService,
		EventsPublisher:    f.EventsPublisher,
		GymOwnerRepository: f.GymOwnerRepository,
	}

	return handler.Handle(command)
}
