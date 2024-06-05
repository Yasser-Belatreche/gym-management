package get_gyms

import (
	"gym-management-gyms/src/components/gyms/core/usecases/gyms"
	"gym-management-gyms/src/lib/primitives/application_specific"
)

type GetGymsQueryResponse = application_specific.PaginatedQueryResponse[gyms.GymToReturn]
