package domain

import (
	"gym-management/src/components/memberships/core/domain/events"
	"gym-management/src/lib/primitives/application_specific"
	"gym-management/src/lib/primitives/generic"
	"time"
)

type Membership struct {
	id                     string
	code                   string
	startDate              time.Time
	endDate                *time.Time
	enabled                bool
	disabledFor            *string
	sessionsPerWeek        int
	withCoach              bool
	monthlyPrice           float64
	unpaidBills            []*Bill
	currentTrainingSession *TrainingSession
	planId                 string
	customerId             string
	renewedAt              *time.Time
	events                 []*events.MembershipEvent[interface{}]
	createdBy              string
	updatedBy              string
}

type MembershipState struct {
	Id                     string
	Code                   string
	StartDate              time.Time
	EndDate                *time.Time
	Enabled                bool
	DisabledFor            *string
	SessionsPerWeek        int
	WithCoach              bool
	MonthlyPrice           float64
	UnpaidBills            []*BillState
	CurrentTrainingSession *TrainingSessionState
	PlanId                 string
	RenewedAt              *time.Time
	CustomerId             string
	CreatedBy              string
	UpdatedBy              string
}

func createMembership(endDate *time.Time, plan *Plan, customer *Customer) (*Membership, *application_specific.ApplicationException) {
	if endDate != nil && endDate.Before(time.Now()) {
		return nil, application_specific.NewValidationException("MEMBERSHIPS.INVALID_END_DATE", "End date cannot be in the past", map[string]string{
			"endDate": endDate.String(),
		})
	}
	if plan.deletedAt != nil {
		return nil, application_specific.NewValidationException("MEMBERSHIPS.INVALID_PLAN", "Plan is deleted", map[string]string{
			"planId": plan.id,
		})
	}
	if customer.deletedAt != nil {
		return nil, application_specific.NewValidationException("MEMBERSHIPS.INVALID_CUSTOMER", "Customer is deleted", map[string]string{
			"customerId": customer.id,
		})
	}
	if customer.restricted {
		return nil, application_specific.NewValidationException("MEMBERSHIPS.INVALID_CUSTOMER", "Customer is restricted", map[string]string{
			"customerId": customer.id,
		})
	}

	membership := &Membership{
		id:                     generic.GenerateUUID(),
		code:                   generic.GenerateRandomString(30),
		startDate:              time.Now(),
		endDate:                endDate,
		enabled:                true,
		disabledFor:            nil,
		sessionsPerWeek:        plan.sessionsPerWeek,
		withCoach:              plan.withCoach,
		monthlyPrice:           plan.monthlyPrice,
		unpaidBills:            make([]*Bill, 0),
		currentTrainingSession: nil,
		events:                 make([]*events.MembershipEvent[interface{}], 0),
		planId:                 plan.id,
		customerId:             customer.id,
		createdBy:              customer.createdBy,
		updatedBy:              customer.createdBy,
	}

	membership.pushEvents(NewMembershipCreatedEvent(membership.State()))

	return membership, nil
}

func MembershipFromState(state *MembershipState) *Membership {
	unpaidBills := make([]*Bill, 0)
	for _, billState := range state.UnpaidBills {
		unpaidBills = append(unpaidBills, BillFromState(billState))
	}

	var trainingSessions *TrainingSession = nil
	if state.CurrentTrainingSession != nil {
		trainingSessions = TrainingSessionFromState(state.CurrentTrainingSession)
	}

	return &Membership{
		id:                     state.Id,
		code:                   state.Code,
		startDate:              state.StartDate,
		endDate:                state.EndDate,
		enabled:                state.Enabled,
		disabledFor:            state.DisabledFor,
		sessionsPerWeek:        state.SessionsPerWeek,
		withCoach:              state.WithCoach,
		monthlyPrice:           state.MonthlyPrice,
		unpaidBills:            unpaidBills,
		currentTrainingSession: trainingSessions,
		planId:                 state.PlanId,
		customerId:             state.CustomerId,
		events:                 make([]*events.MembershipEvent[interface{}], 0),
		createdBy:              state.CreatedBy,
		updatedBy:              state.UpdatedBy,
	}
}

