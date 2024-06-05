package customers

import (
	"gym-management-memberships/src/web/gin/v1/controllers/memberships/contracts"
	"gym-management-memberships/src/web/gin/v1/controllers/memberships/contracts/base"
	"gym-management-memberships/src/web/gin/v1/utils"
)

type GetCustomersUrl struct {
	contracts.MembershipsUrl
}

type GetCustomersRequest struct {
	utils.HttpPaginatedRequest
	Id           []string `form:"customerId" json:"customerId" binding:"omitempty,dive"`
	MembershipId []string `form:"membershipId" json:"membershipId" binding:"omitempty,dive"`
	PlanId       []string `form:"planId" json:"planId" binding:"omitempty,dive"`
	Restricted   *bool    `form:"restricted" json:"restricted" binding:"omitempty"`
	Deleted      bool     `form:"deleted" json:"deleted" binding:"omitempty"`
}

type GetCustomersResponse utils.HttpPaginatedResponse[base.Customer]
