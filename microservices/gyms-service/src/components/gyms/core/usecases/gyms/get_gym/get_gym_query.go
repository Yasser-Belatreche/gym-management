package get_gym

import "gym-management-gyms/src/lib/primitives/application_specific"

type GetGymQuery struct {
	Id      string
	Session *application_specific.UserSession
}
