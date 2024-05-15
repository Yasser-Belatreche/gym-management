package base

import "time"

type Bill struct {
	Id        string       `json:"id"`
	Amount    float64      `json:"amount"`
	Paid      bool         `json:"paid"`
	Customer  BillCustomer `json:"customer"`
	Plan      BillPlan     `json:"plan"`
	GymId     string       `json:"gymId"`
	PaidAt    *time.Time   `json:"paidAt"`
	DueDate   time.Time    `json:"dueDate"`
	CreatedAt time.Time    `json:"createdAt"`
}

type BillCustomer struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type BillPlan struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
