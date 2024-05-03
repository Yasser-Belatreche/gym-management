package domain

import "gym-management/src/lib/primitives/application_specific"

type GymOwnerRepository interface {
	FindByID(id string) (*GymOwner, *application_specific.ApplicationException)

	Create(owner *GymOwner) *application_specific.ApplicationException

	Update(owner *GymOwner) *application_specific.ApplicationException
}
