package get_user

import "gym-management-auth/src/lib/primitives/application_specific"

type GetUserQuery struct {
	Id      string
	Session *application_specific.UserSession
}
