package events

const GymEnabledEventType = "Gyms.Enabled"

type GymEnabledEventPayload struct {
	Id        string
	EnabledBy string
	OwnerId   string
}
