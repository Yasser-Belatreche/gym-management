package delete_gym

import "gym-management/src/lib/primitives/application_specific"

type DeleteGymCommand struct {
	GymId   string
	OwnerId string
	Session *application_specific.UserSession
}
