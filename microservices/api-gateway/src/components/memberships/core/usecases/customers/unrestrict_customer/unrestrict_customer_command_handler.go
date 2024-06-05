package unrestrict_customer

import (
	"gym-management/src/components/memberships/core/domain"
	"gym-management/src/lib/primitives/application_specific"
)

type UnrestrictCustomerCommandHandler struct {
	CustomerRepository domain.CustomerRepository
	EventsPublisher    domain.EventsPublisher
}

func (h *UnrestrictCustomerCommandHandler) Handle(command *UnrestrictCustomerCommand) (*UnrestrictCustomerCommandResponse, *application_specific.ApplicationException) {
	customer, err := h.CustomerRepository.FindByID(command.Id, command.Session.Session)
	if err != nil {
		return nil, err
	}

	err = customer.Unrestrict(command.Session.User.Id)
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

	return &UnrestrictCustomerCommandResponse{
		Id: customer.State().Id,
	}, nil
}
