package bills

import (
	"gym-management/src/web/gin/v1/controllers/memberships/contracts"
	"gym-management/src/web/gin/v1/controllers/memberships/contracts/base"
	"gym-management/src/web/gin/v1/utils"
)

type GetBillsUrl struct {
	contracts.MembershipsUrl
	utils.HttpPaginatedRequest
	MembershipId string `uri:"membershipId" binding:"required"`

	BillId     []string `uri:"billId" binding:"omitempty,dive"`
	CustomerId []string `uri:"customerId" binding:"omitempty,dive"`
	PlanId     []string `uri:"planId" binding:"omitempty,dive"`
	Paid       *bool    `uri:"paid" binding:"omitempty"`
}

type GetBillsResponse utils.HttpPaginatedResponse[base.Bill]
