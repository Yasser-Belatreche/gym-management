package customers

import (
	"gym-management/src/web/gin/v1/controllers/memberships/contracts"
	"time"
)

type CreateCustomerUrl struct {
	contracts.MembershipsUrl
}

type CreateCustomerRequest struct {
	FirstName         string     `form:"firstName" json:"firstName" binding:"required"`
	LastName          string     `form:"lastName" json:"lastName" binding:"required"`
	Email             string     `form:"email" json:"email" binding:"required"`
	PhoneNumber       string     `form:"phoneNumber" json:"phoneNumber" binding:"required"`
	BirthYear         int        `form:"birthYear" json:"birthYear" binding:"required"`
	Gender            string     `form:"gender" json:"gender" binding:"required"`
	Password          string     `form:"password" json:"password" binding:"required"`
	PlanId            string     `form:"planId" json:"planId" binding:"required"`
	MembershipEndDate *time.Time `form:"membershipEndDate" json:"membershipEndDate" binding:"required"`
}

type CreateCustomerResponse struct {
	CustomerId   string `json:"customerId"`
	MembershipId string `json:"membershipId"`
}
