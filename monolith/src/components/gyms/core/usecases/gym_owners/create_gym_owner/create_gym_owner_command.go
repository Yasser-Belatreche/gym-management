package create_gym_owner

import "gym-management/src/lib/primitives/application_specific"

type CreateGymOwnerCommand struct {
	Name        string
	PhoneNumber string
	Email       string
	Password    string
	Session     *application_specific.UserSession
}
