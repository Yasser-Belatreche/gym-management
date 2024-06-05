package get_gym_owners

import "gym-management-gyms/src/lib/primitives/application_specific"

type GetGymOwnersQuery struct {
	application_specific.PaginatedQuery

	Id         []string
	Search     string
	Restricted *bool
	Deleted    bool
	Session    *application_specific.UserSession
}
