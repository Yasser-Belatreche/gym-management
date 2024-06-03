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
	} `gorm:"embedded; embeddedPrefix:customer_"`
	Plan struct {
		Id   string
		Name string
	} `gorm:"embedded; embeddedPrefix:plan_"`
	GymId     string
	CreatedBy string
	CreatedAt time.Time
	UpdatedBy string
	UpdatedAt time.Time
	RenewedAt *time.Time
}
