package training_sessions

import (
	"gym-management-memberships/src/web/gin/v1/controllers/memberships/contracts"
)

type StartTrainingSessionUrl struct {
	contracts.MembershipsUrl
	MembershipId string `uri:"membershipId" binding:"required"`
}

type StartTrainingSessionResponse struct {
	Id string `json:"id"`
}
