package get_gyms

import "gym-management/src/lib/primitives/application_specific"

type GetGymsQueryHandler struct {
}

func (h *GetGymsQueryHandler) Handle(query *GetGymsQuery) (*GetGymsQueryResponse, *application_specific.ApplicationException) {
	return nil, nil
}
