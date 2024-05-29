package create_gym

import (
	"gym-management/src/components/gyms/core/domain"
	"gym-management/src/lib/primitives/application_specific"
)

type CreateGymCommandHandler struct {
	EventsPublisher    domain.EventsPublisher
	GymOwnerRepository domain.GymOwnerRepository
}

func (h CreateGymCommandHandler) Handle(command *CreateGymCommand) (*CreateGymCommandResponse, *application_specific.ApplicationException) {
	owner, err := h.GymOwnerRepository.FindByID(command.OwnerId, command.Session.Session)
	if err != nil {
		return nil, err
	}
	if owner == nil {
		return nil, application_specific.NewNotFoundException("GYMS.OWNERS.NOT_FOUND", "Gym owner not found", map[string]string{
			"id": command.OwnerId,
		})
	}

	gym, err := owner.CreateGym(command.Name, command.Address, command.Session.User.Id)
	if err != nil {
		return nil, err
	}

	err = h.GymOwnerRepository.Update(owner, command.Session.Session)
	if err != nil {
		return nil, err
	}

	err = h.EventsPublisher.Publish(owner.PullEvents(), command.Session.Session)
	if err != nil {
		return nil, err
	}

	return &CreateGymCommandResponse{
		Id: gym.State().Id,
	}, nil
}
