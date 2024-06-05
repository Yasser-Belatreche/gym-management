package bills

import (
	"gym-management-memberships/src/web/gin/v1/controllers/memberships/contracts"
	"gym-management-memberships/src/web/gin/v1/controllers/memberships/contracts/base"
)

type GetBillUrl struct {
	contracts.MembershipsUrl
	Id           string `uri:"billId" binding:"required"`
	MembershipId string `uri:"membershipId" binding:"required"`
}

type GetBillResponse base.Bill
