package get_gym_owners

import (
	"gym-management/src/components/gyms/core/usecases/gym_owners"
	"gym-management/src/lib/primitives/application_specific"
)

type GetGymOwnersQueryResponse = application_specific.PaginatedQueryResponse[gym_owners.GymOwnerToReturn]
