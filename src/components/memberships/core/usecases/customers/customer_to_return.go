package customers

import "time"

type CustomerToReturn struct {
	Id          string
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Restricted  bool
	BirthYear   int
	Gender      string
	CreatedBy   string
	UpdatedBy   string
	Membership  struct {
		Id              string
		Enabled         bool
		SessionsPerWeek int
		WithCoach       bool
		MonthlyPrice    float64
		Plan            struct {
			Id   string
			Name string
		}
	}
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedBy *string
	DeletedAt *time.Time
}
