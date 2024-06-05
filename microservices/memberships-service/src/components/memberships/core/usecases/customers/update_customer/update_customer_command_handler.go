package update_customer

import (
	"gym-management-memberships/src/components/memberships/core/domain"
	"gym-management-memberships/src/lib/primitives/application_specific"
)

type UpdateCustomerCommandHandler struct {
	CustomerRepository domain.CustomerRepository
	EmailsService      domain.EmailService
	EventsPublisher    domain.EventsPublisher
}

func (h *UpdateCustomerCommandHandler) Handle(command *UpdateCustomerCommand) (*UpdateCustomerCommandResponse, *application_specific.ApplicationException) {
	customer, err := h.CustomerRepository.FindByID(command.Id, command.Session.Session)
	if err != nil {
		return nil, err
	}

	email, err := application_specific.NewEmail(command.Email)
	if err != nil {
		return nil, err
	}

	if !customer.EmailIs(email) && h.EmailsService.IsUsed(email, command.Session.Session) {
		return nil, application_specific.NewValidationException("CUSTOMER.EMAIL.USED", "Email is already used", map[string]string{
			"email": email.Value,
		})
	}

	gender, err := domain.GenderFrom(command.Gender)
	if err != nil {
		return nil, err
	}

	err = customer.Update(
		command.FirstName,
		command.LastName,
		command.PhoneNumber,
		email,
		command.BirthYear,
		gender,
		command.NewPassword,
		command.Session.User.Id)
	if err != nil {
		return nil, err
	}

	err = h.CustomerRepository.Update(customer, command.Session.Session)
	if err != nil {
		return nil, err
	}

	err = h.EventsPublisher.Publish(customer.PullEvents(), command.Session.Session)
	if err != nil {
		return nil, err
	}

	return &UpdateCustomerCommandResponse{
		Id: customer.State().Id,
	}, nil
}
