package get_memberships

import "gym-management/src/lib/primitives/application_specific"

type GetMembershipsQuery struct {
	application_specific.PaginatedQuery

	Id         []string
	GymId      []string
	CustomerId []string
	PlanId     []string
	Enabled    *bool
	Session    *application_specific.UserSession
}
