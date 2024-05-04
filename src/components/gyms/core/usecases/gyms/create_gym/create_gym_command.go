package create_gym

import "gym-management/src/lib/primitives/application_specific"

type CreateGymCommand struct {
	Name    string
	Address string
	OwnerId string
	Session *application_specific.UserSession
}
