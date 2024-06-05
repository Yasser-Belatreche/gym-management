package update_plan

import "gym-management/src/lib/primitives/application_specific"

type UpdatePlanCommand struct {
	Id              string
	Name            string
	Featured        bool
	SessionsPerWeek int
	WithCoach       bool
	MonthlyPrice    float64
	Session         *application_specific.UserSession
}
