package events

const PlanCreatedEventType = "Memberships.Plans.Created"

type PlanCreatedEventPayload struct {
	Id              string
	Name            string
	Featured        bool
	SessionsPerWeek int
	WithCoach       bool
	MonthlyPrice    float64
	GymId           string
	CreatedBy       string
}
