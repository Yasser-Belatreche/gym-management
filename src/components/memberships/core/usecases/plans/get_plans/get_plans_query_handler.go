package get_plan

import "gym-management/src/lib/primitives/application_specific"

type GetPlansQueryHandler struct{}

func (h *GetPlansQueryHandler) Handle(query *GetPlansQuery) (*GetPlansQueryResponse, *application_specific.ApplicationException) {
	return nil, nil
}
