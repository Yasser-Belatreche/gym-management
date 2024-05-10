package mark_bill_as_paid

import (
	"gym-management/src/components/memberships/core/domain"
	"gym-management/src/lib/primitives/application_specific"
)

type MarkBillAsPaidCommandHandler struct {
	MembershipRepository domain.MembershipRepository
	EventPublisher       domain.EventsPublisher
}

func (h *MarkBillAsPaidCommandHandler) Handle(command *MarkBillAsPaidCommand) (*MarkBillAsPaidCommandResponse, *application_specific.ApplicationException) {
	membership, err := h.MembershipRepository.FindByID(command.MembershipId, command.Session.Session)
	if err != nil {
		return nil, err
	}

	err = membership.PayBill(command.BillId)
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
