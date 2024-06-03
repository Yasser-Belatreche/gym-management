package get_customers

import (
	"gym-management/src/components/memberships/core/usecases/customers"
	"gym-management/src/lib"
	"gym-management/src/lib/primitives/application_specific"
)

type GetCustomersQueryHandler struct{}

func (h *GetCustomersQueryHandler) Handle(query *GetCustomersQuery) (*GetCustomersQueryResponse, *application_specific.ApplicationException) {
	db := lib.GormDB(query.Session)

	options := application_specific.NewPaginationOptions(&query.PaginatedQuery)

	var list []customers.CustomerToReturn

	membershipsSubQuery := db.Table("memberships").
		Select("*").
		Order("memberships.created_at DESC").
		Limit(1)

	dbQuery := db.Table("customers AS c").
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
			p.gym_id AS gym_id,
			p.name AS membership_plan_name
		`).
		Joins("INNER JOIN (?) AS m ON m.customer_id = c.id", membershipsSubQuery).
		Joins("INNER JOIN plans AS p ON p.id = m.plan_id")

	if len(query.Id) > 0 {
		dbQuery.Where("c.id IN (?)", query.Id)
	}

	if len(query.GymId) > 0 {
		dbQuery.Where("p.gym_id IN (?)", query.GymId)
	}

	if len(query.MembershipId) > 0 {
		dbQuery.Where("m.plan_id IN (?)", query.MembershipId)
	}

	if len(query.PlanId) > 0 {
		dbQuery.Where("m.plan_id IN (?)", query.PlanId)
	}

	if query.Restricted != nil {
		dbQuery.Where("c.restricted = ?", *query.Restricted)
	}

	if query.Deleted {
		dbQuery.Where("c.deleted_at IS NOT NULL")
	} else {
		dbQuery.Where("c.deleted_at IS NULL")
	}

	err := dbQuery.
		Offset(options.Skip).
		Limit(options.PerPage).
		Order("c.updated_at DESC").
		Find(&list).
		Error

	if err != nil {
		return nil, application_specific.NewUnknownException("CUSTOMERS.FAILED_TO_GET_CUSTOMERS", err.Error(), nil)
	}

	var total int64
	err = dbQuery.Count(&total).Error
	if err != nil {
		return nil, application_specific.NewUnknownException("CUSTOMERS.FAILED_TO_GET_CUSTOMERS", err.Error(), nil)
	}

	response := GetCustomersQueryResponse(application_specific.NewPaginatedResponse(options, total, list, func(item customers.CustomerToReturn) customers.CustomerToReturn {
		return item
	}))

	return &response, nil

}
