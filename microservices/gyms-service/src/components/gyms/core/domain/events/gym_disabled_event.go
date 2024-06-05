package events

const GymDisabledEventType = "Gyms.Disabled"

type GymDisabledEventPayload struct {
	Id          string
	DisabledBy  string
	DisabledFor string
	OwnerId     string
}
