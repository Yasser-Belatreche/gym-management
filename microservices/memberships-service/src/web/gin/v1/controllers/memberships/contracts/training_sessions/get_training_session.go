package training_sessions

import (
	"gym-management-memberships/src/web/gin/v1/controllers/memberships/contracts"
	"gym-management-memberships/src/web/gin/v1/controllers/memberships/contracts/base"
)

type GetTrainingSessionUrl struct {
	contracts.MembershipsUrl
	SessionId    string `uri:"sessionId" binding:"required"`
	MembershipId string `uri:"membershipId" binding:"required"`
}

type GetTrainingSessionResponse base.TrainingSession
