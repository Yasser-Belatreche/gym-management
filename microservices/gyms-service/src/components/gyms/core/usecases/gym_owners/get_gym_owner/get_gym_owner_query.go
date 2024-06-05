package get_gym_owner

import "gym-management-gyms/src/lib/primitives/application_specific"

type GetGymOwnerQuery struct {
	Id      string
	Session *application_specific.UserSession
}
