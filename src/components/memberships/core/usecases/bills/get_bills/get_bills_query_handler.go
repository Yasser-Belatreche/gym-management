package get_bills

import (
	bills2 "gym-management/src/components/memberships/core/usecases/bills"
	"gym-management/src/lib"
	"gym-management/src/lib/persistence/psql/gorm/models"
	"gym-management/src/lib/primitives/application_specific"
)

type GetBillsQueryHandler struct{}

func (h *GetBillsQueryHandler) Handle(query *GetBillsQuery) (*GetBillsQueryResponse, *application_specific.ApplicationException) {
	db := lib.GormDB(query.Session)

	options := application_specific.NewPaginationOptions(&query.PaginatedQuery)

	var bills []models.Bill

	dbQuery := db.Model(&models.Bill{})
	dbQuery = dbQuery.Joins("Membership").Select("")
	dbQuery = dbQuery.Joins("Membership.Plan").Select("plans.id", "plans.name", "plans.gym_id")
	dbQuery = dbQuery.Joins("Membership.Customer").Select("customers.id", "customers.first_name", "customers.last_name")
	//dbQuery = dbQuery.Joins("JOIN memberships ON memberships.id = bills.membership_id")
	//dbQuery = dbQuery.Joins("JOIN plans ON plans.id = memberships.plan_id")
	//dbQuery = dbQuery.Joins("JOIN customers ON customers.id = memberships.customer_id")

	if len(query.Id) > 0 {
		dbQuery = dbQuery.Where("id IN (?)", query.Id)
	}

	if len(query.MembershipId) > 0 {
		dbQuery = dbQuery.Where("membership_id IN (?)", query.MembershipId)
	}

	if len(query.CustomerId) > 0 {
		dbQuery = dbQuery.Where("memberships.customer_id IN (?)", query.CustomerId)
	}

	if len(query.PlanId) > 0 {
		dbQuery = dbQuery.Where("memberships.plan_id IN (?)", query.PlanId)
	}

	if len(query.GymId) > 0 {
		dbQuery = dbQuery.Where("plans.gym_id IN (?)", query.GymId)
	}

	if query.Paid != nil {
		dbQuery = dbQuery.Where("paid = ?", *query.Paid)
	}

	result := dbQuery.Offset(options.Skip).Limit(options.PerPage).Order("updated_at DESC").Find(&bills)
	if result.Error != nil {
		return nil, application_specific.NewUnknownException("BILLS.FAILED_TO_GET_BILLS", result.Error.Error(), nil)
	}

	var total int64
	result = dbQuery.Count(&total)
	if result.Error != nil {
		return nil, application_specific.NewUnknownException("BILLS.FAILED_TO_GET_BILLS", result.Error.Error(), nil)
	}

	response := GetBillsQueryResponse(application_specific.NewPaginatedResponse(options, total, bills, func(item models.Bill) bills2.BillToReturn {
		return bills2.BillToReturn{
			Id:     item.Id,
			Amount: item.Amount,
			Paid:   item.Paid,
			Customer: bills2.BillToReturnCustomer{
				Id:        item.Membership.Customer.Id,
				FirstName: item.Membership.Customer.FirstName,
				LastName:  item.Membership.Customer.LastName,
			},
			Plan: bills2.BillToReturnPlan{
				Id:   item.Membership.Plan.Id,
				Name: item.Membership.Plan.Name,
			},
			MembershipId: item.MembershipId,
			GymId:        item.Membership.Plan.GymId,
			PaidAt:       item.PaidAt,
			DueDate:      item.DueTo,
			CreatedAt:    item.CreatedAt,
		}
	}))

	return &response, nil
}
