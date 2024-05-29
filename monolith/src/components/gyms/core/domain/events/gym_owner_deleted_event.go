package events

const GymOwnerDeletedEventType = "Gyms.Owners.Deleted"

type GymOwnerDeletedEventPayload struct {
	Id        string
	DeletedBy string
}
