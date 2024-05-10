package get_bill

import "gym-management/src/lib/primitives/application_specific"

type GetBillsQuery struct {
	application_specific.PaginatedQuery
	Id           []string
	MembershipId []string
	CustomerId   []string
	GymId        []string
	Paid         *bool
	Session      *application_specific.UserSession
}
