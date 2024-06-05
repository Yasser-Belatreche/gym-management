package get_customers

import (
	"gym-management-memberships/src/components/memberships/core/usecases/customers"
	"gym-management-memberships/src/lib/primitives/application_specific"
)

type GetCustomersQueryResponse = application_specific.PaginatedQueryResponse[customers.CustomerToReturn]
