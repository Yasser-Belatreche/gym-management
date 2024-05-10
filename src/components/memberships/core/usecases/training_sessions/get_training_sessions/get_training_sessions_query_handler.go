package get_training_sessions

import "gym-management/src/lib/primitives/application_specific"

type GetTrainingSessionsQueryHandler struct{}

func (h *GetTrainingSessionsQueryHandler) Handle(query *GetTrainingSessionsQuery) (*GetTrainingSessionsQueryResponse, *application_specific.ApplicationException) {
	return nil, nil
}
