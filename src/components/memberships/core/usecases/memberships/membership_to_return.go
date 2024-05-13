package memberships

import "time"

type MembershipToReturn struct {
	Id              string
	StartDate       time.Time
	EndDate         *time.Time
	Enabled         bool
	DisabledFor     *string
	SessionsPerWeek int
	WithCoach       bool
	MonthlyPrice    float64
	Customer        struct {
		Id        string
		FirstName string
		LastName  string
	}
	Plan struct {
		Id   string
		Name string
	}
	Gym struct {
		Id   string
		Name string
	}
	CreatedBy string
	CreatedAt time.Time
	UpdatedBy string
	UpdatedAt time.Time
	RenewedAt *time.Time
}
