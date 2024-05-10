package get_customers

import (
	"gym-management/src/components/memberships/core/usecases/customers"
	"gym-management/src/lib/primitives/application_specific"
)

type GetCustomersQueryResponse application_specific.PaginatedQueryResponse[customers.CustomerToReturn]
