package get_customer

import (
	"errors"
	"gorm.io/gorm"
	"gym-management/src/components/memberships/core/usecases/customers"
	"gym-management/src/lib"
	"gym-management/src/lib/primitives/application_specific"
)

type GetCustomerQueryHandler struct{}

func (h *GetCustomerQueryHandler) Handle(query *GetCustomerQuery) (*GetCustomerQueryResponse, *application_specific.ApplicationException) {
	db := lib.GormDB(query.Session)

	var customer customers.CustomerToReturn

	err := db.Table("memberships AS m").
		Select(`
			c.id AS id,
			c.first_name AS first_name,
			c.last_name AS last_name,
			c.email AS email,
			c.phone_number AS phone_number,
			c.restricted AS restricted,
			c.birth_year AS birth_year,
			c.gender AS gender,
			c.created_by AS created_by,
			c.updated_by AS updated_by,
			c.created_at AS created_at,
			c.updated_at AS updated_at,
			c.deleted_at AS deleted_at,
			c.deleted_by AS deleted_by,

			m.id AS membership_id,
			m.enabled AS membership_enabled,
			m.sessions_per_week AS membership_sessions_per_week,
			m.with_coach AS membership_with_coach,
			m.monthly_price AS membership_monthly_price,
			
			p.id AS membership_plan_id,
			p.name AS membership_plan_name,
			p.gym_id AS gym_id
		`).
		Joins("INNER JOIN plans AS p ON p.id = m.plan_id").
		Joins("INNER JOIN customers AS c ON c.id = m.customer_id").
		Where("m.customer_id = ?", query.Id).
		Order("m.created_at DESC").
		First(&customer).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, application_specific.NewNotFoundException("CUSTOMERS.NOT_FOUND", "customer not found", map[string]string{
				"id": query.Id,
			})
		}

		return nil, application_specific.NewUnknownException("CUSTOMERS.FAILED_TO_GET_CUSTOMER", err.Error(), nil)
	}

	response := GetCustomerQueryResponse(customer)

	return &response, nil
}
