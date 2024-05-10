package get_plan

import "gym-management/src/lib/primitives/application_specific"

type GetPlanQueryHandler struct{}

func (h *GetPlanQueryHandler) Handle(query *GetPlanQuery) (*GetPlanQueryResponse, *application_specific.ApplicationException) {
	return nil, nil
}
