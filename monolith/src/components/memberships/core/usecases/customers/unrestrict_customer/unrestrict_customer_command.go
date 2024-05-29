package unrestrict_customer

import (
	"gym-management/src/lib/primitives/application_specific"
)

type UnrestrictCustomerCommand struct {
	Id      string
	Session *application_specific.UserSession
}
