package domain

import (
	"gym-management/src/components/memberships/core/domain/events"
	"gym-management/src/lib/primitives/application_specific"
	"gym-management/src/lib/primitives/generic"
	"strconv"
	"strings"
	"time"
)

type Plan struct {
	id              string
	name            string
	featured        bool
	sessionsPerWeek int
	withCoach       bool
	monthlyPrice    float64
	gymId           string
	createdBy       string
	updatedBy       string
	deletedBy       *string
	deletedAt       *time.Time
	events          []*events.MembershipEvent[interface{}]
}

type PlanState struct {
	Id              string
	Name            string
	Featured        bool
	SessionsPerWeek int
	WithCoach       bool
	MonthlyPrice    float64
	GymId           string
	CreatedBy       string
	UpdatedBy       string
	DeletedBy       *string
	DeletedAt       *time.Time
}

func CreatePlan(name string, featured bool, sessionsPerWeek int, withCoach bool, monthlyPrice float64, gymId string, createdBy string) (*Plan, *application_specific.ApplicationException) {
	name = strings.TrimSpace(name)

	if len(name) == 0 {
		return nil, application_specific.NewValidationException("MEMBERSHIPS.PLANS.INVALID_NAME", "Name cannot be empty", map[string]string{
			"name": name,
		})
	}

	if sessionsPerWeek <= 0 {
		return nil, application_specific.NewValidationException("MEMBERSHIPS.PLANS.INVALID_SESSIONS_PER_WEEK", "Sessions per week must be greater than 0", map[string]string{
			"sessionsPerWeek": strconv.Itoa(sessionsPerWeek),
		})
	}

	if monthlyPrice <= 0 {
		return nil, application_specific.NewValidationException("MEMBERSHIPS.PLANS.INVALID_MONTHLY_PRICE", "Monthly price must be greater than 0", map[string]string{
			"monthlyPrice": strconv.FormatFloat(monthlyPrice, 'f', -1, 64),
		})
	}

	plan := &Plan{
		id:              generic.GenerateUUID(),
		name:            name,
		featured:        featured,
		sessionsPerWeek: sessionsPerWeek,
		withCoach:       withCoach,
		monthlyPrice:    monthlyPrice,
		gymId:           gymId,
		createdBy:       createdBy,
		updatedBy:       createdBy,
		deletedBy:       nil,
		deletedAt:       nil,
		events:          make([]*events.MembershipEvent[interface{}], 0),
	}

	plan.pushEvents(NewPlanCreatedEvent(plan.State()))

	return plan, nil
}

func PlanFromState(state *PlanState) *Plan {
	return &Plan{
		id:              state.Id,
		name:            state.Name,
		featured:        state.Featured,
		sessionsPerWeek: state.SessionsPerWeek,
		withCoach:       state.WithCoach,
		monthlyPrice:    state.MonthlyPrice,
		gymId:           state.GymId,
		createdBy:       state.CreatedBy,
		updatedBy:       state.UpdatedBy,
		deletedBy:       state.DeletedBy,
		deletedAt:       state.DeletedAt,
		events:          make([]*events.MembershipEvent[interface{}], 0),
	}
}

func (p *Plan) Update(name string, featured bool, sessionsPerWeek int, withCoach bool, monthlyPrice float64, updatedBy string) *application_specific.ApplicationException {
	if p.deletedAt != nil {
		return application_specific.NewValidationException("MEMBERSHIPS.PLANS.DELETED", "Plan is deleted", map[string]string{
			"id": p.id,
		})
	}

	name = strings.TrimSpace(name)

	if len(name) == 0 {
		return application_specific.NewValidationException("MEMBERSHIPS.PLANS.INVALID_NAME", "Name cannot be empty", map[string]string{
			"name": name,
		})
	}

	if sessionsPerWeek <= 0 {
		return application_specific.NewValidationException("MEMBERSHIPS.PLANS.INVALID_SESSIONS_PER_WEEK", "Sessions per week must be greater than 0", map[string]string{
			"sessionsPerWeek": strconv.Itoa(sessionsPerWeek),
		})
	}

	if monthlyPrice <= 0 {
		return application_specific.NewValidationException("MEMBERSHIPS.PLANS.INVALID_MONTHLY_PRICE", "Monthly price must be greater than 0", map[string]string{
			"monthlyPrice": strconv.FormatFloat(monthlyPrice, 'f', -1, 64),
		})
	}

	p.name = name
	p.featured = featured
	p.sessionsPerWeek = sessionsPerWeek
	p.withCoach = withCoach
	p.monthlyPrice = monthlyPrice
	p.updatedBy = updatedBy

	p.pushEvents(NewPlanUpdatedEvent(p.State()))

	return nil
}

func (p *Plan) Delete(deletedBy string) *application_specific.ApplicationException {
	if p.deletedAt != nil {
		return application_specific.NewValidationException("MEMBERSHIPS.PLANS.DELETED", "Plan is deleted", map[string]string{
			"id": p.id,
		})
	}

	p.deletedBy = &deletedBy
	now := time.Now()
	p.deletedAt = &now

	p.pushEvents(NewPlanDeletedEvent(p.State()))

	return nil
}

func (p *Plan) pushEvents(event ...*events.MembershipEvent[interface{}]) {
	for _, e := range event {
		p.events = append(p.events, e)
	}
}

func (p *Plan) PullEvents() []*events.MembershipEvent[interface{}] {
	list := p.events
	p.events = nil
	return list
}

func (p *Plan) State() *PlanState {
	return &PlanState{
		Id:              p.id,
		Name:            p.name,
		Featured:        p.featured,
		SessionsPerWeek: p.sessionsPerWeek,
		WithCoach:       p.withCoach,
		MonthlyPrice:    p.monthlyPrice,
		GymId:           p.gymId,
		CreatedBy:       p.createdBy,
		UpdatedBy:       p.updatedBy,
		DeletedBy:       p.deletedBy,
		DeletedAt:       p.deletedAt,
	}
}
