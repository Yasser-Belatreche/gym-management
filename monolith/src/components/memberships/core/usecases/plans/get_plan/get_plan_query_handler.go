package get_plan

import (
	"errors"
	"gorm.io/gorm"
	plans2 "gym-management/src/components/memberships/core/usecases/plans"
	"gym-management/src/lib"
	"gym-management/src/lib/persistence/psql/gorm/models"
	"gym-management/src/lib/primitives/application_specific"
)

type GetPlanQueryHandler struct{}

func (h *GetPlanQueryHandler) Handle(query *GetPlanQuery) (*GetPlanQueryResponse, *application_specific.ApplicationException) {
	db := lib.GormDB(query.Session)

	var plan models.Plan

	dbQuery := db.Model(&models.Plan{})

	if err := dbQuery.Where("id = ?", query.Id).First(&plan).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, application_specific.NewNotFoundException("PLANS.NOT_FOUND", "plan not found", map[string]string{
				"id": query.Id,
			})

		}

		return nil, application_specific.NewUnknownException("PLANS.FAILED_TO_GET_PLAN", err.Error(), nil)
	}

	response := GetPlanQueryResponse(
		plans2.PlanToReturn{
			Id:             plan.Id,
			Name:           plan.Name,
			Featured:       plan.Featured,
			SessionPerWeek: plan.SessionsPerWeek,
			WithCoach:      plan.WithCoach,
			MonthlyPrice:   plan.MonthlyPrice,
			GymId:          plan.GymId,
			CreatedBy:      plan.CreatedBy,
			UpdatedBy:      plan.UpdatedBy,
			CreatedAt:      plan.CreatedAt,
			UpdatedAt:      plan.UpdatedAt,
			DeletedBy:      plan.DeletedBy,
			DeletedAt:      plan.DeletedAt,
		})

	return &response, nil
}
