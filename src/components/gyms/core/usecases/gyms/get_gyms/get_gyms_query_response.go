package get_gyms

import (
	"gym-management/src/components/gyms/core/usecases/gyms"
	"gym-management/src/lib/primitives/application_specific"
)

type GetGymsQueryResponse struct {
	application_specific.PaginatedQueryResponse[gym_owners.GymToReturn]
}
