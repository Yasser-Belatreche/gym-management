package memberships

import (
	auth "gym-management/src/components/auth/core/domain"
	"gym-management/src/components/memberships/core/usecases/bills/get_bill"
	"gym-management/src/components/memberships/core/usecases/bills/get_bills"
	"gym-management/src/components/memberships/core/usecases/bills/mark_bill_as_paid"
	"gym-management/src/components/memberships/core/usecases/customers/change_customer_plan"
	"gym-management/src/components/memberships/core/usecases/customers/create_customer"
	"gym-management/src/components/memberships/core/usecases/customers/delete_customer"
	"gym-management/src/components/memberships/core/usecases/customers/get_customer"
	"gym-management/src/components/memberships/core/usecases/customers/get_customers"
	"gym-management/src/components/memberships/core/usecases/customers/restrict_customer"
	"gym-management/src/components/memberships/core/usecases/customers/unrestrict_customer"
	"gym-management/src/components/memberships/core/usecases/customers/update_customer"
	"gym-management/src/components/memberships/core/usecases/memberships/cancel_membership"
	"gym-management/src/components/memberships/core/usecases/memberships/get_membership"
	"gym-management/src/components/memberships/core/usecases/memberships/get_membership_badge"
	"gym-management/src/components/memberships/core/usecases/memberships/get_memberships"
	"gym-management/src/components/memberships/core/usecases/memberships/renew_membership"
	"gym-management/src/components/memberships/core/usecases/plans/create_plan"
	"gym-management/src/components/memberships/core/usecases/plans/delete_plan"
	"gym-management/src/components/memberships/core/usecases/plans/get_plan"
	"gym-management/src/components/memberships/core/usecases/plans/get_plans"
	"gym-management/src/components/memberships/core/usecases/plans/update_plan"
	"gym-management/src/components/memberships/core/usecases/training_sessions/end_training_session"
	"gym-management/src/components/memberships/core/usecases/training_sessions/get_training_session"
	"gym-management/src/components/memberships/core/usecases/training_sessions/get_training_sessions"
	"gym-management/src/components/memberships/core/usecases/training_sessions/start_training_session"
	"gym-management/src/lib/primitives/application_specific"
)

type AuthorizationDecorator struct {
	manager Manager
}

func (a *AuthorizationDecorator) CreateCustomer(command *create_customer.CreateCustomerCommand) (*create_customer.CreateCustomerCommandResponse, *application_specific.ApplicationException) {
	err := application_specific.AssertUserRole(command.Session, auth.RoleGymOwner)
	if err != nil {
		return nil, err
	}

	plan, err := a.manager.GetPlan(&get_plan.GetPlanQuery{Id: command.PlanId, Session: command.Session})
	if err != nil {
		return nil, err
	}

	if !command.Session.IsOwnerOfEnabledGym(plan.Gym.Id) {
		return nil, application_specific.NewForbiddenException("OPERATION_NOT_ALLOWED", "Target gym disabled or not your gym", map[string]string{
			"planId": command.PlanId,
			"gymId":  plan.Gym.Id,
		})
	}

	return a.manager.CreateCustomer(command)
}

func (a *AuthorizationDecorator) UpdateCustomer(command *update_customer.UpdateCustomerCommand) (*update_customer.UpdateCustomerCommandResponse, *application_specific.ApplicationException) {
	err := application_specific.AssertUserRole(command.Session, auth.RoleGymOwner)
	if err != nil {
		return nil, err
	}

	customer, err := a.manager.GetCustomer(&get_customer.GetCustomerQuery{
		Id:      command.Id,
		Session: command.Session,
	})
	if err != nil {
		return nil, err
	}

	if !command.Session.IsOwnerOfEnabledGym(customer.Gym.Id) {
		return nil, application_specific.NewForbiddenException("OPERATION_NOT_ALLOWED", "Target gym disabled or not your gym", map[string]string{
			"customerId": command.Id,
			"gymId":      customer.Gym.Id,
		})
	}

	return a.manager.UpdateCustomer(command)
}

