package domain

type GymOwner struct {
	id          string
	name        string
	phoneNumber string
	email       string
	restricted  bool
	createdBy   string
	updatedBy   string
}
