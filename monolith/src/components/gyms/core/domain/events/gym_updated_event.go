package events

const GymUpdatedEventType = "Gyms.Updated"

type GymUpdatedEventPayload struct {
	Id        string
	Name      string
	Address   string
	OwnerId   string
	UpdatedBy string
}