func (a *AuthorizationDecorator) RestrictCustomer(command *restrict_customer.RestrictCustomerCommand) (*restrict_customer.RestrictCustomerCommandResponse, *application_specific.ApplicationException) {
	err := application_specific.AssertUserRole(command.Session, auth.RoleGymOwner)
	if err != nil {
		return nil, err
	}

	customer, err := a.manager.GetCustomer(&get_customer.GetCustomerQuery{
		Id:      command.Id,
		Session: command.Session,
	})
	if err != nil {
		return nil, err
	}

	if !command.Session.IsOwnerOfEnabledGym(customer.Gym.Id) {
		return nil, application_specific.NewForbiddenException("OPERATION_NOT_ALLOWED", "Target gym disabled or not your gym", map[string]string{
			"customerId": command.Id,
			"gymId":      customer.Gym.Id,
		})
	}

	return a.manager.RestrictCustomer(command)
}

func (a *AuthorizationDecorator) UnrestrictCustomer(command *unrestrict_customer.UnrestrictCustomerCommand) (*unrestrict_customer.UnrestrictCustomerCommandResponse, *application_specific.ApplicationException) {
	err := application_specific.AssertUserRole(command.Session, auth.RoleGymOwner)
	if err != nil {
		return nil, err
	}

	customer, err := a.manager.GetCustomer(&get_customer.GetCustomerQuery{
		Id:      command.Id,
		Session: command.Session,
	})
	if err != nil {
		return nil, err
	}

	if !command.Session.IsOwnerOfEnabledGym(customer.Gym.Id) {
		return nil, application_specific.NewForbiddenException("OPERATION_NOT_ALLOWED", "Target gym disabled or not your gym", map[string]string{
			"customerId": command.Id,
			"gymId":      customer.Gym.Id,
		})
	}

	return a.manager.UnrestrictCustomer(command)
}

func (a *AuthorizationDecorator) DeleteCustomer(command *delete_customer.DeleteCustomerCommand) (*delete_customer.DeleteCustomerCommandResponse, *application_specific.ApplicationException) {
	err := application_specific.AssertUserRole(command.Session, auth.RoleGymOwner)
	if err != nil {
		return nil, err
	}

	customer, err := a.manager.GetCustomer(&get_customer.GetCustomerQuery{
		Id:      command.Id,
		Session: command.Session,
	})
	if err != nil {
		return nil, err
	}

	if !command.Session.IsOwnerOfEnabledGym(customer.Gym.Id) {
		return nil, application_specific.NewForbiddenException("OPERATION_NOT_ALLOWED", "Target gym disabled or not your gym", map[string]string{
			"customerId": command.Id,
			"gymId":      customer.Gym.Id,
		})
	}

	return a.manager.DeleteCustomer(command)
}

func (a *AuthorizationDecorator) ChangeCustomerPlan(command *change_customer_plan.ChangeCustomerPlanCommand) (*change_customer_plan.ChangeCustomerPlanCommandResponse, *application_specific.ApplicationException) {
	err := application_specific.AssertUserRole(command.Session, auth.RoleGymOwner)
	if err != nil {
		return nil, err
	}

	customer, err := a.manager.GetCustomer(&get_customer.GetCustomerQuery{
		Id:      command.CustomerId,
		Session: command.Session,
	})
	if err != nil {
		return nil, err
	}

	if !command.Session.IsOwnerOfEnabledGym(customer.Gym.Id) {
		return nil, application_specific.NewForbiddenException("OPERATION_NOT_ALLOWED", "Target gym disabled or not your gym", map[string]string{
			"customerId": command.CustomerId,
			"gymId":      customer.Gym.Id,
		})
	}

	plan, err := a.manager.GetPlan(&get_plan.GetPlanQuery{Id: command.PlanId, Session: command.Session})
	if err != nil {
		return nil, err
	}

	if !command.Session.IsOwnerOfEnabledGym(plan.Gym.Id) {
		return nil, application_specific.NewForbiddenException("OPERATION_NOT_ALLOWED", "Target gym disabled or not your gym", map[string]string{
			"planId": command.PlanId,
			"gymId":  plan.Gym.Id,
		})
	}

	return a.manager.ChangeCustomerPlan(command)
}

