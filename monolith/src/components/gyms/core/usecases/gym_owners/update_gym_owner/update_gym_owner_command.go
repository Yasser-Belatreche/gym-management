package update_gym_owner

import "gym-management/src/lib/primitives/application_specific"

type UpdateGymOwnerCommand struct {
	Id          string
	Name        string
	PhoneNumber string
	Email       string
	NewPassword *string
	Session     *application_specific.UserSession
}
