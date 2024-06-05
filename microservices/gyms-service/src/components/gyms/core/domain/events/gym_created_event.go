package events

const GymCreatedEventType = "Gyms.Created"

type GymCreatedEventPayload struct {
	Id        string
	Name      string
	Address   string
	OwnerId   string
	CreatedBy string
}
