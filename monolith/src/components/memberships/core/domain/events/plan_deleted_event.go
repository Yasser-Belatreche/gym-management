package events

const PlanDeletedEventType = "Memberships.Plans.Deleted"

type PlanDeletedEventPayload struct {
	Id        string
	DeletedBy string
}
