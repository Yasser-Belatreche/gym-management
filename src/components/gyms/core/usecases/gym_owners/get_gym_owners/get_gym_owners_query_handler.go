package get_gym_owners

import "gym-management/src/lib/primitives/application_specific"

type GetGymOwnersQueryHandler struct{}

func (h *GetGymOwnersQueryHandler) Handle(query *GetGymOwnersQuery) (*GetGymOwnersQueryResponse, *application_specific.ApplicationException) {
	return nil, nil
}
