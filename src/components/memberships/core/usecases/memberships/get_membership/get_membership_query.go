package get_membership

import "gym-management/src/lib/primitives/application_specific"

type GetMembershipQuery struct {
	MembershipId   string
	MembershipCode string
	Session        *application_specific.UserSession
}
