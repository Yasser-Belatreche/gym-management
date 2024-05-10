package bills

import "time"

type BillToReturn struct {
	Id       string
	Amount   float64
	Paid     bool
	Customer struct {
		Id        string
		FirstName string
		LastName  string
	}
	Plan struct {
		Id   string
		Name string
	}
	PaidAt    *time.Time
	DueDate   time.Time
	CreatedAt time.Time
}
