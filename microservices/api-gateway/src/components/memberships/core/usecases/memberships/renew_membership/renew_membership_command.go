package renew_membership

import (
	"gym-management/src/lib/primitives/application_specific"
	"time"
)

type RenewMembershipCommand struct {
	MembershipId string
	EndDate      *time.Time
	Session      *application_specific.UserSession
}
