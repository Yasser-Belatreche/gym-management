package cancel_membership

import (
	"gym-management-memberships/src/components/memberships/core/domain"
	"gym-management-memberships/src/lib/primitives/application_specific"
)

type CancelMembershipCommandHandler struct {
	MembershipRepository domain.MembershipRepository
	EventPublisher       domain.EventsPublisher
}

func (h *CancelMembershipCommandHandler) Handle(command *CancelMembershipCommand) (*CancelMembershipCommandResponse, *application_specific.ApplicationException) {
	membership, err := h.MembershipRepository.FindByID(command.MembershipId, command.Session.Session)
	if err != nil {
		return nil, err
	}

	err = membership.Cancel(command.Session.User.Id)
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

	return &CancelMembershipCommandResponse{
		Id: membership.State().Id,
	}, nil
}
