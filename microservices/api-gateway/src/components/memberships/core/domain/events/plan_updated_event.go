package events

const PlanUpdatedEventType = "Memberships.Plans.Updated"

type PlanUpdatedEventPayload struct {
	Id              string
	Name            string
	Featured        bool
	SessionsPerWeek int
	WithCoach       bool
	MonthlyPrice    float64
	GymId           string
	UpdatedBy       string
}
