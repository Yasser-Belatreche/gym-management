package create_customer

import (
	"gym-management-memberships/src/components/memberships/core/domain"
	"gym-management-memberships/src/lib/primitives/application_specific"
)

type CreateCustomerCommandHandler struct {
	CustomerRepository   domain.CustomerRepository
	PlanRepository       domain.PlanRepository
	MembershipRepository domain.MembershipRepository
	EmailsService        domain.EmailService
	EventsPublisher      domain.EventsPublisher
}

func (h *CreateCustomerCommandHandler) Handle(command *CreateCustomerCommand) (*CreateCustomerCommandResponse, *application_specific.ApplicationException) {
	plan, err := h.PlanRepository.FindByID(command.PlanId, command.Session.Session)
	if err != nil {
		return nil, err
	}

	email, err := application_specific.NewEmail(command.Email)
	if err != nil {
		return nil, err
	}

	if h.EmailsService.IsUsed(email, command.Session.Session) {
		return nil, application_specific.NewValidationException("CUSTOMER.EMAIL.USED", "Email is already used", map[string]string{
			"email": email.Value,
		})
	}

	gender, err := domain.GenderFrom(command.Gender)
	if err != nil {
		return nil, err
	}

	customer, membership, err := domain.CreateCustomer(
		command.FirstName,
		command.LastName,
		command.PhoneNumber,
		email,
		command.Password,
		command.BirthYear,
		gender,
		command.Session.User.Id,
		command.MembershipEndDate,
		plan)
	if err != nil {
		return nil, err
	}

	err = h.CustomerRepository.Create(customer, command.Session.Session)
	if err != nil {
		return nil, err
	}

	err = h.MembershipRepository.Create(membership, command.Session.Session)
	if err != nil {
		return nil, err
	}

	err = h.EventsPublisher.Publish(customer.PullEvents(), command.Session.Session)
	if err != nil {
		return nil, err
	}

	err = h.EventsPublisher.Publish(membership.PullEvents(), command.Session.Session)
	if err != nil {
		return nil, err
	}

	return &CreateCustomerCommandResponse{
		CustomerId:   customer.State().Id,
		MembershipId: membership.State().Id,
	}, nil
}
