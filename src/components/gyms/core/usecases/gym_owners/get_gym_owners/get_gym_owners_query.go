package get_gym_owners

import "gym-management/src/lib/primitives/application_specific"

type GetGymOwnersQuery struct {
	application_specific.PaginatedQuery

	Id         []string
	Search     string
	Restricted *bool
	Deleted    bool
	session    *application_specific.UserSession
}
