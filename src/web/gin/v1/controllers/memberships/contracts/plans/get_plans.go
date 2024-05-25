package plans

import (
	"gym-management/src/web/gin/v1/controllers/memberships/contracts"
	"gym-management/src/web/gin/v1/controllers/memberships/contracts/base"
	"gym-management/src/web/gin/v1/utils"
)

type GetPlansUrl struct {
	contracts.MembershipsUrl
	utils.HttpPaginatedRequest
	Id       []string `uri:"planId" binding:"omitempty,dive"`
	Featured *bool    `uri:"featured" binding:"omitempty"`
	Deleted  bool     `uri:"deleted" binding:"omitempty"`
}

type GetPlansResponse utils.HttpPaginatedResponse[base.Plan]
