package training_sessions

import "time"

type TrainingSessionToReturn struct {
	Id        string
	startedAt time.Time
	endedAt   *time.Time
	Customer  struct {
		Id        string
		FirstName string
		LastName  string
	}
	Gym struct {
		Id   string
		Name string
	}
}
