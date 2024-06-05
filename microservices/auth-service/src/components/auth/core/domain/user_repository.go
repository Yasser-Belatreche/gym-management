package domain

import "gym-management-auth/src/lib/primitives/application_specific"

type UserRepository interface {
	FindByUsername(username Username, session *application_specific.Session) (*User, *application_specific.ApplicationException)

	UsernameUsed(username Username, session *application_specific.Session) (bool, *application_specific.ApplicationException)

	FindByID(id string, session *application_specific.Session) (*User, *application_specific.ApplicationException)

	Create(user *User, session *application_specific.Session) *application_specific.ApplicationException

	Update(user *User, session *application_specific.Session) *application_specific.ApplicationException
}
