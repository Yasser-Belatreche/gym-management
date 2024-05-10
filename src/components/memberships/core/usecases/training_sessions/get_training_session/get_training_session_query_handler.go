package get_training_session

import "gym-management/src/lib/primitives/application_specific"

type GetTrainingSessionQueryHandler struct{}

func (h *GetTrainingSessionQueryHandler) Handle(query *GetTrainingSessionQuery) (*GetTrainingSessionQueryResponse, *application_specific.ApplicationException) {
	return nil, nil
}
