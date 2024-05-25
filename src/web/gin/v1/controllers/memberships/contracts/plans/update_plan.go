package plans

import (
	"gym-management/src/web/gin/v1/controllers/memberships/contracts"
)

type UpdatePlanUrl struct {
	contracts.MembershipsUrl
	PlanId string `uri:"planId" binding:"required"`
}

type UpdatePlanRequest struct {
	Name            string  `form:"name" binding:"required"`
	Featured        bool    `form:"featured" binding:"required"`
	SessionsPerWeek int     `form:"sessionsPerWeek" binding:"required"`
	WithCoach       bool    `form:"withCoach" binding:"required"`
	MonthlyPrice    float64 `form:"monthlyPrice" binding:"required"`
}

type UpdatePlanResponse struct {
	Id string `json:"id"`
}
