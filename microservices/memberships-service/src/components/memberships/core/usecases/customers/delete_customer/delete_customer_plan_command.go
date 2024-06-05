package delete_customer

import (
	"gym-management-memberships/src/lib/primitives/application_specific"
)

type DeleteCustomerCommand struct {
	Id      string
	Session *application_specific.UserSession
}
