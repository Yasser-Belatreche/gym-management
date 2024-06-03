package get_training_session

import (
	"errors"
	"gorm.io/gorm"
	"gym-management/src/components/memberships/core/usecases/training_sessions"
	"gym-management/src/lib"
	"gym-management/src/lib/primitives/application_specific"
)

type GetTrainingSessionQueryHandler struct{}

func (h *GetTrainingSessionQueryHandler) Handle(query *GetTrainingSessionQuery) (*GetTrainingSessionQueryResponse, *application_specific.ApplicationException) {
	db := lib.GormDB(query.Session)

	var session training_sessions.TrainingSessionToReturn

	err := db.Table("training_sessions as ts").
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
		Joins("INNER JOIN plans AS p ON p.id = m.plan_id").
		Where("ts.id = ?", query.Id).
		First(&session).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, application_specific.NewNotFoundException("TRAINING_SESSIONS.NOT_FOUND", "training session not found", map[string]string{
				"id": query.Id,
			})
		}

		return nil, application_specific.NewUnknownException("TRAINING_SESSIONS.FAILED_TO_GET_TRAINING_SESSION", err.Error(), nil)
	}

	response := GetTrainingSessionQueryResponse(session)

	return &response, nil
}
