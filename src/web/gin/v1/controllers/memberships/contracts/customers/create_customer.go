package customers

import (
	"gym-management/src/web/gin/v1/controllers/memberships/contracts"
	"time"
)

type CreateCustomerUrl struct {
	contracts.MembershipsUrl
}

type CreateCustomerRequest struct {
	FirstName         string     `form:"firstName" binding:"required"`
	LastName          string     `form:"lastName" binding:"required"`
	Email             string     `form:"email" binding:"required"`
	PhoneNumber       string     `form:"phoneNumber" binding:"required"`
	BirthYear         int        `form:"birthYear" binding:"required"`
	Gender            string     `form:"gender" binding:"required"`
	Password          string     `form:"password" binding:"required"`
	PlanId            string     `form:"planId" binding:"required"`
	MembershipEndDate *time.Time `form:"membershipEndDate" binding:"required"`
}

type CreateCustomerResponse struct {
	CustomerId   string `json:"customerId"`
	MembershipId string `json:"membershipId"`
}
