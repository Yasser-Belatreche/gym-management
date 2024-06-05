package get_training_session

import "gym-management-memberships/src/lib/primitives/application_specific"

type GetTrainingSessionQuery struct {
	Id      string
	Session *application_specific.UserSession
}
