package memberships

import "gym-management/src/web/gin/v1/controllers/memberships/contracts"

type CancelMembershipUrl struct {
	contracts.MembershipsUrl
	Id string `uri:"membershipId" binding:"required"`
}

type CancelMembershipResponse struct {
	Id string `json:"id"`
}
