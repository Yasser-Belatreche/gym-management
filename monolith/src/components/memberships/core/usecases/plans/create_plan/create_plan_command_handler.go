package create_plan

import (
	"gym-management/src/components/memberships/core/domain"
	"gym-management/src/lib/primitives/application_specific"
)

type CreatePlanCommandHandler struct {
	PlanRepository  domain.PlanRepository
	EventsPublisher domain.EventsPublisher
}

func (h *CreatePlanCommandHandler) Handle(command *CreatePlanCommand) (*CreatePlanCommandResponse, *application_specific.ApplicationException) {
	plan, err := domain.CreatePlan(
		command.Name,
		command.Featured,
		command.SessionsPerWeek,
		command.WithCoach,
		command.MonthlyPrice,
		command.GymId,
		command.Session.User.Id)

	if err != nil {
		return nil, err
	}

	err = h.PlanRepository.Create(plan, command.Session.Session)
	if err != nil {
		return nil, err
	}

	err = h.EventsPublisher.Publish(plan.PullEvents(), command.Session.Session)
	if err != nil {
		return nil, err
	}

	return &CreatePlanCommandResponse{
		Id: plan.State().Id,
	}, nil
}
