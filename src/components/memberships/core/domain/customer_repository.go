package domain

import "gym-management/src/lib/primitives/application_specific"

type CustomerRepository interface {
	FindByID(id string) (*Customer, *application_specific.ApplicationException)

	Create(customer *Customer) *application_specific.ApplicationException

	Update(customer *Customer) *application_specific.ApplicationException
}
