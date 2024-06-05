package memberships

import (
	"gym-management-memberships/src/web/gin/v1/controllers/memberships/contracts"
)

type GetMembershipBadgeUrl struct {
	contracts.MembershipsUrl
	Id string `uri:"membershipId" binding:"required"`
}

type GetMembershipBadgeResponse struct {
	Id   string `json:"id"`
	Code string `json:"code"`
}
