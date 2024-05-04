package domain

import "gym-management/src/lib/primitives/application_specific"

type UserRepository interface {
	FindByUsername(username Username) (*User, *application_specific.ApplicationException)

	FindByID(id string) (*User, *application_specific.ApplicationException)

	Create(user *User) *application_specific.ApplicationException

	Update(user *User) *application_specific.ApplicationException
}
