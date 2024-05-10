package start_training_session

import (
	"gym-management/src/components/memberships/core/domain"
	"gym-management/src/lib/primitives/application_specific"
)

type StartTrainingSessionCommandHandler struct {
	MembershipRepository domain.MembershipRepository
	EventsPublisher      domain.EventsPublisher
}

func (h *StartTrainingSessionCommandHandler) Handle(command *StartTrainingSessionCommand) (*StartTrainingSessionCommandResponse, *application_specific.ApplicationException) {
	membership, err := h.MembershipRepository.FindByCode(command.MembershipCode, command.Session.Session)
	if err != nil {
		return nil, err
	}

	trainingSession, err := membership.StartTrainingSession()
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

	if membership.IsDisabled() {
		return nil, application_specific.NewValidationException("MEMBERSHIPS.DISABLED", "Membership is disabled", map[string]string{
			"membershipId": membership.State().Id,
			"reason":       *membership.State().DisabledFor,
		})
	}

	return &StartTrainingSessionCommandResponse{
		Id: trainingSession.State().Id,
	}, nil
}
