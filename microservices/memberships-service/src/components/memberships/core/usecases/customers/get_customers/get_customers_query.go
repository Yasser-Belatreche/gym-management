package get_customers

import "gym-management-memberships/src/lib/primitives/application_specific"

type GetCustomersQuery struct {
	application_specific.PaginatedQuery

	Id           []string
	GymId        []string
	MembershipId []string
	PlanId       []string
	Restricted   *bool
	Deleted      bool
	Session      *application_specific.UserSession
}
