package events

const CustomerCreatedEventType = "Memberships.Customers.Created"

type CustomerCreatedEventPayload struct {
	Id          string
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Password    string
	BirthYear   int
	Gender      string
	CreatedBy   string
	PlanId      string
}
