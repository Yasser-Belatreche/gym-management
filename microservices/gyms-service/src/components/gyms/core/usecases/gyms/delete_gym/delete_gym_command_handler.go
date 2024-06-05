package delete_gym

import (
	"gym-management-gyms/src/components/gyms/core/domain"
	"gym-management-gyms/src/lib/primitives/application_specific"
)

type DeleteGymCommandHandler struct {
	EventsPublisher    domain.EventsPublisher
	GymOwnerRepository domain.GymOwnerRepository
}

func (h DeleteGymCommandHandler) Handle(command *DeleteGymCommand) (*DeleteGymCommandResponse, *application_specific.ApplicationException) {
	owner, err := h.GymOwnerRepository.FindByID(command.OwnerId, command.Session.Session)
	if err != nil {
		return nil, err
	}
	if owner == nil {
		return nil, application_specific.NewNotFoundException("GYMS.OWNERS.NOT_FOUND", "Gym owner not found", map[string]string{
			"id": command.OwnerId,
		})
	}

	err = owner.DeleteGym(command.GymId, command.Session.User.Id)
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

	return &DeleteGymCommandResponse{
		Id: command.GymId,
	}, nil
}
