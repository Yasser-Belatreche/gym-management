package cancel_membership

import "gym-management-memberships/src/lib/primitives/application_specific"

type CancelMembershipCommand struct {
	MembershipId string
	Session      *application_specific.UserSession
}
