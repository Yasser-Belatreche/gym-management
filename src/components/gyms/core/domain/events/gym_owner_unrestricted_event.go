package events

const GymOwnerUnrestrictedEventType = "Gyms.Owners.Unrestricted"

type GymOwnerUnrestrictedEventPayload struct {
	Id             string
	UnrestrictedBy string
}
