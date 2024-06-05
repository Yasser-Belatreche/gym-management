package infra

import (
	"errors"
	"gorm.io/gorm"
	"gym-management/src/components/memberships/core/domain"
	"gym-management/src/lib"
	"gym-management/src/lib/persistence/psql/gorm/models"
	"gym-management/src/lib/primitives/application_specific"
)

type GormPlanRepository struct{}

func (g *GormPlanRepository) FindByID(id string, session *application_specific.Session) (*domain.Plan, *application_specific.ApplicationException) {
	db := lib.GormDB(session)

	var plan models.Plan

	result := db.Set("gorm:query_option", "FOR UPDATE").First(&plan, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, application_specific.NewNotFoundException("MEMBERSHIPS.PLANS.NOT_FOUND", "Plan not found", map[string]string{
				"id": id,
			})
		}

		return nil, application_specific.NewUnknownException("MEMBERSHIPS.INFRA.FAILED_TO_FIND_PLAN", result.Error.Error(), map[string]string{})
	}

	domainPlan := planToDomain(&plan)

	return domainPlan, nil
}

func (g *GormPlanRepository) Create(plan *domain.Plan, session *application_specific.Session) *application_specific.ApplicationException {
	db := lib.GormDB(session)

	planModel := planToDB(plan)

	result := db.Create(planModel)
	if result.Error != nil {
		return application_specific.NewUnknownException("MEMBERSHIPS.INFRA.FAILED_TO_CREATE_PLAN", result.Error.Error(), map[string]string{})
	}

	return nil
}

func (g *GormPlanRepository) Update(plan *domain.Plan, session *application_specific.Session) *application_specific.ApplicationException {
	db := lib.GormDB(session)

	planModel := planToDB(plan)

	result := db.Save(planModel)
	if result.Error != nil {
		return application_specific.NewUnknownException("MEMBERSHIPS.INFRA.FAILED_TO_UPDATE_PLAN", result.Error.Error(), map[string]string{
			"id": planModel.Id,
		})
	}

	return nil
}

func planToDB(plan *domain.Plan) *models.Plan {
	state := plan.State()

	return &models.Plan{
		Id:              state.Id,
		Name:            state.Name,
		Featured:        state.Featured,
		SessionsPerWeek: state.SessionsPerWeek,
		WithCoach:       state.WithCoach,
		MonthlyPrice:    state.MonthlyPrice,
		GymId:           state.GymId,
		CreatedBy:       state.CreatedBy,
		UpdatedBy:       state.UpdatedBy,
		DeletedBy:       state.DeletedBy,
		DeletedAt:       state.DeletedAt,
	}
}

func planToDomain(plan *models.Plan) *domain.Plan {
	state := &domain.PlanState{
		Id:              plan.Id,
		Name:            plan.Name,
		Featured:        plan.Featured,
		SessionsPerWeek: plan.SessionsPerWeek,
		WithCoach:       plan.WithCoach,
		MonthlyPrice:    plan.MonthlyPrice,
		GymId:           plan.GymId,
		CreatedBy:       plan.CreatedBy,
		UpdatedBy:       plan.UpdatedBy,
		DeletedBy:       plan.DeletedBy,
		DeletedAt:       plan.DeletedAt,
	}

	return domain.PlanFromState(state)
}
