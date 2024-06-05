package events

const TrainingSessionStartedEventType = "Memberships.TrainingSessions.Started"

type TrainingSessionStartedEventPayload struct {
	SessionId    string
	MembershipId string
	CustomerId   string
	PlanId       string
}
