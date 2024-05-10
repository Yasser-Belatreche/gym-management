package get_plan

import (
	"gym-management/src/components/memberships/core/usecases/plans"
	"gym-management/src/lib/primitives/application_specific"
)

type GetPlansQueryResponse application_specific.PaginatedQueryResponse[plans.PlanToReturn]
