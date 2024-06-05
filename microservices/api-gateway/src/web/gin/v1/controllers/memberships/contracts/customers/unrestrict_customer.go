package customers

import "gym-management/src/web/gin/v1/controllers/memberships/contracts"

type UnrestrictCustomerUrl struct {
	contracts.MembershipsUrl
	Id string `uri:"customerId" binding:"required"`
}

type UnrestrictCustomerResponse struct {
	Id string `json:"id"`
}
