package customers

import (
	"gym-management/src/web/gin/v1/controllers/memberships/contracts"
	"gym-management/src/web/gin/v1/controllers/memberships/contracts/base"
)

type GetCustomerUrl struct {
	contracts.MembershipsUrl
	Id string `uri:"customerId" binding:"required"`
}

type GetCustomerResponse base.Customer
