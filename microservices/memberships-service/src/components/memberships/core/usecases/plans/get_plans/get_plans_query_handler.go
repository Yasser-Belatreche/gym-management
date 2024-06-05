package get_plans

import (
	plans2 "gym-management/src/components/memberships/core/usecases/plans"
	"gym-management/src/lib"
	"gym-management/src/lib/persistence/psql/gorm/models"
	"gym-management/src/lib/primitives/application_specific"
)

type GetPlansQueryHandler struct{}

func (h *GetPlansQueryHandler) Handle(query *GetPlansQuery) (*GetPlansQueryResponse, *application_specific.ApplicationException) {
	db := lib.GormDB(query.Session)

	options := application_specific.NewPaginationOptions(&query.PaginatedQuery)

	var plans []models.Plan

	dbQuery := db.Table("plans")

	if len(query.Id) > 0 {
		dbQuery = dbQuery.Where("id IN (?)", query.Id)
	}

	if len(query.GymId) > 0 {
		dbQuery = dbQuery.Where("gym_id IN (?)", query.GymId)
	}

	if query.Featured != nil {
		dbQuery = dbQuery.Where("featured = ?", *query.Featured)
	}

	if query.Deleted {
		dbQuery = dbQuery.Where("deleted_at IS NOT NULL")
	} else {
		dbQuery = dbQuery.Where("deleted_at IS NULL")
	}

	result := dbQuery.Offset(options.Skip).Limit(options.PerPage).Order("updated_at DESC").Find(&plans)
	if result.Error != nil {
		return nil, application_specific.NewUnknownException("PLANS.FAILED_TO_GET_PLANS", result.Error.Error(), nil)
	}

	var total int64
	result = dbQuery.Count(&total)
	if result.Error != nil {
		return nil, application_specific.NewUnknownException("PLANS.FAILED_TO_GET_PLANS", result.Error.Error(), nil)
	}

	response := GetPlansQueryResponse(application_specific.NewPaginatedResponse(options, total, plans, func(item models.Plan) plans2.PlanToReturn {
		return plans2.PlanToReturn{
			Id:             item.Id,
			Name:           item.Name,
			Featured:       item.Featured,
			SessionPerWeek: item.SessionsPerWeek,
			WithCoach:      item.WithCoach,
			MonthlyPrice:   item.MonthlyPrice,
			GymId:          item.GymId,
			CreatedBy:      item.CreatedBy,
			UpdatedBy:      item.UpdatedBy,
			CreatedAt:      item.CreatedAt,
			UpdatedAt:      item.UpdatedAt,
			DeletedBy:      item.DeletedBy,
			DeletedAt:      item.DeletedAt,
		}
	}))

	return &response, nil
}
