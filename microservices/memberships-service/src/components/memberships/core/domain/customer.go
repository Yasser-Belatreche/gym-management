package domain

import (
	"gym-management/src/components/memberships/core/domain/events"
	"gym-management/src/lib/primitives/application_specific"
	"gym-management/src/lib/primitives/generic"
	"strconv"
	"strings"
	"time"
)

type Customer struct {
	id          string
	firstName   string
	lastName    string
	phoneNumber string
	email       application_specific.Email
	restricted  bool
	birthYear   int
	gender      Gender
	createdBy   string
	updatedBy   string
	deletedBy   *string
	deletedAt   *time.Time
	events      []*events.MembershipEvent[interface{}]
}

type CustomerState struct {
	Id          string
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Restricted  bool
	BirthYear   int
	Gender      string
	CreatedBy   string
	UpdatedBy   string
	DeletedBy   *string
	DeletedAt   *time.Time
}

func createCustomer(firstName string, lastName string, phoneNumber string, email application_specific.Email, password string, birthYear int, gender Gender, createdBy string, planId string) (*Customer, *application_specific.ApplicationException) {
	firstName = strings.TrimSpace(firstName)
	lastName = strings.TrimSpace(lastName)
	phoneNumber = strings.TrimSpace(phoneNumber)
	password = strings.TrimSpace(password)

	if len(firstName) == 0 {
		return nil, application_specific.NewValidationException("MEMBERSHIPS.CUSTOMERS.INVALID_FIRST_NAME", "First name cannot be empty", map[string]string{
			"firstName": firstName,
		})
	}
	if len(lastName) == 0 {
		return nil, application_specific.NewValidationException("MEMBERSHIPS.CUSTOMERS.INVALID_LAST_NAME", "Last name cannot be empty", map[string]string{
			"lastName": lastName,
		})
	}
	if len(phoneNumber) == 0 {
		return nil, application_specific.NewValidationException("MEMBERSHIPS.CUSTOMERS.INVALID_PHONE_NUMBER", "Phone number cannot be empty", map[string]string{
			"phoneNumber": phoneNumber,
		})
	}
	if len(password) < 6 {
		return nil, application_specific.NewValidationException("MEMBERSHIPS.CUSTOMERS.INVALID_PASSWORD", "password must be at least 6 characters long", map[string]string{
			"password": password,
		})
	}
	if birthYear < 1900 || birthYear > time.Now().Year() {
		return nil, application_specific.NewValidationException("MEMBERSHIPS.CUSTOMERS.INVALID_BIRTH_YEAR", "Birth year must be between 1900 and the current year", map[string]string{
			"birthYear": strconv.Itoa(birthYear),
		})
	}

	customer := &Customer{
		id:          generic.GenerateULID(),
		firstName:   firstName,
		lastName:    lastName,
		phoneNumber: phoneNumber,
		email:       email,
		restricted:  false,
		birthYear:   birthYear,
		gender:      gender,
		createdBy:   createdBy,
		updatedBy:   createdBy,
		deletedBy:   nil,
		deletedAt:   nil,
		events:      make([]*events.MembershipEvent[interface{}], 0),
	}

	customer.pushEvents(NewCustomerCreatedEvent(customer.State(), password, planId))

	return customer, nil
}

func CustomerFromState(state *CustomerState) *Customer {
	email, _ := application_specific.NewEmail(state.Email)

	return &Customer{
		id:          state.Id,
		firstName:   state.FirstName,
		lastName:    state.LastName,
		phoneNumber: state.PhoneNumber,
		email:       email,
		restricted:  state.Restricted,
		birthYear:   state.BirthYear,
		gender:      Gender(state.Gender),
		createdBy:   state.CreatedBy,
		updatedBy:   state.UpdatedBy,
		deletedBy:   state.DeletedBy,
		deletedAt:   state.DeletedAt,
		events:      make([]*events.MembershipEvent[interface{}], 0),
	}
}

