package enable_gym

import "gym-management/src/lib/primitives/application_specific"

type EnableGymCommand struct {
	GymId   string
	OwnerId string
	Session *application_specific.UserSession
}