func (a *AuthorizationDecorator) GetCustomer(query *get_customer.GetCustomerQuery) (*get_customer.GetCustomerQueryResponse, *application_specific.ApplicationException) {
	err := application_specific.AssertUserRole(query.Session, auth.RoleGymOwner, auth.RoleCustomer, auth.RoleAdmin)
	if err != nil {
		return nil, err
	}

	customer, err := a.manager.GetCustomer(query)
	if err != nil {
		return nil, err
	}

	if query.Session.RoleIsOneOf(auth.RoleGymOwner) && !query.Session.IsOwnerOfEnabledGym(customer.Gym.Id) {
		return nil, application_specific.NewForbiddenException("OPERATION_NOT_ALLOWED", "Target gym disabled or not your gym", map[string]string{
			"customerId": query.Id,
			"gymId":      customer.Gym.Id,
		})
	}
	if query.Session.RoleIsOneOf(auth.RoleCustomer) && query.Session.UserId() != query.Id {
		return nil, application_specific.NewForbiddenException("OPERATION_NOT_ALLOWED", "You are not the target customer", map[string]string{
			"customerId": query.Id,
		})
	}

	return customer, nil
}

func (a *AuthorizationDecorator) GetCustomers(query *get_customers.GetCustomersQuery) (*get_customers.GetCustomersQueryResponse, *application_specific.ApplicationException) {
	err := application_specific.AssertUserRole(query.Session, auth.RoleGymOwner, auth.RoleCustomer, auth.RoleAdmin)
	if err != nil {
		return nil, err
	}

	if query.Session.RoleIsOneOf(auth.RoleGymOwner) {
		query.GymId = query.Session.User.Profile.EnabledOwnedGyms
	}
	if query.Session.RoleIsOneOf(auth.RoleCustomer) {
		query.Id = []string{query.Session.UserId()}
	}

	return a.manager.GetCustomers(query)
}

func (a *AuthorizationDecorator) CancelMembership(command *cancel_membership.CancelMembershipCommand) (*cancel_membership.CancelMembershipCommandResponse, *application_specific.ApplicationException) {
	err := application_specific.AssertUserRole(command.Session, auth.RoleGymOwner)
	if err != nil {
		return nil, err
	}

	membership, err := a.manager.GetMembership(&get_membership.GetMembershipQuery{MembershipId: command.MembershipId, Session: command.Session})
	if err != nil {
		return nil, err
	}

	if !command.Session.IsOwnerOfEnabledGym(membership.Gym.Id) {
		return nil, application_specific.NewForbiddenException("OPERATION_NOT_ALLOWED", "Target gym disabled or not your gym", map[string]string{
			"membershipId": command.MembershipId,
			"gymId":        membership.Gym.Id,
		})
	}

	return a.manager.CancelMembership(command)
}

func (a *AuthorizationDecorator) RenewMembership(command *renew_membership.RenewMembershipCommand) (*renew_membership.RenewMembershipCommandResponse, *application_specific.ApplicationException) {
	err := application_specific.AssertUserRole(command.Session, auth.RoleGymOwner)
	if err != nil {
		return nil, err
	}

	membership, err := a.manager.GetMembership(&get_membership.GetMembershipQuery{MembershipId: command.MembershipId, Session: command.Session})
	if err != nil {
		return nil, err
	}

	if !command.Session.IsOwnerOfEnabledGym(membership.Gym.Id) {
		return nil, application_specific.NewForbiddenException("OPERATION_NOT_ALLOWED", "Target gym disabled or not your gym", map[string]string{
			"membershipId": command.MembershipId,
			"gymId":        membership.Gym.Id,
		})
	}

	return a.manager.RenewMembership(command)
}

func (a *AuthorizationDecorator) GetMembership(query *get_membership.GetMembershipQuery) (*get_membership.GetMembershipQueryResponse, *application_specific.ApplicationException) {
	err := application_specific.AssertUserRole(query.Session, auth.RoleGymOwner, auth.RoleCustomer)
	if err != nil {
		return nil, err
	}

	membership, err := a.manager.GetMembership(&get_membership.GetMembershipQuery{MembershipId: query.MembershipId, Session: query.Session})
	if err != nil {
		return nil, err
	}

	if query.Session.RoleIsOneOf(auth.RoleGymOwner) && !query.Session.IsOwnerOfEnabledGym(membership.Gym.Id) {
		return nil, application_specific.NewForbiddenException("OPERATION_NOT_ALLOWED", "Target gym disabled or not your gym", map[string]string{
			"membershipId": query.MembershipId,
			"gymId":        membership.Gym.Id,
		})
	}

	if query.Session.RoleIsOneOf(auth.RoleCustomer) && query.Session.UserId() != membership.Customer.Id {
		return nil, application_specific.NewForbiddenException("OPERATION_NOT_ALLOWED", "You are not the target customer", map[string]string{
			"membershipId": query.MembershipId,
		})
	}

	return membership, nil
}

