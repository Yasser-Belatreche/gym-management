package training_sessions

import (
	"gym-management/src/web/gin/v1/controllers/memberships/contracts"
	"gym-management/src/web/gin/v1/controllers/memberships/contracts/base"
	"gym-management/src/web/gin/v1/utils"
)

type GetTrainingSessionsUrl struct {
	contracts.MembershipsUrl
	utils.HttpPaginatedRequest
	MembershipId string `uri:"membershipId" binding:"required"`

	Id         []string `uri:"sessionId" binding:"omitempty,dive"`
	CustomerId []string `uri:"customerId" binding:"dive,omitempty"`
	Ended      *bool    `uri:"ended" binding:"omitempty"`
}

type GetTrainingSessionsResponse utils.HttpPaginatedResponse[base.TrainingSession]
