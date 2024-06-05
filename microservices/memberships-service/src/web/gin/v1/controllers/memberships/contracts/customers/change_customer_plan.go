package customers

import (
	"gym-management/src/web/gin/v1/controllers/memberships/contracts"
	"time"
)

type ChangeCustomerPlanUrl struct {
	contracts.MembershipsUrl
	CustomerId string `uri:"customerId" binding:"required"`
}

type ChangeCustomerPlanRequest struct {
	PlanId  string     `form:"planId" binding:"required"`
	EndDate *time.Time `form:"endDate" binding:"omitempty"`
}

type ChangeCustomerPlanResponse struct {
	CustomerId   string `json:"customerId"`
	MembershipId string `json:"membershipId"`
}
