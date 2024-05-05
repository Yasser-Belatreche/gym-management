package events

const CustomerUpdatedEventType = "Memberships.Customers.Updated"

type CustomerUpdatedEventPayload struct {
	Id          string
	FirstName   string
	LastName    string
	Email       string
	Restricted  bool
	PhoneNumber string
	NewPassword *string
	BirthYear   int
	Gender      string
	UpdatedBy   string
}
