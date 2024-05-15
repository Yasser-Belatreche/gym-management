package plans

import "time"

type PlanToReturn struct {
	Id             string
	Name           string
	Featured       bool
	SessionPerWeek int
	WithCoach      bool
	MonthlyPrice   float64
	GymId          string
	CreatedBy      string
	UpdatedBy      string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedBy      *string
	DeletedAt      *time.Time
}
