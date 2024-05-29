package create_customer

import (
	"gym-management/src/lib/primitives/application_specific"
	"time"
)

type CreateCustomerCommand struct {
	FirstName         string
	LastName          string
	Email             string
	PhoneNumber       string
	BirthYear         int
	Gender            string
	Password          string
	PlanId            string
	MembershipEndDate *time.Time
	Session           *application_specific.UserSession
}
