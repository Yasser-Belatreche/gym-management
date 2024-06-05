package events

const TrainingSessionEndedEventType = "Memberships.TrainingSessions.Ended"

type TrainingSessionEndedEventPayload struct {
	SessionId    string
	MembershipId string
	CustomerId   string
	PlanId       string
}
