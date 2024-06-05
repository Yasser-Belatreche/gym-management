package get_training_session

import "gym-management/src/lib/primitives/application_specific"

type GetTrainingSessionQuery struct {
	Id      string
	Session *application_specific.UserSession
}
