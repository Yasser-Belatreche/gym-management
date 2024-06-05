package get_bills

import "gym-management-memberships/src/lib/primitives/application_specific"

type GetBillsQuery struct {
	application_specific.PaginatedQuery
	Id           []string
	MembershipId []string
	CustomerId   []string
	GymId        []string
	PlanId       []string
	Paid         *bool
	Session      *application_specific.UserSession
}
