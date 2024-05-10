package get_bill

import "gym-management/src/lib/primitives/application_specific"

type GetBillQueryHandler struct{}

func (h *GetBillQueryHandler) Handle(query *GetBillQuery) (*GetBillQueryResponse, *application_specific.ApplicationException) {
	return nil, nil
}
