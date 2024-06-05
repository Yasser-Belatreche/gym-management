package get_membership

import "gym-management-memberships/src/lib/primitives/application_specific"

type GetMembershipQuery struct {
	MembershipId string
	Session      *application_specific.UserSession
}
