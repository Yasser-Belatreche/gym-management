package gyms

import (
	"gym-management-gyms/src/components/gyms/core/domain"
	"gym-management-gyms/src/components/gyms/core/usecases/gym_owners/create_gym_owner"
	"gym-management-gyms/src/components/gyms/core/usecases/gym_owners/delete_gym_owner"
	"gym-management-gyms/src/components/gyms/core/usecases/gym_owners/get_gym_owner"
	"gym-management-gyms/src/components/gyms/core/usecases/gym_owners/get_gym_owners"
	"gym-management-gyms/src/components/gyms/core/usecases/gym_owners/restrict_gym_owner"
	"gym-management-gyms/src/components/gyms/core/usecases/gym_owners/unrestrict_gym_owner"
	"gym-management-gyms/src/components/gyms/core/usecases/gym_owners/update_gym_owner"
	"gym-management-gyms/src/components/gyms/core/usecases/gyms/create_gym"
	"gym-management-gyms/src/components/gyms/core/usecases/gyms/delete_gym"
	"gym-management-gyms/src/components/gyms/core/usecases/gyms/disable_gym"
	"gym-management-gyms/src/components/gyms/core/usecases/gyms/enable_gym"
	"gym-management-gyms/src/components/gyms/core/usecases/gyms/get_gym"
	"gym-management-gyms/src/components/gyms/core/usecases/gyms/get_gyms"
	"gym-management-gyms/src/components/gyms/core/usecases/gyms/update_gym"
	"gym-management-gyms/src/lib/primitives/application_specific"
)

type Facade struct {
	EmailService       domain.EmailService
	EventsPublisher    domain.EventsPublisher
	GymOwnerRepository domain.GymOwnerRepository
}

func (f *Facade) GetGymOwner(query *get_gym_owner.GetGymOwnerQuery) (*get_gym_owner.GetGymOwnerQueryResponse, *application_specific.ApplicationException) {
	handler := &get_gym_owner.GetGymOwnerQueryHandler{}

	return handler.Handle(query)
}

func (f *Facade) GetGymOwners(query *get_gym_owners.GetGymOwnersQuery) (*get_gym_owners.GetGymOwnersQueryResponse, *application_specific.ApplicationException) {
	handler := &get_gym_owners.GetGymOwnersQueryHandler{}

	return handler.Handle(query)
}

func (f *Facade) CreateGymOwner(command *create_gym_owner.CreateGymOwnerCommand) (*create_gym_owner.CreateGymOwnerCommandResponse, *application_specific.ApplicationException) {
	handler := &create_gym_owner.CreateGymOwnerCommandHandler{
		EmailService:       f.EmailService,
		EventsPublisher:    f.EventsPublisher,
		GymOwnerRepository: f.GymOwnerRepository,
	}

	return handler.Handle(command)
}

func (f *Facade) UpdateGymOwner(command *update_gym_owner.UpdateGymOwnerCommand) (*update_gym_owner.UpdateGymOwnerCommandResponse, *application_specific.ApplicationException) {
	handler := &update_gym_owner.UpdateGymOwnerCommandHandler{
		EmailService:       f.EmailService,
		EventsPublisher:    f.EventsPublisher,
		GymOwnerRepository: f.GymOwnerRepository,
	}

	return handler.Handle(command)
}

func (f *Facade) DeleteGymOwner(command *delete_gym_owner.DeleteGymOwnerCommand) (*delete_gym_owner.DeleteGymOwnerCommandResponse, *application_specific.ApplicationException) {
	handler := &delete_gym_owner.DeleteGymOwnerCommandHandler{
		EventsPublisher:    f.EventsPublisher,
		GymOwnerRepository: f.GymOwnerRepository,
	}

	return handler.Handle(command)
}

func (f *Facade) RestrictGymOwner(command *restrict_gym_owner.RestrictGymOwnerCommand) (*restrict_gym_owner.RestrictGymOwnerCommandResponse, *application_specific.ApplicationException) {
	handler := &restrict_gym_owner.RestrictGymOwnerCommandHandler{
		EventsPublisher:    f.EventsPublisher,
		GymOwnerRepository: f.GymOwnerRepository,
	}

	return handler.Handle(command)
}

func (f *Facade) UnrestrictGymOwner(command *unrestrict_gym_owner.UnrestrictGymOwnerCommand) (*unrestrict_gym_owner.UnrestrictGymOwnerCommandResponse, *application_specific.ApplicationException) {
	handler := &unrestrict_gym_owner.UnrestrictGymOwnerCommandHandler{
		EventsPublisher:    f.EventsPublisher,
		GymOwnerRepository: f.GymOwnerRepository,
	}

	return handler.Handle(command)
}

func (f *Facade) GetGym(query *get_gym.GetGymQuery) (*get_gym.GetGymQueryResponse, *application_specific.ApplicationException) {
	handler := &get_gym.GetGymQueryHandler{}

	return handler.Handle(query)
}

func (f *Facade) GetGyms(query *get_gyms.GetGymsQuery) (*get_gyms.GetGymsQueryResponse, *application_specific.ApplicationException) {
	handler := &get_gyms.GetGymsQueryHandler{}

	return handler.Handle(query)
}

func (f *Facade) CreateGym(command *create_gym.CreateGymCommand) (*create_gym.CreateGymCommandResponse, *application_specific.ApplicationException) {
	handler := &create_gym.CreateGymCommandHandler{
		EventsPublisher:    f.EventsPublisher,
		GymOwnerRepository: f.GymOwnerRepository,
	}

	return handler.Handle(command)
}

func (f *Facade) UpdateGym(command *update_gym.UpdateGymCommand) (*update_gym.UpdateGymCommandResponse, *application_specific.ApplicationException) {
	handler := &update_gym.UpdateGymCommandHandler{
		EventsPublisher:    f.EventsPublisher,
		GymOwnerRepository: f.GymOwnerRepository,
	}

	return handler.Handle(command)
}

func (f *Facade) DeleteGym(command *delete_gym.DeleteGymCommand) (*delete_gym.DeleteGymCommandResponse, *application_specific.ApplicationException) {
	handler := &delete_gym.DeleteGymCommandHandler{
		EventsPublisher:    f.EventsPublisher,
		GymOwnerRepository: f.GymOwnerRepository,
	}

	return handler.Handle(command)
}

func (f *Facade) EnableGym(command *enable_gym.EnableGymCommand) (*enable_gym.EnableGymCommandResponse, *application_specific.ApplicationException) {
	handler := &enable_gym.EnableGymCommandHandler{
		EventsPublisher:    f.EventsPublisher,
		GymOwnerRepository: f.GymOwnerRepository,
	}

	return handler.Handle(command)
}

func (f *Facade) DisableGym(command *disable_gym.DisableGymCommand) (*disable_gym.DisableGymCommandResponse, *application_specific.ApplicationException) {
	handler := &disable_gym.DisableGymCommandHandler{
		EventsPublisher:    f.EventsPublisher,
		GymOwnerRepository: f.GymOwnerRepository,
	}

	return handler.Handle(command)
}
