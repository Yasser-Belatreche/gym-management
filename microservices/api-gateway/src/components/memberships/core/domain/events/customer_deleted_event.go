package events

const CustomerDeletedEventType = "Memberships.Customers.Deleted"

type CustomerDeletedEventPayload struct {
	Id        string
	DeletedBy string
}
