package get_training_sessions

import (
	"gym-management-memberships/src/components/memberships/core/usecases/training_sessions"
	"gym-management-memberships/src/lib/primitives/application_specific"
)

type GetTrainingSessionsQueryResponse = application_specific.PaginatedQueryResponse[training_sessions.TrainingSessionToReturn]
