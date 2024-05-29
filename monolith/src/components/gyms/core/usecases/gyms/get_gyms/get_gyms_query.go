package get_gyms

import "gym-management/src/lib/primitives/application_specific"

type GetGymsQuery struct {
	application_specific.PaginatedQuery

	Id      []string
	OwnerId []string
	Search  string
	Enabled *bool
	Deleted bool
	Session *application_specific.UserSession
}
