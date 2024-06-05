package plans

import (
	"gym-management-memberships/src/web/gin/v1/controllers/memberships/contracts"
)

type CreatePlanUrl struct {
	contracts.MembershipsUrl
}

type CreatePlanRequest struct {
	Name            string  `form:"name" json:"name" binding:"required"`
	Featured        *bool   `form:"featured" json:"featured" binding:"required"`
	SessionsPerWeek int     `form:"sessionsPerWeek" json:"sessionsPerWeek" binding:"required"`
	WithCoach       *bool   `form:"withCoach" json:"withCoach" binding:"required"`
	MonthlyPrice    float64 `form:"monthlyPrice" json:"monthlyPrice" binding:"required"`
}

type CreatePlanResponse struct {
	Id string `json:"id"`
}
