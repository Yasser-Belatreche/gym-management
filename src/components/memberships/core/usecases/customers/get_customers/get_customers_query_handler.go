package get_customers

import "gym-management/src/lib/primitives/application_specific"

type GetCustomersQueryHandler struct{}

func (h *GetCustomersQueryHandler) Handle(query *GetCustomersQuery) (*GetCustomersQueryResponse, *application_specific.ApplicationException) {
	return nil, nil
}
