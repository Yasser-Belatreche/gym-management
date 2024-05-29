package get_customer

import (
	"errors"
	"gorm.io/gorm"
	"gym-management/src/components/memberships/core/usecases/customers"
	"gym-management/src/lib"
	"gym-management/src/lib/persistence/psql/gorm/models"
	"gym-management/src/lib/primitives/application_specific"
)

type GetCustomerQueryHandler struct{}

func (h *GetCustomerQueryHandler) Handle(query *GetCustomerQuery) (*GetCustomerQueryResponse, *application_specific.ApplicationException) {
	db := lib.GormDB(query.Session)

	var customer models.Customer

	dbQuery := db.Model(&models.Customer{})
	dbQuery = dbQuery.Joins("Memberships", "ORDER BY memberships.updated_at DESC LIMIT 1")
	dbQuery = dbQuery.Joins("Memberships.Plan")

	if err := dbQuery.Where("id = ?", query.Id).First(&customer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, application_specific.NewNotFoundException("CUSTOMERS.NOT_FOUND", "customer not found", map[string]string{
				"id": query.Id,
			})
		}

		return nil, application_specific.NewUnknownException("CUSTOMERS.FAILED_TO_GET_CUSTOMER", err.Error(), nil)
	}

	response := GetCustomerQueryResponse(
		customers.CustomerToReturn{
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
			Membership: customers.CustomerToReturnMembership{
				Id:              customer.Memberships[0].Id,
				Enabled:         customer.Memberships[0].Enabled,
				SessionsPerWeek: customer.Memberships[0].SessionsPerWeek,
				WithCoach:       customer.Memberships[0].WithCoach,
				MonthlyPrice:    customer.Memberships[0].MonthlyPrice,
				Plan: customers.CustomerToReturnMembershipPlan{
					Id:   customer.Memberships[0].Plan.Id,
					Name: customer.Memberships[0].Plan.Name,
				},
			},
			GymId:     customer.Memberships[0].Plan.GymId,
			CreatedAt: customer.CreatedAt,
			UpdatedAt: customer.UpdatedAt,
			DeletedBy: customer.DeletedBy,
			DeletedAt: customer.DeletedAt,
		})

	return &response, nil
}
