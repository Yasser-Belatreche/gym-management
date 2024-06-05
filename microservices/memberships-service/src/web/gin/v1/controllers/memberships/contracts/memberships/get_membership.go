package memberships

import (
	"gym-management-memberships/src/web/gin/v1/controllers/memberships/contracts"
	"gym-management-memberships/src/web/gin/v1/controllers/memberships/contracts/base"
)

type GetMembershipUrl struct {
	contracts.MembershipsUrl
	Id string `uri:"membershipId" binding:"required"`
}

type GetMembershipResponse base.Membership
