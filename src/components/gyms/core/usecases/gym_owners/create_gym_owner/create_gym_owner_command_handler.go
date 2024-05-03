package create_gym_owner

import (
	"gym-management/src/components/gyms/core/domain"
	"gym-management/src/lib/primitives/application_specific"
)

type CreateGymOwnerCommandHandler struct {
	EmailService       domain.EmailService
	EventsPublisher    domain.EventsPublisher
	GymOwnerRepository domain.GymOwnerRepository
}

func (h CreateGymOwnerCommandHandler) Handle(command *CreateGymOwnerCommand) (*CreateGymOwnerCommandResponse, *application_specific.ApplicationException) {
	email, err := application_specific.NewEmail(command.Email)
	if err != nil {
		return nil, err
	}

	if h.EmailService.IsUsed(email) {
		return nil, application_specific.NewValidationException("GYMS.OWNERS.EMAIL_USED", "Email is already used", map[string]string{
			"email": command.Email,
		})
	}

	owner, err := domain.CreateGymOwner(command.Name, command.PhoneNumber, email, command.Password, command.Session.User.Id)
	if err != nil {
		return nil, err
	}

	err = h.GymOwnerRepository.Create(owner)
	if err != nil {
		return nil, err
	}

	err = h.EventsPublisher.Publish(owner.PullEvents())
	if err != nil {
		return nil, err
	}

	return &CreateGymOwnerCommandResponse{
		Id: owner.State().Id,
	}, nil
}
