package get_customers

import (
	customers2 "gym-management/src/components/memberships/core/usecases/customers"
	"gym-management/src/lib"
	"gym-management/src/lib/persistence/psql/gorm/models"
	"gym-management/src/lib/primitives/application_specific"
)

type GetCustomersQueryHandler struct{}

func (h *GetCustomersQueryHandler) Handle(query *GetCustomersQuery) (*GetCustomersQueryResponse, *application_specific.ApplicationException) {
	db := lib.GormDB(query.Session)

	options := application_specific.NewPaginationOptions(&query.PaginatedQuery)

	var customers []models.Customer

	dbQuery := db.Model(&models.Customer{})
	dbQuery = dbQuery.Joins("Memberships", "ORDER BY memberships.updated_at DESC LIMIT 1")
	dbQuery = dbQuery.Joins("Memberships.Plan")

	if len(query.Id) > 0 {
		dbQuery = dbQuery.Where("id IN (?)", query.Id)
	}

	if len(query.GymId) > 0 {
		dbQuery = dbQuery.Where("memberships.gym_id IN (?)", query.GymId)
	}

	if len(query.MembershipId) > 0 {
		dbQuery = dbQuery.Where("memberships.plan_id IN (?)", query.MembershipId)
	}

	if len(query.PlanId) > 0 {
		dbQuery = dbQuery.Where("memberships.plan_id IN (?)", query.PlanId)
	}

	if query.Restricted != nil {
		dbQuery = dbQuery.Where("restricted = ?", *query.Restricted)
	}

	if query.Deleted {
		dbQuery = dbQuery.Where("deleted_at IS NOT NULL")
	} else {
		dbQuery = dbQuery.Where("deleted_at IS NULL")
	}

	result := dbQuery.Offset(options.Skip).Limit(options.PerPage).Order("customers.updated_at DESC").Find(&customers)
	if result.Error != nil {
		return nil, application_specific.NewUnknownException("CUSTOMERS.FAILED_TO_GET_CUSTOMERS", result.Error.Error(), nil)
	}

	var total int64
	result = dbQuery.Count(&total)
	if result.Error != nil {
		return nil, application_specific.NewUnknownException("CUSTOMERS.FAILED_TO_GET_CUSTOMERS", result.Error.Error(), nil)
	}

	response := GetCustomersQueryResponse(application_specific.NewPaginatedResponse(options, total, customers, func(item models.Customer) customers2.CustomerToReturn {
		return customers2.CustomerToReturn{
			Id:          item.Id,
			FirstName:   item.FirstName,
			LastName:    item.LastName,
			Email:       item.Email,
			PhoneNumber: item.PhoneNumber,
			Restricted:  item.Restricted,
			BirthYear:   item.BirthYear,
			Gender:      item.Gender,
			CreatedBy:   item.CreatedBy,
			UpdatedBy:   item.UpdatedBy,
			Membership: customers2.CustomerToReturnMembership{
				Id:              item.Memberships[0].Id,
				Enabled:         item.Memberships[0].Enabled,
				SessionsPerWeek: item.Memberships[0].SessionsPerWeek,
				WithCoach:       item.Memberships[0].WithCoach,
				MonthlyPrice:    item.Memberships[0].MonthlyPrice,
				Plan: customers2.CustomerToReturnMembershipPlan{
					Id:   item.Memberships[0].Plan.Id,
					Name: item.Memberships[0].Plan.Name,
				},
			},
			GymId:     item.Memberships[0].Plan.GymId,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
			DeletedBy: item.DeletedBy,
			DeletedAt: item.DeletedAt,
		}
	}))

	return &response, nil

}
