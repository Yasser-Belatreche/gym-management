package training_sessions

import (
	"gym-management-memberships/src/web/gin/v1/controllers/memberships/contracts"
)

type EndTrainingSessionUrl struct {
	contracts.MembershipsUrl
	Id           string `uri:"sessionId" binding:"required"`
	MembershipId string `uri:"membershipId" binding:"required"`
}

type EndTrainingSessionResponse struct {
	Id string `json:"id"`
}
