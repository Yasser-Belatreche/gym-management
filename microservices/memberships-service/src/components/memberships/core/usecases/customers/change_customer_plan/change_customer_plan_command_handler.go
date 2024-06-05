package change_customer_plan

import (
	"gym-management-memberships/src/components/memberships/core/domain"
	"gym-management-memberships/src/lib/primitives/application_specific"
)

type ChangeCustomerPlanCommandHandler struct {
	CustomerRepository   domain.CustomerRepository
	PlanRepository       domain.PlanRepository
	MembershipRepository domain.MembershipRepository
	EventsPublisher      domain.EventsPublisher
}

func (h *ChangeCustomerPlanCommandHandler) Handle(command *ChangeCustomerPlanCommand) (*ChangeCustomerPlanCommandResponse, *application_specific.ApplicationException) {
	customer, err := h.CustomerRepository.FindByID(command.CustomerId, command.Session.Session)
	if err != nil {
		return nil, err
	}

	membership, err := h.MembershipRepository.FindLatestCustomerMembership(command.CustomerId, command.Session.Session)
	if err != nil {
		return nil, err
	}

	plan, err := h.PlanRepository.FindByID(command.PlanId, command.Session.Session)

	service, err := domain.NewCustomerMembershipService(customer, membership)
	if err != nil {
		return nil, err
	}

	newMembership, err := service.ChangeCustomerPlanTo(plan, command.EndDate, command.Session.User.Id)
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

	err = h.MembershipRepository.Create(newMembership, command.Session.Session)
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

	err = h.EventsPublisher.Publish(newMembership.PullEvents(), command.Session.Session)
	if err != nil {
		return nil, err
	}

	return &ChangeCustomerPlanCommandResponse{
		CustomerId:   customer.State().Id,
		MembershipId: newMembership.State().Id,
	}, nil
}
