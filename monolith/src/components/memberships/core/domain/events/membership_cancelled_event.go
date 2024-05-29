package events

const MembershipCancelledEventType = "Memberships.Cancelled"

type MembershipCancelledEventPayload struct {
	Id     string
	Reason string
}
