package plans

import (
	"gym-management/src/web/gin/v1/controllers/memberships/contracts"
	"gym-management/src/web/gin/v1/controllers/memberships/contracts/base"
)

type GetPlanUrl struct {
	contracts.MembershipsUrl
	Id string `uri:"planId" binding:"required"`
}

type GetPlanResponse base.Plan
