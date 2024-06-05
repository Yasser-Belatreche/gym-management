package update_plan

import (
	"gym-management/src/components/memberships/core/domain"
	"gym-management/src/lib/primitives/application_specific"
)

type UpdatePlanCommandHandler struct {
	PlanRepository  domain.PlanRepository
	EventsPublisher domain.EventsPublisher
}

func (h *UpdatePlanCommandHandler) Handle(command *UpdatePlanCommand) (*UpdatePlanCommandResponse, *application_specific.ApplicationException) {
	plan, err := h.PlanRepository.FindByID(command.Id, command.Session.Session)
	if err != nil {
		return nil, err
	}

	err = plan.Update(
		command.Name,
		command.Featured,
		command.SessionsPerWeek,
		command.WithCoach,
		command.MonthlyPrice,
		command.Session.User.Id)
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

	return &UpdatePlanCommandResponse{
		Id: plan.State().Id,
	}, nil
}
