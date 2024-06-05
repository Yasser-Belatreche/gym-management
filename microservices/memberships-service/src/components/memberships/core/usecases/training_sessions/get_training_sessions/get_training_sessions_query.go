package get_training_sessions

import "gym-management-memberships/src/lib/primitives/application_specific"

type GetTrainingSessionsQuery struct {
	application_specific.PaginatedQuery
	Id           []string
	CustomerId   []string
	MembershipId []string
	GymId        []string
	Ended        *bool
	Session      *application_specific.UserSession
}
