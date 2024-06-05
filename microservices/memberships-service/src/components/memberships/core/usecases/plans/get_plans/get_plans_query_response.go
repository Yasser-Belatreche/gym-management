package get_plans

import (
	"gym-management-memberships/src/components/memberships/core/usecases/plans"
	"gym-management-memberships/src/lib/primitives/application_specific"
)

type GetPlansQueryResponse = application_specific.PaginatedQueryResponse[plans.PlanToReturn]
