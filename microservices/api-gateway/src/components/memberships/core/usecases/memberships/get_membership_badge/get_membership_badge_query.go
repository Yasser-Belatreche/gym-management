package get_membership_badge

import "gym-management/src/lib/primitives/application_specific"

type GetMembershipBadgeQuery struct {
	MembershipId string
	Session      *application_specific.UserSession
}
