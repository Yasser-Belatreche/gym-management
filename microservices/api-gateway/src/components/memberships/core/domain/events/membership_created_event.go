package events

import "time"

const MembershipCreatedEventType = "Memberships.Created"

type MembershipCreatedEventPayload struct {
	Id              string
	Code            string
	StartDate       time.Time
	EndDate         *time.Time
	SessionsPerWeek int
	WithCoach       bool
	MonthlyPrice    float64
	PlanId          string
	CustomerId      string
}
