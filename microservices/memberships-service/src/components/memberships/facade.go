package memberships

import (
	"gym-management-memberships/src/components/memberships/core/domain"
	"gym-management-memberships/src/components/memberships/core/usecases/bills/get_bill"
	"gym-management-memberships/src/components/memberships/core/usecases/bills/get_bills"
	"gym-management-memberships/src/components/memberships/core/usecases/bills/mark_bill_as_paid"
	"gym-management-memberships/src/components/memberships/core/usecases/customers/change_customer_plan"
	"gym-management-memberships/src/components/memberships/core/usecases/customers/create_customer"
	"gym-management-memberships/src/components/memberships/core/usecases/customers/delete_customer"
	"gym-management-memberships/src/components/memberships/core/usecases/customers/get_customer"
	"gym-management-memberships/src/components/memberships/core/usecases/customers/get_customers"
	"gym-management-memberships/src/components/memberships/core/usecases/customers/restrict_customer"
	"gym-management-memberships/src/components/memberships/core/usecases/customers/unrestrict_customer"
	"gym-management-memberships/src/components/memberships/core/usecases/customers/update_customer"
	"gym-management-memberships/src/components/memberships/core/usecases/memberships/cancel_membership"
	"gym-management-memberships/src/components/memberships/core/usecases/memberships/get_membership"
	"gym-management-memberships/src/components/memberships/core/usecases/memberships/get_membership_badge"
	"gym-management-memberships/src/components/memberships/core/usecases/memberships/get_memberships"
	"gym-management-memberships/src/components/memberships/core/usecases/memberships/renew_membership"
	"gym-management-memberships/src/components/memberships/core/usecases/plans/create_plan"
	"gym-management-memberships/src/components/memberships/core/usecases/plans/delete_plan"
	"gym-management-memberships/src/components/memberships/core/usecases/plans/get_plan"
	"gym-management-memberships/src/components/memberships/core/usecases/plans/get_plans"
	"gym-management-memberships/src/components/memberships/core/usecases/plans/update_plan"
	"gym-management-memberships/src/components/memberships/core/usecases/training_sessions/end_training_session"
	"gym-management-memberships/src/components/memberships/core/usecases/training_sessions/get_training_session"
	"gym-management-memberships/src/components/memberships/core/usecases/training_sessions/get_training_sessions"
	"gym-management-memberships/src/components/memberships/core/usecases/training_sessions/start_training_session"
	"gym-management-memberships/src/lib/primitives/application_specific"
)

type Facade struct {
	CustomerRepository   domain.CustomerRepository
	PlanRepository       domain.PlanRepository
	MembershipRepository domain.MembershipRepository
	EmailsService        domain.EmailService
	EventsPublisher      domain.EventsPublisher
}

func (f *Facade) CreateCustomer(command *create_customer.CreateCustomerCommand) (*create_customer.CreateCustomerCommandResponse, *application_specific.ApplicationException) {
	handler := &create_customer.CreateCustomerCommandHandler{
		CustomerRepository:   f.CustomerRepository,
		PlanRepository:       f.PlanRepository,
		MembershipRepository: f.MembershipRepository,
		EmailsService:        f.EmailsService,
		EventsPublisher:      f.EventsPublisher,
	}

	return handler.Handle(command)
}

func (f *Facade) UpdateCustomer(command *update_customer.UpdateCustomerCommand) (*update_customer.UpdateCustomerCommandResponse, *application_specific.ApplicationException) {
	handler := &update_customer.UpdateCustomerCommandHandler{
		CustomerRepository: f.CustomerRepository,
		EmailsService:      f.EmailsService,
		EventsPublisher:    f.EventsPublisher,
	}

	return handler.Handle(command)
}

func (f *Facade) RestrictCustomer(command *restrict_customer.RestrictCustomerCommand) (*restrict_customer.RestrictCustomerCommandResponse, *application_specific.ApplicationException) {
	handler := &restrict_customer.RestrictCustomerCommandHandler{
		CustomerRepository:   f.CustomerRepository,
		MembershipRepository: f.MembershipRepository,
		EventsPublisher:      f.EventsPublisher,
	}

	return handler.Handle(command)
}

