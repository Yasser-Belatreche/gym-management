package update_gym

import "gym-management-gyms/src/lib/primitives/application_specific"

type UpdateGymCommand struct {
	Name    string
	Address string
	GymId   string
	OwnerId string
	Session *application_specific.UserSession
}
