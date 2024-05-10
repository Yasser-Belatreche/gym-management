package get_training_sessions

import "gym-management/src/lib/primitives/application_specific"

type GetTrainingSessionsQuery struct {
	Id           []string
	CustomerId   []string
	MembershipId []string
	GymId        []string
	Ended        *bool
	Session      *application_specific.UserSession
}
