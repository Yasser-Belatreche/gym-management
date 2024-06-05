package base

import "time"

type Membership struct {
	Id              string             `json:"id"`
	StartDate       time.Time          `json:"startDate"`
	EndDate         *time.Time         `json:"endDate"`
	Enabled         bool               `json:"enabled"`
	DisabledFor     *string            `json:"disabledFor"`
	SessionsPerWeek int                `json:"sessionsPerWeek"`
	WithCoach       bool               `json:"withCoach"`
	MonthlyPrice    float64            `json:"monthlyPrice"`
	Customer        MembershipCustomer `json:"customer"`
	Plan            MembershipPlan     `json:"plan"`
	GymId           string             `json:"gymId"`
	CreatedBy       string             `json:"createdBy"`
	CreatedAt       time.Time          `json:"createdAt"`
	UpdatedBy       string             `json:"updatedBy"`
	UpdatedAt       time.Time          `json:"updatedAt"`
	RenewedAt       *time.Time         `json:"renewedAt"`
}

type MembershipCustomer struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type MembershipPlan struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
