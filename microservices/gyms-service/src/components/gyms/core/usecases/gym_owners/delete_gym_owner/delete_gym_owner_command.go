package delete_gym_owner

import "gym-management-gyms/src/lib/primitives/application_specific"

type DeleteGymOwnerCommand struct {
	Id      string
	Session *application_specific.UserSession
}
