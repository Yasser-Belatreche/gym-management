package create_plan

import "gym-management/src/lib/primitives/application_specific"

type CreatePlanCommand struct {
	Name           string
	Featured       bool
	SessionPerWeek int
	WithCoach      bool
	MonthlyPrice   float64
	GymId          string
	Session        *application_specific.UserSession
}
