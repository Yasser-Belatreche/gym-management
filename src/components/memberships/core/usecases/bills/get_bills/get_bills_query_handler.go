package get_bill

import "gym-management/src/lib/primitives/application_specific"

type GetBillsQueryHandler struct{}

func (h *GetBillsQueryHandler) Handle(query *GetBillsQuery) (*GetBillsQueryResponse, *application_specific.ApplicationException) {
	return nil, nil
}
