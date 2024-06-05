package memberships

import (
	"gym-management-memberships/src/web/gin/v1/controllers/memberships/contracts"
	"time"
)

type RenewMembershipUrl struct {
	contracts.MembershipsUrl
	Id string `uri:"membershipId" binding:"required"`
}

type RenewMembershipRequest struct {
	EndDate *time.Time `json:"endDate" binding:"omitempty"`
}

type RenewMembershipResponse struct {
	Id string `json:"id"`
}
