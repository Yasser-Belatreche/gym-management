package get_membership

import "gym-management/src/lib/primitives/application_specific"

type GetMembershipQueryHandler struct{}

func (h *GetMembershipQueryHandler) Handle(query *GetMembershipQuery) (*GetMembershipQueryResponse, *application_specific.ApplicationException) {
	return nil, nil
}
