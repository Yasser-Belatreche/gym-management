package events

import "time"

const BillGeneratedEventType = "Memberships.Bills.Generated"

type BillGeneratedEventPayload struct {
	Id           string
	Amount       float64
	DueDate      time.Time
	MembershipId string
	CustomerId   string
	PlanId       string
}
