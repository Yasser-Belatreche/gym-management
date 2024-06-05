package events

const BillPaidEventType = "Memberships.Bills.Paid"

type BillPaidEventPayload struct {
	Id           string
	Amount       float64
	MembershipId string
	CustomerId   string
	PlanId       string
}
