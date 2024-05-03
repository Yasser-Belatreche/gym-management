package events

const GymOwnerRestrictedEventType = "Gyms.Owners.Restricted"

type GymOwnerRestrictedEventPayload struct {
	Id           string
	RestrictedBy string
}