func (f *Facade) UnrestrictCustomer(command *unrestrict_customer.UnrestrictCustomerCommand) (*unrestrict_customer.UnrestrictCustomerCommandResponse, *application_specific.ApplicationException) {
	handler := &unrestrict_customer.UnrestrictCustomerCommandHandler{
		CustomerRepository: f.CustomerRepository,
		EventsPublisher:    f.EventsPublisher,
	}

	return handler.Handle(command)
}

func (f *Facade) DeleteCustomer(command *delete_customer.DeleteCustomerCommand) (*delete_customer.DeleteCustomerCommandResponse, *application_specific.ApplicationException) {
	handler := &delete_customer.DeleteCustomerCommandHandler{
		CustomerRepository:   f.CustomerRepository,
		MembershipRepository: f.MembershipRepository,
		EventsPublisher:      f.EventsPublisher,
	}

	return handler.Handle(command)
}

func (f *Facade) ChangeCustomerPlan(command *change_customer_plan.ChangeCustomerPlanCommand) (*change_customer_plan.ChangeCustomerPlanCommandResponse, *application_specific.ApplicationException) {
	handler := &change_customer_plan.ChangeCustomerPlanCommandHandler{
		CustomerRepository:   f.CustomerRepository,
		PlanRepository:       f.PlanRepository,
		MembershipRepository: f.MembershipRepository,
		EventsPublisher:      f.EventsPublisher,
	}

	return handler.Handle(command)
}

func (f *Facade) GetCustomer(query *get_customer.GetCustomerQuery) (*get_customer.GetCustomerQueryResponse, *application_specific.ApplicationException) {
	handler := &get_customer.GetCustomerQueryHandler{}

	return handler.Handle(query)
}

func (f *Facade) GetCustomers(query *get_customers.GetCustomersQuery) (*get_customers.GetCustomersQueryResponse, *application_specific.ApplicationException) {
	handler := &get_customers.GetCustomersQueryHandler{}

	return handler.Handle(query)
}

func (f *Facade) CancelMembership(command *cancel_membership.CancelMembershipCommand) (*cancel_membership.CancelMembershipCommandResponse, *application_specific.ApplicationException) {
	handler := &cancel_membership.CancelMembershipCommandHandler{
		MembershipRepository: f.MembershipRepository,
		EventPublisher:       f.EventsPublisher,
	}

	return handler.Handle(command)
}

func (f *Facade) RenewMembership(command *renew_membership.RenewMembershipCommand) (*renew_membership.RenewMembershipCommandResponse, *application_specific.ApplicationException) {
	handler := &renew_membership.RenewMembershipCommandHandler{
		MembershipRepository: f.MembershipRepository,
		EventPublisher:       f.EventsPublisher,
	}

	return handler.Handle(command)
}

func (f *Facade) GetMembership(query *get_membership.GetMembershipQuery) (*get_membership.GetMembershipQueryResponse, *application_specific.ApplicationException) {
	handler := &get_membership.GetMembershipQueryHandler{}

	return handler.Handle(query)
}

func (f *Facade) GetMembershipBadge(query *get_membership_badge.GetMembershipBadgeQuery) (*get_membership_badge.GetMembershipBadgeQueryResponse, *application_specific.ApplicationException) {
	handler := &get_membership_badge.GetMembershipBadgeQueryHandler{}

	return handler.Handle(query)
}

func (f *Facade) GetMemberships(query *get_memberships.GetMembershipsQuery) (*get_memberships.GetMembershipsQueryResponse, *application_specific.ApplicationException) {
	handler := &get_memberships.GetMembershipsQueryHandler{}

	return handler.Handle(query)
}

