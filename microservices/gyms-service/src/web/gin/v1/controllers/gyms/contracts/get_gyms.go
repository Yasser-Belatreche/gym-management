package contracts

import (
	"gym-management-gyms/src/web/gin/v1/controllers/gyms/contracts/base"
	"gym-management-gyms/src/web/gin/v1/utils"
)

type GetGymsRequest struct {
	utils.HttpPaginatedRequest

	Id      []string `form:"id" binding:"omitempty,dive"`
	Search  string   `form:"search" binding:"omitempty"`
	Enabled *bool    `form:"enabled" binding:"omitempty"`
	Deleted bool     `form:"deleted" binding:"omitempty"`
}

type GetGymsUrl struct {
	OwnerId string `uri:"ownerId" binding:"required"`
}

type GetGymsResponse utils.HttpPaginatedResponse[base.Gym]
