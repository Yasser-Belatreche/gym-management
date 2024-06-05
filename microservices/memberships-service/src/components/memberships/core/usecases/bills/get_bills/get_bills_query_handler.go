package get_bills

import (
	bills2 "gym-management-memberships/src/components/memberships/core/usecases/bills"
	"gym-management-memberships/src/lib"
	"gym-management-memberships/src/lib/persistence/psql/gorm/models"
	"gym-management-memberships/src/lib/primitives/application_specific"
)

type GetBillsQueryHandler struct{}

func (h *GetBillsQueryHandler) Handle(query *GetBillsQuery) (*GetBillsQueryResponse, *application_specific.ApplicationException) {
	db := lib.GormDB(query.Session)

	options := application_specific.NewPaginationOptions(&query.PaginatedQuery)

	var bills []models.Bill

	dbQuery := db.Model(&models.Bill{}).
		Joins("Membership").
		Joins("Membership.Plan").
		Joins("Membership.Customer")

	if len(query.Id) > 0 {
		dbQuery.Where("id IN (?)", query.Id)
	}

	if len(query.MembershipId) > 0 {
		dbQuery.Where("membership_id IN (?)", query.MembershipId)
	}

	if len(query.CustomerId) > 0 {
		dbQuery.Where("memberships.customer_id IN (?)", query.CustomerId)
	}

	if len(query.PlanId) > 0 {
		dbQuery.Where("memberships.plan_id IN (?)", query.PlanId)
	}

	if len(query.GymId) > 0 {
		dbQuery.Where("gym_id IN (?)", query.GymId)
	}

	if query.Paid != nil {
		dbQuery.Where("paid = ?", *query.Paid)
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
