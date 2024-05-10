package delete_plan

import (
	"gym-management/src/components/memberships/core/domain"
	"gym-management/src/lib/primitives/application_specific"
)

type DeletePlanCommandHandler struct {
	PlanRepository  domain.PlanRepository
	EventsPublisher domain.EventsPublisher
}

func (h *DeletePlanCommandHandler) Handle(command *DeletePlanCommand) (*DeletePlanCommandResponse, *application_specific.ApplicationException) {
	plan, err := h.PlanRepository.FindByID(command.Id, command.Session.Session)
	if err != nil {
		return nil, err
	}

	err = plan.Delete(command.Session.User.Id)
	if err != nil {
		return nil, err
	}

	err = h.PlanRepository.Update(plan, command.Session.Session)
	if err != nil {
		return nil, err
	}

	err = h.EventsPublisher.Publish(plan.PullEvents(), command.Session.Session)
	if err != nil {
		return nil, err
	}

	return &DeletePlanCommandResponse{
		Id: plan.State().Id,
	}, nil
}