func (c *Customer) Update(firstName string, lastName string, phoneNumber string, email application_specific.Email, birthYear int, gender Gender, newPassword *string, updatedBy string) *application_specific.ApplicationException {
	if c.restricted {
		return application_specific.NewValidationException("MEMBERSHIPS.CUSTOMERS.ALREADY_RESTRICTED", "Customer is already restricted", map[string]string{
			"customerId": c.id,
		})
	}
	if c.deletedAt != nil {
		return application_specific.NewNotFoundException("MEMBERSHIPS.CUSTOMERS.DELETED", "Customer is deleted", map[string]string{
			"customerId": c.id,
		})
	}

	firstName = strings.TrimSpace(firstName)
	lastName = strings.TrimSpace(lastName)
	phoneNumber = strings.TrimSpace(phoneNumber)

	if len(firstName) == 0 {
		return application_specific.NewValidationException("MEMBERSHIPS.CUSTOMERS.INVALID_FIRST_NAME", "First name cannot be empty", map[string]string{
			"firstName": firstName,
		})
	}
	if len(lastName) == 0 {
		return application_specific.NewValidationException("MEMBERSHIPS.CUSTOMERS.INVALID_LAST_NAME", "Last name cannot be empty", map[string]string{
			"lastName": lastName,
		})
	}
	if len(phoneNumber) == 0 {
		return application_specific.NewValidationException("MEMBERSHIPS.CUSTOMERS.INVALID_PHONE_NUMBER", "Phone number cannot be empty", map[string]string{
			"phoneNumber": phoneNumber,
		})
	}
	if birthYear < 1900 || birthYear > time.Now().Year() {
		return application_specific.NewValidationException("MEMBERSHIPS.CUSTOMERS.INVALID_BIRTH_YEAR", "Birth year must be between 1900 and the current year", map[string]string{
			"birthYear": strconv.Itoa(birthYear),
		})
	}
	if newPassword != nil && len(*newPassword) < 6 {
		return application_specific.NewValidationException("MEMBERSHIPS.CUSTOMERS.INVALID_PASSWORD", "password must be at least 6 characters long", map[string]string{
			"password": *newPassword,
		})
	}

	c.gender = gender
	c.firstName = firstName
	c.lastName = lastName
	c.phoneNumber = phoneNumber
	c.email = email
	c.birthYear = birthYear
	c.updatedBy = updatedBy

	c.pushEvents(NewCustomerUpdatedEvent(c.State(), newPassword))

	return nil
}

func (c *Customer) restrict(by string) *application_specific.ApplicationException {
	if c.restricted {
		return application_specific.NewValidationException("MEMBERSHIPS.CUSTOMERS.ALREADY_RESTRICTED", "Customer is already restricted", map[string]string{
			"customerId": c.id,
		})
	}
	if c.deletedAt != nil {
		return application_specific.NewNotFoundException("MEMBERSHIPS.CUSTOMERS.DELETED", "Customer is deleted", map[string]string{
			"customerId": c.id,
		})
	}

	c.restricted = true
	c.updatedBy = by

	c.pushEvents(NewCustomerRestrictedEvent(c.State()), NewCustomerUpdatedEvent(c.State(), nil))

	return nil
}

func (c *Customer) Unrestrict(by string) *application_specific.ApplicationException {
	if !c.restricted {
		return application_specific.NewValidationException("MEMBERSHIPS.CUSTOMERS.NOT_RESTRICTED", "Customer is not restricted", map[string]string{
			"customerId": c.id,
		})
	}
	if c.deletedAt != nil {
		return application_specific.NewNotFoundException("MEMBERSHIPS.CUSTOMERS.DELETED", "Customer is deleted", map[string]string{
			"customerId": c.id,
		})
	}

	c.restricted = false
	c.updatedBy = by

	c.pushEvents(NewCustomerUnrestrictedEvent(c.State()), NewCustomerUpdatedEvent(c.State(), nil))

	return nil
}

func (c *Customer) EmailIs(email application_specific.Email) bool {
	return c.email.Equals(email)
}

func (c *Customer) delete(by string) *application_specific.ApplicationException {
	if c.deletedAt != nil {
		return application_specific.NewValidationException("MEMBERSHIPS.CUSTOMERS.ALREADY_DELETED", "Customer is already deleted", map[string]string{
			"customerId": c.id,
		})
	}

	c.deletedBy = &by
	now := time.Now()
	c.deletedAt = &now

	c.pushEvents(NewCustomerDeletedEvent(c.State()))

	return nil
}

func (c *Customer) changePlan(plan *Plan, by string) *application_specific.ApplicationException {
	if c.deletedAt != nil {
		return application_specific.NewNotFoundException("MEMBERSHIPS.CUSTOMERS.DELETED", "Customer is deleted", map[string]string{
			"customerId": c.id,
		})
	}
	if c.restricted {
		return application_specific.NewValidationException("MEMBERSHIPS.CUSTOMERS.RESTRICTED", "Customer is restricted", map[string]string{
			"customerId": c.id,
		})
	}

	c.updatedBy = by

	c.pushEvents(NewCustomerPlanChangedEvent(c.State(), plan.id))

	return nil
}

func (c *Customer) State() *CustomerState {
	return &CustomerState{
		Id:          c.id,
		FirstName:   c.firstName,
		LastName:    c.lastName,
		Email:       c.email.Value,
		PhoneNumber: c.phoneNumber,
		Restricted:  c.restricted,
		BirthYear:   c.birthYear,
		Gender:      c.gender.Value(),
		CreatedBy:   c.createdBy,
		UpdatedBy:   c.updatedBy,
		DeletedBy:   c.deletedBy,
		DeletedAt:   c.deletedAt,
	}
}

func (c *Customer) pushEvents(events ...*events.MembershipEvent[interface{}]) {
	for _, event := range events {
		c.events = append(c.events, event)
	}
}

func (c *Customer) PullEvents() []*events.MembershipEvent[interface{}] {
	old := c.events

	c.events = make([]*events.MembershipEvent[interface{}], 0)

	return old
}
