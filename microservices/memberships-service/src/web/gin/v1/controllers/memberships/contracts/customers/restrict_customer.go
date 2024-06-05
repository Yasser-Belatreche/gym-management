package customers

import (
	"gym-management-memberships/src/web/gin/v1/controllers/memberships/contracts"
)

type RestrictCustomerUrl struct {
	contracts.MembershipsUrl
	Id string `uri:"customerId" binding:"required"`
}

type RestrictCustomerResponse struct {
	Id string `json:"id"`
}