func (a *AuthorizationDecorator) GetMembershipBadge(query *get_membership_badge.GetMembershipBadgeQuery) (*get_membership_badge.GetMembershipBadgeQueryResponse, *application_specific.ApplicationException) {
	err := application_specific.AssertUserRole(query.Session, auth.RoleGymOwner)
	if err != nil {
		return nil, err
	}

	membership, err := a.manager.GetMembership(&get_membership.GetMembershipQuery{MembershipId: query.MembershipId, Session: query.Session})
	if err != nil {
		return nil, err
	}

	if !query.Session.IsOwnerOfEnabledGym(membership.Gym.Id) {
		return nil, application_specific.NewForbiddenException("OPERATION_NOT_ALLOWED", "Target gym disabled or not your gym", map[string]string{
			"membershipId": membership.Id,
			"gymId":        membership.Gym.Id,
		})
	}

	return a.manager.GetMembershipBadge(query)
}

func (a *AuthorizationDecorator) GetMemberships(query *get_memberships.GetMembershipsQuery) (*get_memberships.GetMembershipsQueryResponse, *application_specific.ApplicationException) {
	err := application_specific.AssertUserRole(query.Session, auth.RoleGymOwner, auth.RoleAdmin, auth.RoleCustomer)
	if err != nil {
		return nil, err
	}

	if query.Session.RoleIsOneOf(auth.RoleGymOwner) {
		query.GymId = query.Session.User.Profile.EnabledOwnedGyms
	}

	if query.Session.RoleIsOneOf(auth.RoleCustomer) {
		query.CustomerId = []string{query.Session.UserId()}
	}

	return a.manager.GetMemberships(query)
}

func (a *AuthorizationDecorator) StartTrainingSession(command *start_training_session.StartTrainingSessionCommand) (*start_training_session.StartTrainingSessionCommandResponse, *application_specific.ApplicationException) {
	err := application_specific.AssertUserRole(command.Session, auth.RoleGymOwner)
	if err != nil {
		return nil, err
	}

	membership, err := a.manager.GetMembership(&get_membership.GetMembershipQuery{MembershipCode: command.MembershipCode, Session: command.Session})
	if err != nil {
		return nil, err
	}

	if !command.Session.IsOwnerOfEnabledGym(membership.Gym.Id) {
		return nil, application_specific.NewForbiddenException("OPERATION_NOT_ALLOWED", "Target gym disabled or not your gym", map[string]string{
			"membershipCode": command.MembershipCode,
			"gymId":          membership.Gym.Id,
		})
	}

	return a.manager.StartTrainingSession(command)
}

func (a *AuthorizationDecorator) EndTrainingSession(command *end_training_session.EndTrainingSessionCommand) (*end_training_session.EndTrainingSessionCommandResponse, *application_specific.ApplicationException) {
	err := application_specific.AssertUserRole(command.Session, auth.RoleGymOwner)
	if err != nil {
		return nil, err
	}

	membership, err := a.manager.GetMembership(&get_membership.GetMembershipQuery{MembershipCode: command.MembershipCode, Session: command.Session})
	if err != nil {
		return nil, err
	}

	if !command.Session.IsOwnerOfEnabledGym(membership.Gym.Id) {
		return nil, application_specific.NewForbiddenException("OPERATION_NOT_ALLOWED", "Target gym disabled or not your gym", map[string]string{
			"membershipCode": command.MembershipCode,
			"gymId":          membership.Gym.Id,
		})
	}

	return a.manager.EndTrainingSession(command)
}

