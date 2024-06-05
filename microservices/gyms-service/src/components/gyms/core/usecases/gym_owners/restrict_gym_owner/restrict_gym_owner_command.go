package restrict_gym_owner

import "gym-management-gyms/src/lib/primitives/application_specific"

type RestrictGymOwnerCommand struct {
	Id      string
	Session *application_specific.UserSession
}
