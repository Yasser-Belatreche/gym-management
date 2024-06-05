package events

const MembershipDisabledEventType = "Memberships.Disabled"

type MembershipDisabledEventPayload struct {
	Id     string
	Reason string
}
