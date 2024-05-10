package memberships

import (
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

type Manager interface {
	CreateCustomer(command *create_customer.CreateCustomerCommand) (*create_customer.CreateCustomerCommandResponse, *application_specific.ApplicationException)
	UpdateCustomer(command *update_customer.UpdateCustomerCommand) (*update_customer.UpdateCustomerCommandResponse, *application_specific.ApplicationException)
	RestrictCustomer(command *restrict_customer.RestrictCustomerCommand) (*restrict_customer.RestrictCustomerCommandResponse, *application_specific.ApplicationException)
	UnrestrictCustomer(command *unrestrict_customer.UnrestrictCustomerCommand) (*unrestrict_customer.UnrestrictCustomerCommandResponse, *application_specific.ApplicationException)
	DeleteCustomer(command *delete_customer.DeleteCustomerCommand) (*delete_customer.DeleteCustomerCommandResponse, *application_specific.ApplicationException)
	ChangeCustomerPlan(command *change_customer_plan.ChangeCustomerPlanCommand) (*change_customer_plan.ChangeCustomerPlanCommandResponse, *application_specific.ApplicationException)
	GetCustomer(query *get_customer.GetCustomerQuery) (*get_customer.GetCustomerQueryResponse, *application_specific.ApplicationException)
	GetCustomers(query *get_customers.GetCustomersQuery) (*get_customers.GetCustomersQueryResponse, *application_specific.ApplicationException)

	CancelMembership(command *cancel_membership.CancelMembershipCommand) (*cancel_membership.CancelMembershipCommandResponse, *application_specific.ApplicationException)
	RenewMembership(command *renew_membership.RenewMembershipCommand) (*renew_membership.RenewMembershipCommandResponse, *application_specific.ApplicationException)
	GetMembership(query *get_membership.GetMembershipQuery) (*get_membership.GetMembershipQueryResponse, *application_specific.ApplicationException)
	GetMembershipBadge(query *get_membership_badge.GetMembershipBadgeQuery) (*get_membership_badge.GetMembershipBadgeQueryResponse, *application_specific.ApplicationException)
	GetMemberships(query *get_memberships.GetMembershipsQuery) (*get_memberships.GetMembershipsQueryResponse, *application_specific.ApplicationException)

	StartTrainingSession(command *start_training_session.StartTrainingSessionCommand) (*start_training_session.StartTrainingSessionCommandResponse, *application_specific.ApplicationException)
	EndTrainingSession(command *end_training_session.EndTrainingSessionCommand) (*end_training_session.EndTrainingSessionCommandResponse, *application_specific.ApplicationException)
	GetTrainingSession(query *get_training_session.GetTrainingSessionQuery) (*get_training_session.GetTrainingSessionQueryResponse, *application_specific.ApplicationException)
	GetTrainingSessions(query *get_training_sessions.GetTrainingSessionsQuery) (*get_training_sessions.GetTrainingSessionsQueryResponse, *application_specific.ApplicationException)

	MarkBillAsPaid(command *mark_bill_as_paid.MarkBillAsPaidCommand) (*mark_bill_as_paid.MarkBillAsPaidCommandResponse, *application_specific.ApplicationException)
	GetBill(query *get_bill.GetBillQuery) (*get_bill.GetBillQueryResponse, *application_specific.ApplicationException)
	GetBills(query *get_bills.GetBillsQuery) (*get_bills.GetBillsQueryResponse, *application_specific.ApplicationException)

	CreatePlan(command *create_plan.CreatePlanCommand) (*create_plan.CreatePlanCommandResponse, *application_specific.ApplicationException)
	UpdatePlan(command *update_plan.UpdatePlanCommand) (*update_plan.UpdatePlanCommandResponse, *application_specific.ApplicationException)
	DeletePlan(command *delete_plan.DeletePlanCommand) (*delete_plan.DeletePlanCommandResponse, *application_specific.ApplicationException)
	GetPlan(query *get_plan.GetPlanQuery) (*get_plan.GetPlanQueryResponse, *application_specific.ApplicationException)
	GetPlans(query *get_plans.GetPlansQuery) (*get_plans.GetPlansQueryResponse, *application_specific.ApplicationException)
}
