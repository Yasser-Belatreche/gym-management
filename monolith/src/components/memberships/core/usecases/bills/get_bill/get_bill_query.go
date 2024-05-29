package get_bill

import "gym-management/src/lib/primitives/application_specific"

type GetBillQuery struct {
	BillId  string
	Session *application_specific.UserSession
}
