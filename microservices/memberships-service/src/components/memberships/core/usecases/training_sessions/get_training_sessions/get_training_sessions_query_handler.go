package get_training_sessions

import (
	"gym-management/src/components/memberships/core/usecases/training_sessions"
	"gym-management/src/lib"
	"gym-management/src/lib/primitives/application_specific"
)

type GetTrainingSessionsQueryHandler struct{}

func (h *GetTrainingSessionsQueryHandler) Handle(query *GetTrainingSessionsQuery) (*GetTrainingSessionsQueryResponse, *application_specific.ApplicationException) {
	db := lib.GormDB(query.Session)

	options := application_specific.NewPaginationOptions(&query.PaginatedQuery)

	var list []training_sessions.TrainingSessionToReturn

	dbQuery := db.Table("training_sessions as ts").
		Select(`
			ts.id AS id,
			ts.started_at AS started_at,
			ts.ended_at AS ended_at,

			c.id AS customer_id,
			c.first_name AS customer_first_name,
			c.last_name AS customer_last_name,

			p.gym_id AS gym_id
		`).
		Joins("INNER JOIN memberships AS m ON m.id = ts.membership_id").
		Joins("INNER JOIN customers AS c ON c.id = m.customer_id").
		Joins("INNER JOIN plans AS p ON p.id = m.plan_id")

	if len(query.Id) > 0 {
		dbQuery.Where("ts.id IN (?)", query.Id)
	}

	if len(query.MembershipId) > 0 {
		dbQuery.Where("ts.membership_id IN (?)", query.MembershipId)
	}

	if len(query.CustomerId) > 0 {
		dbQuery.Where("m.customer_id IN (?)", query.CustomerId)
	}

	if len(query.GymId) > 0 {
		dbQuery.Where("p.gym_id IN (?)", query.GymId)
	}

	if query.Ended != nil {
		dbQuery.Where("ts.ended_at IS NOT NULL")
	}

	err := dbQuery.
		Offset(options.Skip).
		Limit(options.PerPage).
		Order("ts.started_at DESC").
		Find(&list).
		Error
	if err != nil {
		return nil, application_specific.NewUnknownException("TRAINING_SESSIONS.FAILED_TO_GET_TRAINING_SESSIONS", err.Error(), nil)
	}

	var total int64
	err = dbQuery.Count(&total).Error
	if err != nil {
		return nil, application_specific.NewUnknownException("TRAINING_SESSIONS.FAILED_TO_GET_TRAINING_SESSIONS", err.Error(), nil)
	}

	response := GetTrainingSessionsQueryResponse(application_specific.NewPaginatedResponse(options, total, list, func(item training_sessions.TrainingSessionToReturn) training_sessions.TrainingSessionToReturn {
		return item
	}))

	return &response, nil
}
