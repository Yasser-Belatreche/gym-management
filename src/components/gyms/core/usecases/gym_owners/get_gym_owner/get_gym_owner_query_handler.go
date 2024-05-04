package get_gym_owner

import "gym-management/src/lib/primitives/application_specific"

type GetGymOwnerQueryHandler struct{}

func (h *GetGymOwnerQueryHandler) Handle(query *GetGymOwnerQuery) (*GetGymOwnerQueryResponse, *application_specific.ApplicationException) {
	return nil, nil
}
