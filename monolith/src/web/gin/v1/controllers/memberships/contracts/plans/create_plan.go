package plans

import (
	"gym-management/src/web/gin/v1/controllers/memberships/contracts"
)

type CreatePlanUrl struct {
	contracts.MembershipsUrl
}

type CreatePlanRequest struct {
	Name            string  `form:"name" binding:"required"`
	Featured        bool    `form:"featured" binding:"required"`
	SessionsPerWeek int     `form:"sessionsPerWeek" binding:"required"`
	WithCoach       bool    `form:"withCoach" binding:"required"`
	MonthlyPrice    float64 `form:"monthlyPrice" binding:"required"`
}

type CreatePlanResponse struct {
	Id string `json:"id"`
}
