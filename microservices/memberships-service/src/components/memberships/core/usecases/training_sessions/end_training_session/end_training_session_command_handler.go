package end_training_session

import (
	"gym-management/src/components/memberships/core/domain"
	"gym-management/src/lib/primitives/application_specific"
)

type EndTrainingSessionCommandHandler struct {
	MembershipRepository domain.MembershipRepository
	EventsPublisher      domain.EventsPublisher
}

func (h *EndTrainingSessionCommandHandler) Handle(command *EndTrainingSessionCommand) (*EndTrainingSessionCommandResponse, *application_specific.ApplicationException) {
	membership, err := h.MembershipRepository.FindByID(command.MembershipId, command.Session.Session)
	if err != nil {
		return nil, err
	}

	trainingSession, err := membership.EndTrainingSession()
	if err != nil {
		return nil, err
	}

	err = h.MembershipRepository.Update(membership, command.Session.Session)
	if err != nil {
		return nil, err
	}

	err = h.EventsPublisher.Publish(membership.PullEvents(), command.Session.Session)
	if err != nil {
		return nil, err
	}

	return &EndTrainingSessionCommandResponse{
		Id: trainingSession.State().Id,
	}, nil
}
