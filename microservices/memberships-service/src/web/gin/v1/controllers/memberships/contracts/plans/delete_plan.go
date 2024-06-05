package plans

import (
	"gym-management-memberships/src/web/gin/v1/controllers/memberships/contracts"
	"gym-management-memberships/src/web/gin/v1/controllers/memberships/contracts/base"
)

type DeletePlanUrl struct {
	contracts.MembershipsUrl
	Id string `uri:"planId" binding:"required"`
}

type DeletePlanResponse base.Plan
