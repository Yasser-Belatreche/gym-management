package plans

import (
	"gym-management-memberships/src/web/gin/v1/controllers/memberships/contracts"
	"gym-management-memberships/src/web/gin/v1/controllers/memberships/contracts/base"
	"gym-management-memberships/src/web/gin/v1/utils"
)

type GetPlansUrl struct {
	contracts.MembershipsUrl
}

type GetPlansRequest struct {
	utils.HttpPaginatedRequest
	Id       []string `form:"planId" json:"id" binding:"omitempty,dive"`
	Featured *bool    `form:"featured" json:"featured" binding:"omitempty"`
	Deleted  bool     `form:"deleted" json:"deleted" binding:"omitempty"`
}

type GetPlansResponse utils.HttpPaginatedResponse[base.Plan]
