package end_training_session

import "gym-management/src/lib/primitives/application_specific"

type EndTrainingSessionCommand struct {
	MembershipCode string
	Session        *application_specific.UserSession
}
