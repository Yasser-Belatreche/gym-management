package create_admin

import "gym-management-auth/src/lib/primitives/application_specific"

type CreateAdminCommand struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
	Phone     string
	Session   *application_specific.Session
}
