package get_memberships

import (
	"gym-management-memberships/src/components/memberships/core/usecases/memberships"
	"gym-management-memberships/src/lib/primitives/application_specific"
)

type GetMembershipsQueryResponse = application_specific.PaginatedQueryResponse[memberships.MembershipToReturn]
