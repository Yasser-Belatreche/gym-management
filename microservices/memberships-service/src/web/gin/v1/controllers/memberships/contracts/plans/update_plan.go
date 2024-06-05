package plans

import (
	"gym-management/src/web/gin/v1/controllers/memberships/contracts"
)

type UpdatePlanUrl struct {
	contracts.MembershipsUrl
	PlanId string `uri:"planId" binding:"required"`
}

type UpdatePlanRequest struct {
	Name            string  `form:"name" json:"name" binding:"required"`
	Featured        *bool   `form:"featured" json:"featured" binding:"required"`
	SessionsPerWeek int     `form:"sessionsPerWeek" json:"sessionsPerWeek" binding:"required"`
	WithCoach       *bool   `form:"withCoach" json:"withCoach" binding:"required"`
	MonthlyPrice    float64 `form:"monthlyPrice" json:"monthlyPrice" binding:"required"`
}

type UpdatePlanResponse struct {
	Id string `json:"id"`
}
