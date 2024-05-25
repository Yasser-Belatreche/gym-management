package get_training_sessions

import (
	"gym-management/src/components/memberships/core/usecases/training_sessions"
	"gym-management/src/lib"
	"gym-management/src/lib/persistence/psql/gorm/models"
	"gym-management/src/lib/primitives/application_specific"
)

type GetTrainingSessionsQueryHandler struct{}

func (h *GetTrainingSessionsQueryHandler) Handle(query *GetTrainingSessionsQuery) (*GetTrainingSessionsQueryResponse, *application_specific.ApplicationException) {
	db := lib.GormDB(query.Session)

	options := application_specific.NewPaginationOptions(&query.PaginatedQuery)

	var trainingSessions []models.TrainingSession

	dbQuery := db.Model(&models.TrainingSession{})
	dbQuery = dbQuery.Joins("Membership").Select("")
	dbQuery = dbQuery.Joins("Membership.Plan").Select("plans.gym_id")
	dbQuery = dbQuery.Joins("Membership.Customer").Select("customers.id, customers.first_name, customers.last_name")

	if len(query.Id) > 0 {
		dbQuery = dbQuery.Where("id IN (?)", query.Id)
	}

	if len(query.MembershipId) > 0 {
		dbQuery = dbQuery.Where("membership_id IN (?)", query.MembershipId)
	}

	if len(query.CustomerId) > 0 {
		dbQuery = dbQuery.Where("memberships.customer_id IN (?)", query.CustomerId)
	}

	if len(query.GymId) > 0 {
		dbQuery = dbQuery.Where("plans.gym_id IN (?)", query.GymId)
	}

	if query.Ended != nil {
		dbQuery = dbQuery.Where("ended = ?", *query.Ended)
	}

	results := dbQuery.Offset(options.Skip).Limit(options.PerPage).Order("started_at DESC").Find(&trainingSessions)
	if results.Error != nil {
		return nil, application_specific.NewUnknownException("TRAINING_SESSIONS.FAILED_TO_GET_TRAINING_SESSIONS", results.Error.Error(), nil)
	}

	var total int64
	results = dbQuery.Count(&total)
	if results.Error != nil {
		return nil, application_specific.NewUnknownException("TRAINING_SESSIONS.FAILED_TO_GET_TRAINING_SESSIONS", results.Error.Error(), nil)
	}

	response := GetTrainingSessionsQueryResponse(application_specific.NewPaginatedResponse(options, total, trainingSessions, func(item models.TrainingSession) training_sessions.TrainingSessionToReturn {
		return training_sessions.TrainingSessionToReturn{
			Id:        item.Id,
			StartedAt: item.StartedAt,
			EndedAt:   item.EndedAt,
			Customer: struct {
				Id        string
				FirstName string
				LastName  string
			}{
				Id:        item.Membership.Customer.Id,
				FirstName: item.Membership.Customer.FirstName,
				LastName:  item.Membership.Customer.LastName,
			},
			GymId: item.Membership.Plan.GymId,
		}
	}))

	return &response, nil
}
