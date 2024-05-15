package infra

import (
	"errors"
	"gorm.io/gorm"
	"gym-management/src/components/memberships/core/domain"
	"gym-management/src/lib"
	"gym-management/src/lib/persistence/psql/gorm/models"
	"gym-management/src/lib/primitives/application_specific"
)

type GormCustomerRepository struct {
}

func (g *GormCustomerRepository) FindByID(id string, session *application_specific.Session) (*domain.Customer, *application_specific.ApplicationException) {
	db := lib.GormDB(session)

	var customer models.Customer
	result := db.Set("gorm:query_option", "FOR UPDATE").First(&customer, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, application_specific.NewNotFoundException("MEMBERSHIPS.CUSTOMERS.NOT_FOUND", "Customer not found", map[string]string{
				"id": id,
			})
		}

		return nil, application_specific.NewUnknownException("MEMBERSHIPS.INFRA.FAILED_TO_FIND_CUSTOMER", result.Error.Error(), map[string]string{})
	}

	domainCustomer := customerToDomain(&customer)

	return domainCustomer, nil
}

func (g *GormCustomerRepository) Create(customer *domain.Customer, session *application_specific.Session) *application_specific.ApplicationException {
	db := lib.GormDB(session)

	customerModel := customerToDB(customer)

	result := db.Create(customerModel)
	if result.Error != nil {
		return application_specific.NewUnknownException("MEMBERSHIPS.INFRA.FAILED_TO_CREATE_CUSTOMER", result.Error.Error(), map[string]string{})
	}

	return nil
}

func (g *GormCustomerRepository) Update(customer *domain.Customer, session *application_specific.Session) *application_specific.ApplicationException {
	db := lib.GormDB(session)

	customerModel := customerToDB(customer)

	result := db.Save(customerModel)
	if result.Error != nil {
		return application_specific.NewUnknownException("MEMBERSHIPS.INFRA.FAILED_TO_UPDATE_CUSTOMER", result.Error.Error(), map[string]string{
			"id": customerModel.Id,
		})
	}

	return nil
}

func customerToDB(customer *domain.Customer) *models.Customer {
	state := customer.State()

	return &models.Customer{
		Id:          state.Id,
		FirstName:   state.FirstName,
		LastName:    state.LastName,
		PhoneNumber: state.PhoneNumber,
		Email:       state.Email,
		BirthYear:   state.BirthYear,
		Gender:      state.Gender,
		Restricted:  state.Restricted,
		CreatedBy:   state.CreatedBy,
		UpdatedBy:   state.UpdatedBy,
		DeletedBy:   state.DeletedBy,
		DeletedAt:   state.DeletedAt,
	}
}

func customerToDomain(customer *models.Customer) *domain.Customer {
	state := &domain.CustomerState{
		Id:          customer.Id,
		FirstName:   customer.FirstName,
		LastName:    customer.LastName,
		Email:       customer.Email,
		PhoneNumber: customer.PhoneNumber,
		Restricted:  customer.Restricted,
		BirthYear:   customer.BirthYear,
		Gender:      customer.Gender,
		CreatedBy:   customer.CreatedBy,
		UpdatedBy:   customer.UpdatedBy,
		DeletedBy:   customer.DeletedBy,
		DeletedAt:   customer.DeletedAt,
	}

	return domain.CustomerFromState(state)
}
