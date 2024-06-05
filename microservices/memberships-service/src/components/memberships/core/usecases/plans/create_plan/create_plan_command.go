package create_plan

import "gym-management-memberships/src/lib/primitives/application_specific"

type CreatePlanCommand struct {
	Name            string
	Featured        bool
	SessionsPerWeek int
	WithCoach       bool
	MonthlyPrice    float64
	GymId           string
	Session         *application_specific.UserSession
}
