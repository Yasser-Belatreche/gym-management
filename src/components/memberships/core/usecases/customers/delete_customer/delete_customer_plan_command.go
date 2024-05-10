package delete_customer_plan

import (
	"gym-management/src/lib/primitives/application_specific"
)

type DeleteCustomerCommand struct {
	Id      string
	Session *application_specific.UserSession
}
