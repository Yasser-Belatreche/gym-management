package restrict_customer

import (
	"gym-management/src/components/memberships/core/domain"
	"gym-management/src/lib/primitives/application_specific"
)

type RestrictCustomerCommandHandler struct {
	CustomerRepository   domain.CustomerRepository
	MembershipRepository domain.MembershipRepository
	EventsPublisher      domain.EventsPublisher
}

func (h *RestrictCustomerCommandHandler) Handle(command *RestrictCustomerCommand) (*RestrictCustomerCommandResponse, *application_specific.ApplicationException) {
	customer, err := h.CustomerRepository.FindByID(command.Id, command.Session.Session)
	if err != nil {
		return nil, err
	}

	membership, err := h.MembershipRepository.FindLatestCustomerMembership(command.Id, command.Session.Session)
	if err != nil {
		return nil, err
	}

	service, err := domain.NewCustomerMembershipService(customer, membership)
	if err != nil {
		return nil, err
	}

	err = service.RestrictCustomer(command.Session.User.Id)
	if err != nil {
		return nil, err
	}

	err = h.CustomerRepository.Update(customer, command.Session.Session)
	if err != nil {
		return nil, err
	}

	err = h.MembershipRepository.Update(membership, command.Session.Session)
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

	return &RestrictCustomerCommandResponse{
		Id: customer.State().Id,
	}, nil
}
