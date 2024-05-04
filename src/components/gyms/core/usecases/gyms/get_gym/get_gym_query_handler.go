package get_gym

import "gym-management/src/lib/primitives/application_specific"

type GetGymQueryHandler struct {
}

func (h GetGymQueryHandler) Handle(query *GetGymQuery) (*GetGymQueryResponse, *application_specific.ApplicationException) {
	return nil, nil
}
