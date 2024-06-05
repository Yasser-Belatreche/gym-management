package events

const CustomerRestrictedEventType = "Memberships.Customers.Restricted"

type CustomerRestrictedEventPayload struct {
	Id           string
	RestrictedBy string
}
