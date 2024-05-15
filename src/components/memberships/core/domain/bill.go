package domain

import (
	"gym-management/src/lib/primitives/application_specific"
	"gym-management/src/lib/primitives/generic"
	"time"
)

type Bill struct {
	id        string
	amount    float64
	paid      bool
	paidAt    *time.Time
	dueDate   time.Time
	createdAt time.Time
}

type BillState struct {
	Id        string
	Amount    float64
	Paid      bool
	PaidAt    *time.Time
	DueDate   time.Time
	CreatedAt time.Time
}

func BillFromMembership(membership *Membership) (*Bill, *application_specific.ApplicationException) {
	amount := membership.monthlyPrice

	numberOfDaysSinceStartDate := time.Now().Sub(membership.startDate).Hours() / 24
	if numberOfDaysSinceStartDate < 30 {
		amount = membership.monthlyPrice
	} else {
		dailyPrice := membership.monthlyPrice / 30
		amount = dailyPrice * numberOfDaysSinceStartDate
	}

	bill := &Bill{
		id:        generic.GenerateULID(),
		amount:    amount,
		paid:      false,
		dueDate:   time.Now().AddDate(0, 1, 0),
		createdAt: time.Now(),
	}

	return bill, nil
}

func BillFromState(state *BillState) *Bill {
	return &Bill{
		id:        state.Id,
		amount:    state.Amount,
		paid:      state.Paid,
		paidAt:    state.PaidAt,
		dueDate:   state.DueDate,
		createdAt: state.CreatedAt,
	}
}

func (b *Bill) Pay() {
	b.paid = true
	now := time.Now()
	b.paidAt = &now
}

func (b *Bill) IsDue() bool {
	return !b.paid && time.Now().After(b.dueDate)
}

func (b *Bill) State() *BillState {
	return &BillState{
		Id:        b.id,
		Amount:    b.amount,
		Paid:      b.paid,
		PaidAt:    b.paidAt,
		DueDate:   b.dueDate,
		CreatedAt: b.createdAt,
	}
}
