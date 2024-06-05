package update_customer

import (
	"gym-management-memberships/src/lib/primitives/application_specific"
)

type UpdateCustomerCommand struct {
	Id          string
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	BirthYear   int
	Gender      string
	NewPassword *string
	Session     *application_specific.UserSession
}
