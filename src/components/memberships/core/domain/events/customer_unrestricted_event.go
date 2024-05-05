package events

const CustomerUnrestrictedEventType = "Memberships.Customers.Unrestricted"

type CustomerUnrestrictedEventPayload struct {
	Id             string
	UnrestrictedBy string
}
