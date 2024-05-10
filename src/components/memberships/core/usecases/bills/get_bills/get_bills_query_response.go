package get_bill

import (
	"gym-management/src/components/memberships/core/usecases/bills"
	"gym-management/src/lib/primitives/application_specific"
)

type GetBillsQueryResponse application_specific.PaginatedQueryResponse[bills.BillToReturn]
