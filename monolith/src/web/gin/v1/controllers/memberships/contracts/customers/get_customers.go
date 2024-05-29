package customers

import (
	"gym-management/src/web/gin/v1/controllers/memberships/contracts"
	"gym-management/src/web/gin/v1/controllers/memberships/contracts/base"
	"gym-management/src/web/gin/v1/utils"
)

type GetCustomersUrl struct {
	contracts.MembershipsUrl
	utils.HttpPaginatedRequest
	Id           []string `uri:"customerId" binding:"omitempty,dive"`
	MembershipId []string `uri:"membershipId" binding:"omitempty,dive"`
	PlanId       []string `uri:"planId" binding:"omitempty,dive"`
	Restricted   *bool    `uri:"restricted" binding:"omitempty"`
	Deleted      bool     `uri:"deleted" binding:"omitempty"`
}

type GetCustomersResponse utils.HttpPaginatedResponse[base.Customer]