func (a *AuthorizationDecorator) GetTrainingSession(query *get_training_session.GetTrainingSessionQuery) (*get_training_session.GetTrainingSessionQueryResponse, *application_specific.ApplicationException) {
	err := application_specific.AssertUserRole(query.Session, auth.RoleGymOwner, auth.RoleCustomer, auth.RoleAdmin)
	if err != nil {
		return nil, err
	}

	trainingSession, err := a.manager.GetTrainingSession(query)
	if err != nil {
		return nil, err
	}

	if query.Session.RoleIsOneOf(auth.RoleGymOwner) && !query.Session.IsOwnerOfEnabledGym(trainingSession.Gym.Id) {
		return nil, application_specific.NewForbiddenException("OPERATION_NOT_ALLOWED", "Target gym disabled or not your gym", map[string]string{
			"trainingSessionId": trainingSession.Id,
			"gymId":             trainingSession.Gym.Id,
		})
	}

	if query.Session.RoleIsOneOf(auth.RoleCustomer) && query.Session.UserId() != trainingSession.Customer.Id {
		return nil, application_specific.NewForbiddenException("OPERATION_NOT_ALLOWED", "You are not the target customer", map[string]string{
			"trainingSessionId": trainingSession.Id,
		})
	}

	return trainingSession, nil
}

func (a *AuthorizationDecorator) GetTrainingSessions(query *get_training_sessions.GetTrainingSessionsQuery) (*get_training_sessions.GetTrainingSessionsQueryResponse, *application_specific.ApplicationException) {
	err := application_specific.AssertUserRole(query.Session, auth.RoleGymOwner, auth.RoleCustomer, auth.RoleAdmin)
	if err != nil {
		return nil, err
	}

	if query.Session.RoleIsOneOf(auth.RoleGymOwner) {
		query.GymId = query.Session.User.Profile.EnabledOwnedGyms
	}

	if query.Session.RoleIsOneOf(auth.RoleCustomer) {
		query.CustomerId = []string{query.Session.UserId()}
	}

	return a.manager.GetTrainingSessions(query)
}

func (a *AuthorizationDecorator) MarkBillAsPaid(command *mark_bill_as_paid.MarkBillAsPaidCommand) (*mark_bill_as_paid.MarkBillAsPaidCommandResponse, *application_specific.ApplicationException) {
	err := application_specific.AssertUserRole(command.Session, auth.RoleGymOwner)
	if err != nil {
		return nil, err
	}

	bill, err := a.manager.GetBill(&get_bill.GetBillQuery{BillId: command.BillId, Session: command.Session})
	if err != nil {
		return nil, err
	}

	if !command.Session.IsOwnerOfEnabledGym(bill.Gym.Id) {
		return nil, application_specific.NewForbiddenException("OPERATION_NOT_ALLOWED", "Target gym disabled or not your gym", map[string]string{
			"billId": command.BillId,
			"gymId":  bill.Gym.Id,
		})
	}

	return a.manager.MarkBillAsPaid(command)
}

func (a *AuthorizationDecorator) GetBill(query *get_bill.GetBillQuery) (*get_bill.GetBillQueryResponse, *application_specific.ApplicationException) {
	err := application_specific.AssertUserRole(query.Session, auth.RoleGymOwner, auth.RoleCustomer, auth.RoleAdmin)
	if err != nil {
		return nil, err
	}

	bill, err := a.manager.GetBill(query)
	if err != nil {
		return nil, err
	}

	if query.Session.RoleIsOneOf(auth.RoleGymOwner) && !query.Session.IsOwnerOfEnabledGym(bill.Gym.Id) {
		return nil, application_specific.NewForbiddenException("OPERATION_NOT_ALLOWED", "Target gym disabled or not your gym", map[string]string{
			"billId": query.BillId,
			"gymId":  bill.Gym.Id,
		})
	}

	if query.Session.RoleIsOneOf(auth.RoleCustomer) && query.Session.UserId() != bill.Customer.Id {
		return nil, application_specific.NewForbiddenException("OPERATION_NOT_ALLOWED", "You are not the target customer", map[string]string{
			"billId": query.BillId,
		})
	}

	return bill, nil
}

