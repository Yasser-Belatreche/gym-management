package events

const GymOwnerUpdatedEventType = "Gyms.Owners.Updated"

type GymOwnerUpdatedEventPayload struct {
	Id          string
	Name        string
	PhoneNumber string
	Email       string
	Restricted  bool
	NewPassword *string
	UpdatedBy   string
	Gyms        []string
	EnabledGyms []string
}
