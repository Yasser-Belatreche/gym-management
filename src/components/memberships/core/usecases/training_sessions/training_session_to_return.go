package training_sessions

import "time"

type TrainingSessionToReturn struct {
	Id        string
	StartedAt time.Time
	EndedAt   *time.Time
	Customer  struct {
		Id        string
		FirstName string
		LastName  string
	}
	GymId string
}
