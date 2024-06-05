package get_membership

import (
	"errors"
	"gorm.io/gorm"
	"gym-management-memberships/src/components/memberships/core/usecases/memberships"
	"gym-management-memberships/src/lib"
	"gym-management-memberships/src/lib/primitives/application_specific"
)

type GetMembershipQueryHandler struct{}

func (h *GetMembershipQueryHandler) Handle(query *GetMembershipQuery) (*GetMembershipQueryResponse, *application_specific.ApplicationException) {
	db := lib.GormDB(query.Session)

	var membership memberships.MembershipToReturn

	err := db.Table("memberships as m").
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
		Joins("INNER JOIN customers as c on c.id = m.customer_id").
		Where("m.id = ?", query.MembershipId).
		First(&membership).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, application_specific.NewNotFoundException("MEMBERSHIPS.NOT_FOUND", "membership not found", map[string]string{
				"id": query.MembershipId,
			})
		}

		return nil, application_specific.NewUnknownException("MEMBERSHIPS.FAILED_TO_GET_MEMBERSHIP", err.Error(), nil)
	}

	response := GetMembershipQueryResponse(membership)

	return &response, nil
}
