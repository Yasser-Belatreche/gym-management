package get_memberships

import (
	"gym-management-memberships/src/components/memberships/core/usecases/memberships"
	"gym-management-memberships/src/lib"
	"gym-management-memberships/src/lib/primitives/application_specific"
)

type GetMembershipsQueryHandler struct{}

func (h *GetMembershipsQueryHandler) Handle(query *GetMembershipsQuery) (*GetMembershipsQueryResponse, *application_specific.ApplicationException) {
	db := lib.GormDB(query.Session)

	options := application_specific.NewPaginationOptions(&query.PaginatedQuery)

	var list []memberships.MembershipToReturn

	dbQuery := db.Table("memberships as m").
		Select(`
			m.id as id, 
			m.start_date as start_date, 
			m.end_date as end_date, 
			m.enabled as enabled, 
			m.disabled_for as disabled_for, 
			m.sessions_per_week as sessions_per_week, 
			m.with_coach as with_coach, 
			m.monthly_price as monthly_price, 
			m.created_by as created_by, 
			m.created_at as created_at, 
			m.updated_by as updated_by, 
			m.updated_at as updated_at, 
			m.renewed_at as renewed_at,

			c.id as customer_id,
			c.first_name as customer_first_name, 
			c.last_name as customer_last_name,

			p.id as plan_id,
			p.gym_id as gym_id,
			p.name as plan_name
`).
		Joins("INNER JOIN plans as p on p.id = m.plan_id").
		Joins("INNER JOIN customers as c on c.id = m.customer_id")

	if len(query.Id) > 0 {
		dbQuery.Where("m.id IN (?)", query.Id)
	}

	if len(query.CustomerId) > 0 {
		dbQuery.Where("m.customer_id IN (?)", query.CustomerId)
	}

	if len(query.PlanId) > 0 {
		dbQuery.Where("m.plan_id IN (?)", query.PlanId)
	}

	if len(query.GymId) > 0 {
		dbQuery.Where("p.gym_id IN (?)", query.GymId)
	}

	if query.Enabled != nil {
		dbQuery.Where("m.enabled = ?", *query.Enabled)
	}

	err := dbQuery.
		Offset(options.Skip).
		Limit(options.PerPage).
		Order("m.updated_at DESC").
		Find(&list).
		Error
	if err != nil {
		return nil, application_specific.NewUnknownException("MEMBERSHIPS.FAILED_TO_GET_MEMBERSHIPS", err.Error(), nil)
	}

	var total int64
	err = dbQuery.Count(&total).Error
	if err != nil {
		return nil, application_specific.NewUnknownException("MEMBERSHIPS.FAILED_TO_GET_MEMBERSHIPS", err.Error(), nil)
	}

	response := GetMembershipsQueryResponse(application_specific.NewPaginatedResponse(options, total, list, func(item memberships.MembershipToReturn) memberships.MembershipToReturn {
		return item
	}))

	return &response, nil
}
