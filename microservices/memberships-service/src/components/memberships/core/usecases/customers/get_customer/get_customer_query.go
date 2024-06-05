package get_customer

import "gym-management-memberships/src/lib/primitives/application_specific"

type GetCustomerQuery struct {
	Id      string
	Session *application_specific.UserSession
}
