package get_memberships

import "gym-management/src/lib/primitives/application_specific"

type GetMembershipsQueryHandler struct{}

func (h *GetMembershipsQueryHandler) Handle(query *GetMembershipsQuery) (*GetMembershipsQueryResponse, *application_specific.ApplicationException) {
	return nil, nil
}
