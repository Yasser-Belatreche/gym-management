package get_membership_badge

import (
	"errors"
	"gorm.io/gorm"
	"gym-management/src/components/memberships/core/usecases/memberships"
	"gym-management/src/lib"
	"gym-management/src/lib/persistence/psql/gorm/models"
	"gym-management/src/lib/primitives/application_specific"
)

type GetMembershipBadgeQueryHandler struct {
}

func (h *GetMembershipBadgeQueryHandler) Handle(query *GetMembershipBadgeQuery) (*GetMembershipBadgeQueryResponse, *application_specific.ApplicationException) {
	db := lib.GormDB(query.Session)

	var membership models.Membership

	dbQuery := db.Model(&models.Membership{})
	dbQuery = dbQuery.Joins("Plan").Select("plans.id, plans.name, plans.gym_id")
	dbQuery = dbQuery.Joins("Customer").Select("customers.id, customers.first_name, customer.last_name")

	if err := dbQuery.Where("id = ?", query.MembershipId).First(&membership).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, application_specific.NewNotFoundException("MEMBERSHIPS.NOT_FOUND", "membership not found", map[string]string{
				"id": query.MembershipId,
			})
		}

		return nil, application_specific.NewUnknownException("MEMBERSHIPS.FAILED_TO_GET_MEMBERSHIP", err.Error(), nil)
	}

	response := GetMembershipBadgeQueryResponse(
		memberships.MembershipToReturn{
			Id:              membership.Id,
			StartDate:       membership.StartDate,
			EndDate:         membership.EndDate,
			Enabled:         membership.Enabled,
			DisabledFor:     membership.DisabledFor,
			SessionsPerWeek: membership.SessionsPerWeek,
			WithCoach:       membership.WithCoach,
			MonthlyPrice:    membership.MonthlyPrice,
			Customer: memberships.MembershipToReturnCustomer{
				Id:        membership.Customer.Id,
				FirstName: membership.Customer.FirstName,
				LastName:  membership.Customer.LastName,
			},
			Plan: memberships.MembershipToReturnPlan{
				Id:   membership.Plan.Id,
				Name: membership.Plan.Name,
			},
			GymId:     membership.Plan.GymId,
			CreatedBy: membership.CreatedBy,
			CreatedAt: membership.CreatedAt,
			UpdatedBy: membership.UpdatedBy,
			UpdatedAt: membership.UpdatedAt,
			RenewedAt: membership.RenewedAt,
		})

	return &response, nil
}
