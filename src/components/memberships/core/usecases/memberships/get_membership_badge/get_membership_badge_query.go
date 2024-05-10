package get_membership_badge

import "gym-management/src/lib/primitives/application_specific"

type GetMembershipBadgeQuery struct {
	Id      string
	Session *application_specific.UserSession
}
