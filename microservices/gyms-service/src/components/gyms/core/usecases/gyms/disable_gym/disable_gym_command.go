package disable_gym

import "gym-management-gyms/src/lib/primitives/application_specific"

type DisableGymCommand struct {
	GymId   string
	OwnerId string
	Session *application_specific.UserSession
}
