package domain

import "gym-management-gyms/src/lib/primitives/application_specific"

type GymOwnerRepository interface {
	FindByID(id string, session *application_specific.Session) (*GymOwner, *application_specific.ApplicationException)

	Create(owner *GymOwner, session *application_specific.Session) *application_specific.ApplicationException

	Update(owner *GymOwner, session *application_specific.Session) *application_specific.ApplicationException
}
