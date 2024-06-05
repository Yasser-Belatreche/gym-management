package domain

import "gym-management-memberships/src/lib/primitives/application_specific"

type CustomerRepository interface {
	FindByID(id string, session *application_specific.Session) (*Customer, *application_specific.ApplicationException)

	Create(customer *Customer, session *application_specific.Session) *application_specific.ApplicationException

	Update(customer *Customer, session *application_specific.Session) *application_specific.ApplicationException
}
