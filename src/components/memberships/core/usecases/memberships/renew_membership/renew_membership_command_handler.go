package renew_membership

import (
	"gym-management/src/components/memberships/core/domain"
	"gym-management/src/lib/primitives/application_specific"
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
