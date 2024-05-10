package get_plan

import "gym-management/src/lib/primitives/application_specific"

type GetPlansQuery struct {
	application_specific.PaginatedQuery

	Id       []string
	GymId    []string
	Featured *bool
	Deleted  bool
	Session  *application_specific.UserSession
}
