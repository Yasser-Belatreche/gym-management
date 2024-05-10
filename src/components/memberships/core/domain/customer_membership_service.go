package domain

import (
	"gym-management/src/lib/primitives/application_specific"
	"time"
)

type CustomerMembershipService struct {
	customer   *Customer
	membership *Membership
}

func NewCustomerMembershipService(customer *Customer, membership *Membership) (*CustomerMembershipService, *application_specific.ApplicationException) {
	if customer.id != membership.customerId {
		return nil, application_specific.NewValidationException("GYMS.CUSTOMERS.INVALID_MEMBERSHIP", "Customer and membership do not match", map[string]string{
			"customerId":   customer.id,
			"membershipId": membership.id,
		})
	}

	return &CustomerMembershipService{
		customer:   customer,
		membership: membership,
	}, nil
}

func CreateCustomer(firstName string, lastName string, phoneNumber string, email application_specific.Email, password string, birthYear int, gender Gender, createdBy string, membershipEndDate *time.Time, plan *Plan) (*Customer, *Membership, *application_specific.ApplicationException) {
	customer, err := createCustomer(firstName, lastName, phoneNumber, email, password, birthYear, gender, createdBy, plan.id)
	if err != nil {
		return nil, nil, err
	}

	membership, err := createMembership(membershipEndDate, plan, customer)
	if err != nil {
		return nil, nil, err
	}

	return customer, membership, nil
}

func (s *CustomerMembershipService) RestrictCustomer(by string) *application_specific.ApplicationException {
	err := s.customer.restrict(by)
	if err != nil {
		return err
	}

	if s.membership == nil {
		return nil
	}

	if !s.membership.IsDisabled() {
		err = s.membership.Cancel(by)
		if err != nil {
			return nil
		}
	}

	return nil
}

func (s *CustomerMembershipService) DeleteCustomer(by string) *application_specific.ApplicationException {
	err := s.customer.delete(by)
	if err != nil {
		return err
	}

	if !s.membership.IsDisabled() {
		err = s.membership.Cancel(by)
		if err != nil {
			return nil
		}
	}

	return nil
}

func (s *CustomerMembershipService) ChangeCustomerPlanTo(plan *Plan, endDate *time.Time, by string) (*Membership, *application_specific.ApplicationException) {
	if plan.deletedAt != nil {
		return nil, application_specific.NewValidationException("MEMBERSHIPS.PLANS.DELETED", "Plan is deleted", map[string]string{
			"id": plan.id,
		})
	}

	err := s.customer.changePlan(plan, by)
	if err != nil {
		return nil, err
	}

	newMembership, err := createMembership(endDate, plan, s.customer)
	if err != nil {
		return nil, err
	}

	if !s.membership.IsDisabled() {
		err = s.membership.Cancel(by)
		if err != nil {
			return nil, err
		}
	}

	return newMembership, nil
}
