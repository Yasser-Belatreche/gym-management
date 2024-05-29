package events

const GymDeletedEventType = "Gyms.Deleted"

type GymDeletedEventPayload struct {
	Id        string
	DeletedBy string
	OwnerId   string
}