func (a *AuthorizationDecorator) GetBills(query *get_bills.GetBillsQuery) (*get_bills.GetBillsQueryResponse, *application_specific.ApplicationException) {
	err := application_specific.AssertUserRole(query.Session, auth.RoleGymOwner, auth.RoleCustomer, auth.RoleAdmin)
	if err != nil {
		return nil, err
	}

	if query.Session.RoleIsOneOf(auth.RoleGymOwner) {
		query.GymId = query.Session.User.Profile.EnabledOwnedGyms
	}

	if query.Session.RoleIsOneOf(auth.RoleCustomer) {
		query.CustomerId = []string{query.Session.UserId()}
	}

	return a.manager.GetBills(query)
}

func (a *AuthorizationDecorator) CreatePlan(command *create_plan.CreatePlanCommand) (*create_plan.CreatePlanCommandResponse, *application_specific.ApplicationException) {
	err := application_specific.AssertUserRole(command.Session, auth.RoleGymOwner)
	if err != nil {
		return nil, err
	}

	if !command.Session.IsOwnerOfEnabledGym(command.GymId) {
		return nil, application_specific.NewForbiddenException("OPERATION_NOT_ALLOWED", "Target gym disabled or not your gym", map[string]string{
			"gymId": command.GymId,
		})
	}

	return a.manager.CreatePlan(command)
}

func (a *AuthorizationDecorator) UpdatePlan(command *update_plan.UpdatePlanCommand) (*update_plan.UpdatePlanCommandResponse, *application_specific.ApplicationException) {
	err := application_specific.AssertUserRole(command.Session, auth.RoleGymOwner)
	if err != nil {
		return nil, err
	}

	plan, err := a.manager.GetPlan(&get_plan.GetPlanQuery{Id: command.Id, Session: command.Session})
	if err != nil {
		return nil, err
	}

	if !command.Session.IsOwnerOfEnabledGym(plan.Gym.Id) {
		return nil, application_specific.NewForbiddenException("OPERATION_NOT_ALLOWED", "Target gym disabled or not your gym", map[string]string{
			"planId": plan.Id,
			"gymId":  plan.Gym.Id,
		})
	}

	return a.manager.UpdatePlan(command)
}

func (a *AuthorizationDecorator) DeletePlan(command *delete_plan.DeletePlanCommand) (*delete_plan.DeletePlanCommandResponse, *application_specific.ApplicationException) {
	err := application_specific.AssertUserRole(command.Session, auth.RoleGymOwner)
	if err != nil {
		return nil, err
	}

	plan, err := a.manager.GetPlan(&get_plan.GetPlanQuery{Id: command.Id, Session: command.Session})
	if err != nil {
		return nil, err
	}

	if !command.Session.IsOwnerOfEnabledGym(plan.Gym.Id) {
		return nil, application_specific.NewForbiddenException("OPERATION_NOT_ALLOWED", "Target gym disabled or not your gym", map[string]string{
			"planId": plan.Id,
			"gymId":  plan.Gym.Id,
		})
	}

	return a.manager.DeletePlan(command)
}

func (a *AuthorizationDecorator) GetPlan(query *get_plan.GetPlanQuery) (*get_plan.GetPlanQueryResponse, *application_specific.ApplicationException) {
	err := application_specific.AssertUserRole(query.Session, auth.RoleGymOwner, auth.RoleCustomer, auth.RoleAdmin)
	if err != nil {
		return nil, err
	}

	plan, err := a.manager.GetPlan(query)
	if err != nil {
		return nil, err
	}

	if query.Session.RoleIsOneOf(auth.RoleGymOwner) && !query.Session.IsOwnerOfEnabledGym(plan.Gym.Id) {
		return nil, application_specific.NewForbiddenException("OPERATION_NOT_ALLOWED", "Target gym disabled or not your gym", map[string]string{
			"planId": query.Id,
			"gymId":  plan.Gym.Id,
		})
	}

	return a.manager.GetPlan(query)
}

func (a *AuthorizationDecorator) GetPlans(query *get_plans.GetPlansQuery) (*get_plans.GetPlansQueryResponse, *application_specific.ApplicationException) {
	err := application_specific.AssertUserRole(query.Session, auth.RoleGymOwner, auth.RoleCustomer, auth.RoleAdmin)
	if err != nil {
		return nil, err
	}

	if query.Session.RoleIsOneOf(auth.RoleGymOwner) {
		query.GymId = query.Session.User.Profile.EnabledOwnedGyms
	}

	return a.manager.GetPlans(query)
}
