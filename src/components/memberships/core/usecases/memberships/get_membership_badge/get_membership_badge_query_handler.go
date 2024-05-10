package get_membership_badge

import "gym-management/src/lib/primitives/application_specific"

type GetMembershipBadgeQueryHandler struct {
}

func (h *GetMembershipBadgeQueryHandler) Handle(query *GetMembershipBadgeQuery) (*GetMembershipBadgeQueryResponse, *application_specific.ApplicationException) {
	return nil, nil
}