func (f *Facade) StartTrainingSession(command *start_training_session.StartTrainingSessionCommand) (*start_training_session.StartTrainingSessionCommandResponse, *application_specific.ApplicationException) {
	handler := &start_training_session.StartTrainingSessionCommandHandler{
		MembershipRepository: f.MembershipRepository,
		EventsPublisher:      f.EventsPublisher,
	}

	return handler.Handle(command)
}

func (f *Facade) EndTrainingSession(command *end_training_session.EndTrainingSessionCommand) (*end_training_session.EndTrainingSessionCommandResponse, *application_specific.ApplicationException) {
	handler := &end_training_session.EndTrainingSessionCommandHandler{
		MembershipRepository: f.MembershipRepository,
		EventsPublisher:      f.EventsPublisher,
	}

	return handler.Handle(command)
}

func (f *Facade) GetTrainingSession(query *get_training_session.GetTrainingSessionQuery) (*get_training_session.GetTrainingSessionQueryResponse, *application_specific.ApplicationException) {
	handler := &get_training_session.GetTrainingSessionQueryHandler{}

	return handler.Handle(query)
}

func (f *Facade) GetTrainingSessions(query *get_training_sessions.GetTrainingSessionsQuery) (*get_training_sessions.GetTrainingSessionsQueryResponse, *application_specific.ApplicationException) {
	handler := &get_training_sessions.GetTrainingSessionsQueryHandler{}

	return handler.Handle(query)
}

func (f *Facade) MarkBillAsPaid(command *mark_bill_as_paid.MarkBillAsPaidCommand) (*mark_bill_as_paid.MarkBillAsPaidCommandResponse, *application_specific.ApplicationException) {
	handler := &mark_bill_as_paid.MarkBillAsPaidCommandHandler{
		MembershipRepository: f.MembershipRepository,
		EventPublisher:       f.EventsPublisher,
	}

	return handler.Handle(command)
}

func (f *Facade) GetBill(query *get_bill.GetBillQuery) (*get_bill.GetBillQueryResponse, *application_specific.ApplicationException) {
	handler := &get_bill.GetBillQueryHandler{}

	return handler.Handle(query)
}

func (f *Facade) GetBills(query *get_bills.GetBillsQuery) (*get_bills.GetBillsQueryResponse, *application_specific.ApplicationException) {
	handler := &get_bills.GetBillsQueryHandler{}

	return handler.Handle(query)
}

func (f *Facade) CreatePlan(command *create_plan.CreatePlanCommand) (*create_plan.CreatePlanCommandResponse, *application_specific.ApplicationException) {
	handler := &create_plan.CreatePlanCommandHandler{
		PlanRepository:  f.PlanRepository,
		EventsPublisher: f.EventsPublisher,
	}

	return handler.Handle(command)
}

func (f *Facade) UpdatePlan(command *update_plan.UpdatePlanCommand) (*update_plan.UpdatePlanCommandResponse, *application_specific.ApplicationException) {
	handler := &update_plan.UpdatePlanCommandHandler{
		PlanRepository:  f.PlanRepository,
		EventsPublisher: f.EventsPublisher,
	}

	return handler.Handle(command)
}

func (f *Facade) DeletePlan(command *delete_plan.DeletePlanCommand) (*delete_plan.DeletePlanCommandResponse, *application_specific.ApplicationException) {
	handler := &delete_plan.DeletePlanCommandHandler{
		PlanRepository:  f.PlanRepository,
		EventsPublisher: f.EventsPublisher,
	}

	return handler.Handle(command)
}

func (f *Facade) GetPlan(query *get_plan.GetPlanQuery) (*get_plan.GetPlanQueryResponse, *application_specific.ApplicationException) {
	handler := &get_plan.GetPlanQueryHandler{}

	return handler.Handle(query)
}

func (f *Facade) GetPlans(query *get_plans.GetPlansQuery) (*get_plans.GetPlansQueryResponse, *application_specific.ApplicationException) {
	handler := &get_plans.GetPlansQueryHandler{}

	return handler.Handle(query)
}
