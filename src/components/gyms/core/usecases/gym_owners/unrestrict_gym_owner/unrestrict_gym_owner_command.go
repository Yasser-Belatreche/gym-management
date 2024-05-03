package unrestrict_gym_owner

import "gym-management/src/lib/primitives/application_specific"

type UnrestrictGymOwnerCommand struct {
	Id      string
	Session *application_specific.UserSession
}
