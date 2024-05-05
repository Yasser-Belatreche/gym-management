package events

import "time"

const MembershipRenewedEventType = "Memberships.Renewed"

type MembershipRenewedEventPayload struct {
	Id        string
	StartDate time.Time
	EndDate   *time.Time
	PlanId    string
}
