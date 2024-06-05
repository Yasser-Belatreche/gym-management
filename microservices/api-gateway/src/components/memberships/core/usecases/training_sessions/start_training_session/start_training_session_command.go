package start_training_session

import "gym-management/src/lib/primitives/application_specific"

type StartTrainingSessionCommand struct {
	MembershipId string
	Session      *application_specific.UserSession
}
