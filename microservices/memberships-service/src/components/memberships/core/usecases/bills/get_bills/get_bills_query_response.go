package get_bills

import (
	"gym-management-memberships/src/components/memberships/core/usecases/bills"
	"gym-management-memberships/src/lib/primitives/application_specific"
)

type GetBillsQueryResponse = application_specific.PaginatedQueryResponse[bills.BillToReturn]
