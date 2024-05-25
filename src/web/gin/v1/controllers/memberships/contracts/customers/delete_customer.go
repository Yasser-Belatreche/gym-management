package customers

import (
	"gym-management/src/web/gin/v1/controllers/memberships/contracts"
)

type DeleteCustomerUrl struct {
	contracts.MembershipsUrl
	Id string `uri:"customerId" binding:"required"`
}

type DeleteCustomerResponse struct {
	Id string `json:"id"`
}
