package events

const GymOwnerCreatedEventType = "Gyms.Owners.Created"

type GymOwnerCreatedEventPayload struct {
	Id          string
	Name        string
	PhoneNumber string
	Email       string
	Password    string
	CreatedBy   string
}
