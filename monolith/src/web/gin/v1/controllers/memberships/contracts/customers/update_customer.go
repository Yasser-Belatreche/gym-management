package customers

import (
	"gym-management/src/web/gin/v1/controllers/memberships/contracts"
)

type UpdateCustomerUrl struct {
	contracts.MembershipsUrl
	Id string `uri:"customerId" binding:"required"`
}

type UpdateCustomerRequest struct {
	FirstName   string  `form:"firstName" binding:"required"`
	LastName    string  `form:"lastName" binding:"required"`
	Email       string  `form:"email" binding:"required"`
	PhoneNumber string  `form:"phoneNumber" binding:"required"`
	BirthYear   int     `form:"birthYear" binding:"required"`
	Gender      string  `form:"gender" binding:"required"`
	NewPassword *string `form:"newPassword" binding:"omitempty"`
}

type UpdateCustomerResponse struct {
	Id string `json:"id"`
}
