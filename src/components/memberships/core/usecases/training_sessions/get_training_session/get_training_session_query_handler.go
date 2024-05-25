package get_training_session

import (
	"errors"
	"gorm.io/gorm"
	"gym-management/src/components/memberships/core/usecases/training_sessions"
	"gym-management/src/lib"
	"gym-management/src/lib/persistence/psql/gorm/models"
	"gym-management/src/lib/primitives/application_specific"
)

type GetTrainingSessionQueryHandler struct{}

func (h *GetTrainingSessionQueryHandler) Handle(query *GetTrainingSessionQuery) (*GetTrainingSessionQueryResponse, *application_specific.ApplicationException) {
	db := lib.GormDB(query.Session)

	var session models.TrainingSession

	dbQuery := db.Model(&models.TrainingSession{})
	dbQuery = dbQuery.Joins("Membership").Select("")
	dbQuery = dbQuery.Joins("Membership.Plan").Select("plans.gym_id")
	dbQuery = dbQuery.Joins("Membership.Customer").Select("customers.id, customers.first_name, customers.last_name")

	if err := dbQuery.Where("id = ?", query.Id).First(&session).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, application_specific.NewNotFoundException("TRAINING_SESSIONS.NOT_FOUND", "training session not found", map[string]string{
				"id": query.Id,
			})
		}

		return nil, application_specific.NewUnknownException("TRAINING_SESSIONS.FAILED_TO_GET_TRAINING_SESSION", err.Error(), nil)
	}

	response := GetTrainingSessionQueryResponse(
		training_sessions.TrainingSessionToReturn{
			Id:        session.Id,
			StartedAt: session.StartedAt,
			EndedAt:   session.EndedAt,
			Customer: struct {
				Id        string
				FirstName string
				LastName  string
			}{
				Id:        session.Membership.Customer.Id,
				FirstName: session.Membership.Customer.FirstName,
				LastName:  session.Membership.Customer.LastName,
			},
			GymId: session.Membership.Plan.GymId,
		})

	return &response, nil
}
