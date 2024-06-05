package renew_membership

import (
	"gym-management-memberships/src/components/memberships/core/domain"
	"gym-management-memberships/src/lib/primitives/application_specific"
)

type RenewMembershipCommandHandler struct {
	MembershipRepository domain.MembershipRepository
	EventPublisher       domain.EventsPublisher
}

func (h *RenewMembershipCommandHandler) Handle(command *RenewMembershipCommand) (*RenewMembershipCommandResponse, *application_specific.ApplicationException) {
	membership, err := h.MembershipRepository.FindByID(command.MembershipId, command.Session.Session)
	if err != nil {
		return nil, err
	}

	lastMembership, err := h.MembershipRepository.FindLatestCustomerMembership(membership.State().CustomerId, command.Session.Session)
	if err != nil {
		return nil, err
	}

	if !lastMembership.Equals(membership) {
		return nil, application_specific.NewValidationException("MEMBERSHIPS.NOT_LATEST", "Membership is not the latest", map[string]string{
			"membershipId": membership.State().Id,
		})
	}

	err = membership.Renew(command.EndDate, command.Session.User.Id)
	if err != nil {
		return nil, err
	}

	err = h.MembershipRepository.Update(membership, command.Session.Session)
	if err != nil {
		return nil, err
	}

	err = h.EventPublisher.Publish(membership.PullEvents(), command.Session.Session)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
