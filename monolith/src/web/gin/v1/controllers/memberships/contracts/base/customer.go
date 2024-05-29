package base

import "time"

type Customer struct {
	Id          string             `json:"id"`
	FirstName   string             `json:"firstName"`
	LastName    string             `json:"lastName"`
	Email       string             `json:"email"`
	PhoneNumber string             `json:"phoneNumber"`
	Restricted  bool               `json:"restricted"`
	BirthYear   int                `json:"birthYear"`
	Gender      string             `json:"gender"`
	CreatedBy   string             `json:"createdBy"`
	UpdatedBy   string             `json:"updatedBy"`
	Membership  CustomerMembership `json:"membership"`
	GymId       string             `json:"gymId"`
	CreatedAt   time.Time          `json:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt"`
	DeletedBy   *string            `json:"deletedBy"`
	DeletedAt   *time.Time         `json:"deletedAt"`
}

type CustomerMembership struct {
	Id              string                 `json:"id"`
	Enabled         bool                   `json:"enabled"`
	SessionsPerWeek int                    `json:"sessionsPerWeek"`
	WithCoach       bool                   `json:"withCoach"`
	MonthlyPrice    float64                `json:"monthlyPrice"`
	Plan            CustomerMembershipPlan `json:"plan"`
}

type CustomerMembershipPlan struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