func (m *Membership) GenerateMonthlyBill() *application_specific.ApplicationException {
	if !m.enabled {
		return application_specific.NewValidationException("MEMBERSHIPS.DISABLED", "Membership is disabled", map[string]string{
			"membershipId": m.id,
			"reason":       *m.disabledFor,
		})
	}

	isFirstDayOfMonth := time.Now().Day() == 1
	if !isFirstDayOfMonth {
		return nil
	}

	if m.endDate != nil && time.Now().After(*m.endDate) {
		return m.disable("Membership Ended", "SYSTEM")
	}

	bill, err := BillFromMembership(m)
	if err != nil {
		return err
	}

	m.unpaidBills = append(m.unpaidBills, bill)

	m.pushEvents(NewBillGeneratedEvent(bill.State(), m.State()))

	return nil
}

func (m *Membership) Renew(endDate *time.Time, by string) *application_specific.ApplicationException {
	if m.hasUnpaidBills() {
		return application_specific.NewValidationException("MEMBERSHIPS.UNPAID_BILLS", "Membership has unpaid bills", map[string]string{
			"membershipId": m.id,
		})
	}
	if !m.isDisabledBecauseOf("Membership Ended") {
		return application_specific.NewValidationException("MEMBERSHIPS.NOT_ENDED", "Membership has not ended yet", map[string]string{
			"membershipId": m.id,
		})
	}
	if endDate != nil && endDate.Before(time.Now()) {
		return application_specific.NewValidationException("MEMBERSHIPS.INVALID_END_DATE", "End date cannot be in the past", map[string]string{
			"endDate": endDate.String(),
		})
	}

	m.enabled = true
	m.disabledFor = nil
	now := time.Now()
	m.renewedAt = &now
	m.endDate = endDate
	m.updatedBy = by

	m.pushEvents(NewMembershipRenewedEvent(m.State()))

	return nil
}

func (m *Membership) Cancel(by string) *application_specific.ApplicationException {
	if m.hasUnpaidBills() {
		return application_specific.NewValidationException("MEMBERSHIPS.UNPAID_BILLS", "Membership has unpaid bills", map[string]string{
			"membershipId": m.id,
		})
	}
	if m.IsDisabled() {
		return application_specific.NewValidationException("MEMBERSHIPS.DISABLED", "Membership is disabled", map[string]string{
			"membershipId": m.id,
			"reason":       *m.disabledFor,
		})
	}

	if m.currentTrainingSession != nil {
		_, err := m.EndTrainingSession()
		if err != nil {
			return err
		}
	}

	err := m.disable("Membership Cancelled", by)
	if err != nil {
		return err
	}

	m.pushEvents(NewMembershipCancelledEvent(m.State()))

	return nil
}

func (m *Membership) GymDisabled(by string) *application_specific.ApplicationException {
	if m.currentTrainingSession != nil {
		_, err := m.EndTrainingSession()
		if err != nil {
			return err
		}
	}

	return m.disable("Gym Disabled", by)
}

