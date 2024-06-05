package domain

import (
	"gym-management-memberships/src/components/memberships/core/domain/events"
)

func NewCustomerCreatedEvent(state *CustomerState, password string, planId string) *events.MembershipEvent[interface{}] {
	return events.NewMembershipEvent(events.CustomerCreatedEventType, events.CustomerCreatedEventPayload{
		Id:          state.Id,
		FirstName:   state.FirstName,
		LastName:    state.LastName,
		PhoneNumber: state.PhoneNumber,
		Email:       state.Email,
		Password:    password,
		CreatedBy:   state.CreatedBy,
		PlanId:      planId,
	})
}

func NewCustomerUpdatedEvent(state *CustomerState, newPassword *string) *events.MembershipEvent[interface{}] {
	return events.NewMembershipEvent(events.CustomerUpdatedEventType, events.CustomerUpdatedEventPayload{
		Id:          state.Id,
		FirstName:   state.FirstName,
		LastName:    state.LastName,
		PhoneNumber: state.PhoneNumber,
		Email:       state.Email,
		Restricted:  state.Restricted,
		NewPassword: newPassword,
		UpdatedBy:   state.UpdatedBy,
	})
}

func NewCustomerPlanChangedEvent(state *CustomerState, planId string) *events.MembershipEvent[interface{}] {
	return events.NewMembershipEvent(events.CustomerPlanChangedEventType, events.CustomerPlanChangedEventPayload{
		Id:     state.Id,
		PlanId: planId,
	})
}

func NewCustomerUnrestrictedEvent(state *CustomerState) *events.MembershipEvent[interface{}] {
	return events.NewMembershipEvent(events.CustomerUnrestrictedEventType, events.CustomerUnrestrictedEventPayload{
		Id:             state.Id,
		UnrestrictedBy: state.UpdatedBy,
	})
}

func NewCustomerRestrictedEvent(state *CustomerState) *events.MembershipEvent[interface{}] {
	return events.NewMembershipEvent(events.CustomerRestrictedEventType, events.CustomerRestrictedEventPayload{
		Id:           state.Id,
		RestrictedBy: state.UpdatedBy,
	})
}

func NewCustomerDeletedEvent(state *CustomerState) *events.MembershipEvent[interface{}] {
	return events.NewMembershipEvent(events.CustomerDeletedEventType, events.CustomerDeletedEventPayload{
		Id:        state.Id,
		DeletedBy: *state.DeletedBy,
	})
}

func NewMembershipCreatedEvent(state *MembershipState) *events.MembershipEvent[interface{}] {
	return events.NewMembershipEvent(events.MembershipCreatedEventType, events.MembershipCreatedEventPayload{
		Id:              state.Id,
		Code:            state.Code,
		StartDate:       state.StartDate,
		EndDate:         state.EndDate,
		SessionsPerWeek: state.SessionsPerWeek,
		WithCoach:       state.WithCoach,
		MonthlyPrice:    state.MonthlyPrice,
		CustomerId:      state.CustomerId,
		PlanId:          state.PlanId,
	})
}

func NewMembershipDisabledEvent(state *MembershipState) *events.MembershipEvent[interface{}] {
	return events.NewMembershipEvent(events.MembershipDisabledEventType, events.MembershipDisabledEventPayload{
		Id:     state.Id,
		Reason: *state.DisabledFor,
	})
}

func NewMembershipEnabledEvent(state *MembershipState) *events.MembershipEvent[interface{}] {
	return events.NewMembershipEvent(events.MembershipEnabledEventType, events.MembershipEnabledEventPayload{
		Id: state.Id,
	})
}

func NewMembershipCancelledEvent(state *MembershipState) *events.MembershipEvent[interface{}] {
	return events.NewMembershipEvent(events.MembershipCancelledEventType, events.MembershipCancelledEventPayload{
		Id:     state.Id,
		Reason: *state.DisabledFor,
	})
}

func NewMembershipRenewedEvent(state *MembershipState) *events.MembershipEvent[interface{}] {
	return events.NewMembershipEvent(events.MembershipRenewedEventType, events.MembershipRenewedEventPayload{
		Id:        state.Id,
		StartDate: state.StartDate,
		EndDate:   state.EndDate,
		PlanId:    state.PlanId,
	})
}

func NewBillGeneratedEvent(state *BillState, membership *MembershipState) *events.MembershipEvent[interface{}] {
	return events.NewMembershipEvent(events.BillGeneratedEventType, events.BillGeneratedEventPayload{
		Id:           state.Id,
		Amount:       state.Amount,
		DueDate:      state.DueDate,
		MembershipId: membership.Id,
		CustomerId:   membership.CustomerId,
		PlanId:       membership.PlanId,
	})
}

func NewBillPaidEvent(state *BillState, membership *MembershipState) *events.MembershipEvent[interface{}] {
	return events.NewMembershipEvent(events.BillPaidEventType, events.BillPaidEventPayload{
		Id:           state.Id,
		Amount:       state.Amount,
		MembershipId: membership.Id,
		CustomerId:   membership.CustomerId,
		PlanId:       membership.PlanId,
	})
}

func NewTrainingSessionStartedEvent(state *TrainingSessionState, membership *MembershipState) *events.MembershipEvent[interface{}] {
	return events.NewMembershipEvent(events.TrainingSessionStartedEventType, events.TrainingSessionStartedEventPayload{
		SessionId:    state.Id,
		MembershipId: membership.Id,
		CustomerId:   membership.CustomerId,
		PlanId:       membership.PlanId,
	})
}

func NewTrainingSessionEndedEvent(state *TrainingSessionState, membership *MembershipState) *events.MembershipEvent[interface{}] {
	return events.NewMembershipEvent(events.TrainingSessionEndedEventType, events.TrainingSessionEndedEventPayload{
		SessionId:    state.Id,
		MembershipId: membership.Id,
		CustomerId:   membership.CustomerId,
		PlanId:       membership.PlanId,
	})
}

func NewPlanCreatedEvent(state *PlanState) *events.MembershipEvent[interface{}] {
	return events.NewMembershipEvent(events.PlanCreatedEventType, events.PlanCreatedEventPayload{
		Id:              state.Id,
		Name:            state.Name,
		Featured:        state.Featured,
		SessionsPerWeek: state.SessionsPerWeek,
		WithCoach:       state.WithCoach,
		MonthlyPrice:    state.MonthlyPrice,
		GymId:           state.GymId,
		CreatedBy:       state.CreatedBy,
	})
}

func NewPlanUpdatedEvent(state *PlanState) *events.MembershipEvent[interface{}] {
	return events.NewMembershipEvent(events.PlanUpdatedEventType, events.PlanUpdatedEventPayload{
		Id:              state.Id,
		Name:            state.Name,
		Featured:        state.Featured,
		SessionsPerWeek: state.SessionsPerWeek,
		WithCoach:       state.WithCoach,
		MonthlyPrice:    state.MonthlyPrice,
		GymId:           state.GymId,
		UpdatedBy:       state.UpdatedBy,
	})
}

func NewPlanDeletedEvent(state *PlanState) *events.MembershipEvent[interface{}] {
	return events.NewMembershipEvent(events.PlanDeletedEventType, events.PlanDeletedEventPayload{
		Id:        state.Id,
		DeletedBy: *state.DeletedBy,
	})
}
