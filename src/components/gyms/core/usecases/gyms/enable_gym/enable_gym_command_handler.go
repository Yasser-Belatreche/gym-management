package enable_gym

import (
	"gym-management/src/components/gyms/core/domain"
	"gym-management/src/lib/primitives/application_specific"
)

type EnableGymCommandHandler struct {
	EventsPublisher    domain.EventsPublisher
	GymOwnerRepository domain.GymOwnerRepository
}

func (h EnableGymCommandHandler) Handle(command *EnableGymCommand) (*EnableGymCommandResponse, *application_specific.ApplicationException) {
	owner, err := h.GymOwnerRepository.FindByID(command.OwnerId, command.Session.Session)
	if err != nil {
		return nil, err
	}
	if owner == nil {
		return nil, application_specific.NewNotFoundException("GYMS.OWNERS.NOT_FOUND", "Gym owner not found", map[string]string{
			"id": command.OwnerId,
		})
	}

	err = owner.EnableGym(command.GymId, command.Session.User.Id)
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

	return &EnableGymCommandResponse{
		Id: command.GymId,
	}, nil
}
