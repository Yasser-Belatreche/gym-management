package bills

import "time"

type BillToReturn struct {
	Id           string
	Amount       float64
	Paid         bool
	Customer     BillToReturnCustomer
	Plan         BillToReturnPlan
	MembershipId string
	GymId        string
	PaidAt       *time.Time
	DueDate      time.Time
	CreatedAt    time.Time
}

type BillToReturnCustomer struct {
	Id        string
	FirstName string
	LastName  string
}

type BillToReturnPlan struct {
	Id   string
	Name string
}
