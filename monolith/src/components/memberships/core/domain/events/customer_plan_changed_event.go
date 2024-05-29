package events

const CustomerPlanChangedEventType = "Memberships.Customers.PlanChanged"

type CustomerPlanChangedEventPayload struct {
	Id     string
	PlanId string
}
