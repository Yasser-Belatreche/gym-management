package contracts

import (
	"gym-management-gyms/src/web/gin/v1/controllers/gyms/contracts/base"
	"gym-management-gyms/src/web/gin/v1/utils"
)

type GetGymOwnersRequest struct {
	utils.HttpPaginatedRequest
	Id         []string `form:"id" binding:"omitempty,dive"`
	Search     string   `form:"search" binding:"omitempty"`
	Restricted *bool    `form:"restricted" binding:"omitempty"`
	Deleted    bool     `form:"deleted" binding:"omitempty"`
}

type GetGymOwnersResponse utils.HttpPaginatedResponse[base.GymOwner]
