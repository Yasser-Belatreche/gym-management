package base

import "time"

type TrainingSession struct {
	Id        string                  `json:"id"`
	StartedAt time.Time               `json:"startedAt"`
	EndedAt   *time.Time              `json:"endedAt"`
	Customer  TrainingSessionCustomer `json:"customer"`
	GymId     string                  `json:"gymId"`
}

type TrainingSessionCustomer struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
