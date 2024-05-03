package update_gym_owner

import (
	"gym-management/src/components/gyms/core/domain"
	"gym-management/src/lib/primitives/application_specific"
)

type CreateGymOwnerCommandHandler struct {
	EmailService       domain.EmailService
	EventsPublisher    domain.EventsPublisher
	GymOwnerRepository domain.GymOwnerRepository
}

func (h CreateGymOwnerCommandHandler) Handle(command *UpdateGymOwnerCommand) (*UpdateGymOwnerCommandResponse, *application_specific.ApplicationException) {
	owner, err := h.GymOwnerRepository.FindByID(command.Id)
	if err != nil {
		return nil, err
	}
	if owner == nil {
		return nil, application_specific.NewNotFoundException("GYMS.OWNERS.NOT_FOUND", "Gym owner not found", map[string]string{
			"id": command.Id,
		})
	}

	email, err := application_specific.NewEmail(command.Email)
	if err != nil {
		return nil, err
	}

	if !owner.EmailIs(email) && h.EmailService.IsUsed(email) {
		return nil, application_specific.NewValidationException("GYMS.OWNERS.EMAIL_USED", "Email is already used", map[string]string{
			"email": command.Email,
		})
	}

	err = owner.Update(command.Name, command.PhoneNumber, email, command.NewPassword, command.Session.User.Id)
	if err != nil {
		return nil, err
	}

	err = h.GymOwnerRepository.Update(owner)
	if err != nil {
		return nil, err
	}

	err = h.EventsPublisher.Publish(owner.PullEvents())
	if err != nil {
		return nil, err
	}

	return &UpdateGymOwnerCommandResponse{
		Id: owner.State().Id,
	}, nil
}
