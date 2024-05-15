package get_gyms

import (
	"gym-management/src/components/gyms/core/usecases/gyms"
	"gym-management/src/lib/primitives/application_specific"
)

type GetGymsQueryResponse application_specific.PaginatedQueryResponse[gyms.GymToReturn]
