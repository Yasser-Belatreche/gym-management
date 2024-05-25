package get_memberships

import (
	memberships2 "gym-management/src/components/memberships/core/usecases/memberships"
	"gym-management/src/lib"
	"gym-management/src/lib/persistence/psql/gorm/models"
	"gym-management/src/lib/primitives/application_specific"
)

type GetMembershipsQueryHandler struct{}

func (h *GetMembershipsQueryHandler) Handle(query *GetMembershipsQuery) (*GetMembershipsQueryResponse, *application_specific.ApplicationException) {
	db := lib.GormDB(query.Session)

	options := application_specific.NewPaginationOptions(&query.PaginatedQuery)

	var memberships []models.Membership

	dbQuery := db.Model(&models.Membership{})
	dbQuery = dbQuery.Joins("Plan").Select("plans.id, plans.name, plans.gym_id")
	dbQuery = dbQuery.Joins("Customer").Select("customers.id, customers.first_name, customers.last_name")

	if len(query.Id) > 0 {
		dbQuery = dbQuery.Where("id IN (?)", query.Id)
	}

	if len(query.CustomerId) > 0 {
		dbQuery = dbQuery.Where("customer_id IN (?)", query.CustomerId)
	}

	if len(query.PlanId) > 0 {
		dbQuery = dbQuery.Where("plan_id IN (?)", query.PlanId)
	}

	if len(query.GymId) > 0 {
		dbQuery = dbQuery.Where("plans.gym_id IN (?)", query.GymId)
	}

	if query.Enabled != nil {
		dbQuery = dbQuery.Where("enabled = ?", *query.Enabled)
	}

	var result = dbQuery.Offset(options.Skip).Limit(options.PerPage).Order("updated_at DESC").Find(&memberships)
	if result.Error != nil {
		return nil, application_specific.NewUnknownException("MEMBERSHIPS.FAILED_TO_GET_MEMBERSHIPS", result.Error.Error(), nil)
	}

	var total int64
	result = dbQuery.Count(&total)
	if result.Error != nil {
		return nil, application_specific.NewUnknownException("MEMBERSHIPS.FAILED_TO_GET_MEMBERSHIPS", result.Error.Error(), nil)
	}

	response := GetMembershipsQueryResponse(application_specific.NewPaginatedResponse(options, total, memberships, func(item models.Membership) memberships2.MembershipToReturn {
		return memberships2.MembershipToReturn{
			Id:              item.Id,
			StartDate:       item.StartDate,
			EndDate:         item.EndDate,
			Enabled:         item.Enabled,
			DisabledFor:     item.DisabledFor,
			SessionsPerWeek: item.SessionsPerWeek,
			WithCoach:       item.WithCoach,
			MonthlyPrice:    item.MonthlyPrice,
			Customer: memberships2.MembershipToReturnCustomer{
				Id:        item.Customer.Id,
				FirstName: item.Customer.FirstName,
				LastName:  item.Customer.LastName,
			},
			Plan: memberships2.MembershipToReturnPlan{
				Id:   item.Plan.Id,
				Name: item.Plan.Name,
			},
			GymId:     item.Plan.GymId,
			CreatedBy: item.CreatedBy,
			CreatedAt: item.CreatedAt,
			UpdatedBy: item.UpdatedBy,
			UpdatedAt: item.UpdatedAt,
			RenewedAt: item.RenewedAt,
		}
	}))

	return &response, nil
}
