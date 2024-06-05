package mark_bill_as_paid

import "gym-management/src/lib/primitives/application_specific"

type MarkBillAsPaidCommand struct {
	BillId       string
	MembershipId string
	Session      *application_specific.UserSession
}
