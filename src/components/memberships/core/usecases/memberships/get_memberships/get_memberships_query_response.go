package get_memberships

import (
	"gym-management/src/components/memberships/core/usecases/memberships"
	"gym-management/src/lib/primitives/application_specific"
)

type GetMembershipsQueryResponse application_specific.PaginatedQueryResponse[memberships.MembershipToReturn]
