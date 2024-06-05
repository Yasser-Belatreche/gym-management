package memberships

import (
	"gym-management-memberships/src/web/gin/v1/controllers/memberships/contracts"
	"gym-management-memberships/src/web/gin/v1/controllers/memberships/contracts/base"
	"gym-management-memberships/src/web/gin/v1/utils"
)

type GetMembershipsUrl struct {
	contracts.MembershipsUrl
	utils.HttpPaginatedRequest
	Id         []string `uri:"membershipId" binding:"omitempty,dive"`
	CustomerId []string `uri:"customerId" binding:"omitempty,dive"`
	PlanId     []string `uri:"planId" binding:"omitempty,dive"`
	Enabled    *bool    `uri:"enabled" binding:"omitempty"`
}

type GetMembershipsResponse utils.HttpPaginatedResponse[base.Membership]
