package get_bill

import (
	"errors"
	"gorm.io/gorm"
	bills2 "gym-management/src/components/memberships/core/usecases/bills"
	"gym-management/src/lib"
	"gym-management/src/lib/persistence/psql/gorm/models"
	"gym-management/src/lib/primitives/application_specific"
)

type GetBillQueryHandler struct{}

func (h *GetBillQueryHandler) Handle(query *GetBillQuery) (*GetBillQueryResponse, *application_specific.ApplicationException) {
	db := lib.GormDB(query.Session)

	var bill models.Bill
	dbQuery := db.Model(&models.Bill{})
	dbQuery = dbQuery.Joins("Membership").Select("")
	dbQuery = dbQuery.Joins("Membership.Plan").Select("plans.id, plans.name, plans.gym_id")
	dbQuery = dbQuery.Joins("Membership.Customer").Select("customers.id, customers.first_name, customers.last_name")

	if err := dbQuery.Where("bills.id = ?", query.BillId).First(&bill).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, application_specific.NewNotFoundException("BILLS.NOT_FOUND", "bill not found", map[string]string{
				"id": query.BillId,
			})
		}

		return nil, application_specific.NewUnknownException("BILLS.FAILED_TO_GET_BILL", err.Error(), nil)
	}

	response := GetBillQueryResponse(
		bills2.BillToReturn{
			Id:     bill.Id,
			Amount: bill.Amount,
			Paid:   bill.Paid,
			Customer: bills2.BillToReturnCustomer{
				Id:        bill.Membership.Customer.Id,
				FirstName: bill.Membership.Customer.FirstName,
				LastName:  bill.Membership.Customer.LastName,
			},
			Plan: bills2.BillToReturnPlan{
				Id:   bill.Membership.Plan.Id,
				Name: bill.Membership.Plan.Name,
			},
			MembershipId: bill.MembershipId,
			GymId:        bill.Membership.Plan.GymId,
			PaidAt:       bill.PaidAt,
			DueDate:      bill.DueTo,
			CreatedAt:    bill.CreatedAt,
		})

	return &response, nil

}
