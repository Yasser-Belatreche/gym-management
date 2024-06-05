package restrict_customer

import (
	"gym-management/src/lib/primitives/application_specific"
)

type RestrictCustomerCommand struct {
	Id      string
	Session *application_specific.UserSession
}
