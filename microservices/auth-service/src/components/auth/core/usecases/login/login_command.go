package login

import "gym-management-auth/src/lib/primitives/application_specific"

type LoginCommand struct {
	Username string
	Password string
	Session  *application_specific.Session
}