func (m *Membership) StartTrainingSession(countSessionsThisWeek func() (int, *application_specific.ApplicationException)) (*TrainingSession, *application_specific.ApplicationException) {
	if m.currentTrainingSession != nil {
		return nil, application_specific.NewValidationException("MEMBERSHIPS.TRAINING_SESSIONS.ALREADY_STARTED", "Training session already started", map[string]string{
			"membershipId": m.id,
		})
	}
	if !m.enabled {
		return nil, application_specific.NewValidationException("MEMBERSHIPS.DISABLED", "Membership is disabled", map[string]string{
			"membershipId": m.id,
			"reason":       *m.disabledFor,
		})
	}
	if m.hasMoreThanOneDueBill() {
		err := m.disable("Have more than one due bill", "SYSTEM")
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	count, err := countSessionsThisWeek()
	if err != nil {
		return nil, err
	}

	if count >= m.sessionsPerWeek {
		return nil, application_specific.NewValidationException("MEMBERSHIPS.TRAINING_SESSIONS.LIMIT_REACHED", "Weekly training sessions limit reached", map[string]string{
			"membershipId": m.id,
		})
	}

	trainingSession := CreateTrainingSession()
	m.currentTrainingSession = trainingSession

	m.pushEvents(NewTrainingSessionStartedEvent(trainingSession.State(), m.State()))

	return trainingSession, nil
}

func (m *Membership) EndTrainingSession() (*TrainingSession, *application_specific.ApplicationException) {
	if m.currentTrainingSession == nil {
		return nil, application_specific.NewValidationException("MEMBERSHIPS.TRAINING_SESSIONS.NOT_STARTED", "Training session not started", map[string]string{
			"membershipId": m.id,
		})
	}

	err := m.currentTrainingSession.End()
	if err != nil {
		return nil, err
	}

	m.pushEvents(NewTrainingSessionEndedEvent(m.currentTrainingSession.State(), m.State()))

	return m.currentTrainingSession, nil
}

func (m *Membership) PayBill(id string) *application_specific.ApplicationException {
	bill := m.findBill(id)
	if bill == nil {
		return application_specific.NewNotFoundException("MEMBERSHIPS.BILLS.NOT_FOUND", "Bill not found", map[string]string{
			"billId": id,
		})
	}
	if bill.paid {
		return application_specific.NewValidationException("MEMBERSHIPS.BILLS.ALREADY_PAID", "Bill already paid", map[string]string{
			"billId": id,
		})
	}

	bill.Pay()

	m.pushEvents(NewBillPaidEvent(bill.State(), m.State()))

	if m.isDisabledBecauseOf("Have more than one due bill") && !m.hasMoreThanOneDueBill() {
		m.enabled = true
		m.disabledFor = nil
		m.pushEvents(NewMembershipEnabledEvent(m.State()))
	}

	return nil
}

func (m *Membership) findBill(id string) *Bill {
	for _, bill := range m.unpaidBills {
		if bill.id == id {
			return bill
		}
	}
	return nil
}

func (m *Membership) hasMoreThanOneDueBill() bool {
	count := 0
	for _, bill := range m.unpaidBills {
		if bill.IsDue() {
			count++
		}
	}
	return count > 1
}

func (m *Membership) hasUnpaidBills() bool {
	return len(m.unpaidBills) > 0
}

func (m *Membership) disable(reason string, by string) *application_specific.ApplicationException {
	if !m.enabled {
		return application_specific.NewValidationException("MEMBERSHIPS.ALREADY_DISABLED", "Membership is already disabled", map[string]string{
			"id": m.id,
		})
	}

	m.enabled = false
	m.disabledFor = &reason
	m.updatedBy = by

	bill, err := BillFromMembership(m)
	if err != nil {
		return err
	}
	m.unpaidBills = append(m.unpaidBills, bill)

	m.pushEvents(NewBillGeneratedEvent(bill.State(), m.State()))
	m.pushEvents(NewMembershipDisabledEvent(m.State()))

	return nil
}

func (m *Membership) IsDisabled() bool {
	return !m.enabled
}

func (m *Membership) Equals(another *Membership) bool {
	return m.id == another.id
}

func (m *Membership) State() *MembershipState {
	unpaidBillsState := make([]*BillState, 0)
	for _, bill := range m.unpaidBills {
		unpaidBillsState = append(unpaidBillsState, bill.State())
	}

	var currentTrainingSession *TrainingSessionState = nil
	if m.currentTrainingSession != nil {
		currentTrainingSession = m.currentTrainingSession.State()
	}

	return &MembershipState{
		Id:                     m.id,
		Code:                   m.code,
		StartDate:              m.startDate,
		EndDate:                m.endDate,
		Enabled:                m.enabled,
		DisabledFor:            m.disabledFor,
		SessionsPerWeek:        m.sessionsPerWeek,
		WithCoach:              m.withCoach,
		MonthlyPrice:           m.monthlyPrice,
		CurrentTrainingSession: currentTrainingSession,
		UnpaidBills:            unpaidBillsState,
		RenewedAt:              m.renewedAt,
		PlanId:                 m.planId,
		CustomerId:             m.customerId,
	}
}

func (m *Membership) pushEvents(events ...*events.MembershipEvent[interface{}]) {
	for _, event := range events {
		m.events = append(m.events, event)
	}
}

func (m *Membership) isDisabledBecauseOf(str string) bool {
	return m.disabledFor != nil && *m.disabledFor == str
}

func (m *Membership) PullEvents() []*events.MembershipEvent[interface{}] {
	list := m.events
	m.events = make([]*events.MembershipEvent[interface{}], 0)
	return list
}
