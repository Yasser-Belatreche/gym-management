package start_training_session

import "gym-management/src/lib/primitives/application_specific"

type StartTrainingSessionCommand struct {
	MembershipCode string
	Session        *application_specific.UserSession
}
