package delete_gym_owner

import (
	"gym-management-gyms/src/components/gyms/core/domain"
	"gym-management-gyms/src/lib/primitives/application_specific"
)

type DeleteGymOwnerCommandHandler struct {
	EventsPublisher    domain.EventsPublisher
	GymOwnerRepository domain.GymOwnerRepository
}

func (h DeleteGymOwnerCommandHandler) Handle(command *DeleteGymOwnerCommand) (*DeleteGymOwnerCommandResponse, *application_specific.ApplicationException) {
	owner, err := h.GymOwnerRepository.FindByID(command.Id, command.Session.Session)
	if err != nil {
		return nil, err
	}
	if owner == nil {
		return nil, application_specific.NewNotFoundException("GYMS.OWNERS.NOT_FOUND", "Gym owner not found", map[string]string{
			"id": command.Id,
		})
	}

	err = owner.Delete(command.Session.User.Id)
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

	return &DeleteGymOwnerCommandResponse{
		Id: owner.State().Id,
	}, nil
}
