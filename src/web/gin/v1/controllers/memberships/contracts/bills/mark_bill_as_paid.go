package bills

import (
	"gym-management/src/web/gin/v1/controllers/memberships/contracts"
)

type MarkBillAsPaidUrl struct {
	contracts.MembershipsUrl
	BillId       string `uri:"billId" binding:"required"`
	MembershipId string `uri:"membershipId" binding:"required"`
}

type MarkBillAsPaidResponse struct {
	Id string `json:"id"`
}
