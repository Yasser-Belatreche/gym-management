package get_customer

import "gym-management/src/lib/primitives/application_specific"

type GetCustomerQueryHandler struct{}

func (h *GetCustomerQueryHandler) Handle(query *GetCustomerQuery) (*GetCustomerQueryResponse, *application_specific.ApplicationException) {
	return nil, nil
}
