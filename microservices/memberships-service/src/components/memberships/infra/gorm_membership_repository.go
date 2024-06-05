package infra

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gym-management/src/components/memberships/core/domain"
	"gym-management/src/lib"
	"gym-management/src/lib/persistence/psql/gorm/models"
	"gym-management/src/lib/primitives/application_specific"
	"time"
)

type GormMembershipRepository struct{}

func (g *GormMembershipRepository) FindByID(id string, session *application_specific.Session) (*domain.Membership, *application_specific.ApplicationException) {
	db := lib.GormDB(session)

	var membership models.Membership
	result := db.Clauses(clause.Locking{Strength: "UPDATE"}).First(&membership, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, application_specific.NewNotFoundException("MEMBERSHIPS.NOT_FOUND", "Membership not found", map[string]string{
				"id": id,
			})
		}

		return nil, application_specific.NewUnknownException("MEMBERSHIPS.INFRA.FAILED_TO_FIND_MEMBERSHIP", result.Error.Error(), map[string]string{})
	}

	var t models.TrainingSession
	var currentTrainingSession = &t
	result = db.First(&t, "membership_id = ? AND ended_at IS NULL", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			currentTrainingSession = nil
		} else {
			return nil, application_specific.NewUnknownException("MEMBERSHIPS.INFRA.FAILED_TO_FIND_MEMBERSHIP", result.Error.Error(), map[string]string{})
		}
	}

	var unpaidBills []models.Bill
	result = db.Find(&unpaidBills, "membership_id = ? AND paid = false", id)
	if result.Error != nil {
		return nil, application_specific.NewUnknownException("MEMBERSHIPS.INFRA.FAILED_TO_FIND_MEMBERSHIP", result.Error.Error(), map[string]string{})
	}

	return membershipToDomain(&membership, currentTrainingSession, unpaidBills), nil
}

func (g *GormMembershipRepository) FindLatestCustomerMembership(customerId string, session *application_specific.Session) (*domain.Membership, *application_specific.ApplicationException) {
	db := lib.GormDB(session)

	var membership models.Membership
	result := db.Clauses(clause.Locking{Strength: "UPDATE"}).Order("created_at DESC").First(&membership, "customer_id = ?", customerId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, application_specific.NewNotFoundException("MEMBERSHIPS.NOT_FOUND", "Membership not found", map[string]string{
				"customerId": customerId,
			})
		}

		return nil, application_specific.NewUnknownException("MEMBERSHIPS.INFRA.FAILED_TO_FIND_MEMBERSHIP", result.Error.Error(), map[string]string{})
	}

	var t models.TrainingSession
	var currentTrainingSession = &t
	result = db.First(&t, "membership_id = ? AND ended_at IS NULL", membership.Id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			currentTrainingSession = nil
		} else {
			return nil, application_specific.NewUnknownException("MEMBERSHIPS.INFRA.FAILED_TO_FIND_MEMBERSHIP", result.Error.Error(), map[string]string{})
		}
	}

	var unpaidBills []models.Bill
	result = db.Find(&unpaidBills, "membership_id = ? AND paid = false", membership.Id)
	if result.Error != nil {
		return nil, application_specific.NewUnknownException("MEMBERSHIPS.INFRA.FAILED_TO_FIND_MEMBERSHIP", result.Error.Error(), map[string]string{})
	}

	return membershipToDomain(&membership, currentTrainingSession, unpaidBills), nil
}

func (g *GormMembershipRepository) Create(membership *domain.Membership, session *application_specific.Session) *application_specific.ApplicationException {
	db := lib.GormDB(session)

	membershipModel, currentTrainingSessionModel, billsModels := membershipToDB(membership)

	result := db.Create(membershipModel)
	if result.Error != nil {
		return application_specific.NewUnknownException("MEMBERSHIPS.INFRA.FAILED_TO_CREATE_MEMBERSHIP", result.Error.Error(), map[string]string{})
	}

	if currentTrainingSessionModel != nil {
		result = db.Create(currentTrainingSessionModel)
		if result.Error != nil {
			return application_specific.NewUnknownException("MEMBERSHIPS.INFRA.FAILED_TO_CREATE_MEMBERSHIP", result.Error.Error(), map[string]string{})
		}

	}

	if billsModels != nil && len(billsModels) > 0 {
		result = db.Create(billsModels)
		if result.Error != nil {
			return application_specific.NewUnknownException("MEMBERSHIPS.INFRA.FAILED_TO_CREATE_MEMBERSHIP", result.Error.Error(), map[string]string{})
		}
	}

	return nil
}

func (g *GormMembershipRepository) Update(membership *domain.Membership, session *application_specific.Session) *application_specific.ApplicationException {
	db := lib.GormDB(session)

	membershipModel, currentTrainingSessionModel, billsModels := membershipToDB(membership)

	result := db.Save(membershipModel)
	if result.Error != nil {
		return application_specific.NewUnknownException("MEMBERSHIPS.INFRA.FAILED_TO_UPDATE_MEMBERSHIP", result.Error.Error(), map[string]string{})
	}

	if currentTrainingSessionModel != nil {
		result = db.Save(currentTrainingSessionModel)
		if result.Error != nil {
			return application_specific.NewUnknownException("MEMBERSHIPS.INFRA.FAILED_TO_UPDATE_MEMBERSHIP", result.Error.Error(), map[string]string{})
		}
	}

	if billsModels != nil && len(billsModels) > 0 {
		result = db.Save(billsModels)
		if result.Error != nil {
			return application_specific.NewUnknownException("MEMBERSHIPS.INFRA.FAILED_TO_UPDATE_MEMBERSHIP", result.Error.Error(), map[string]string{})
		}
	}

	return nil
}

func (g *GormMembershipRepository) FindEnabledMemberships(session *application_specific.Session) ([]*domain.Membership, *application_specific.ApplicationException) {
	db := lib.GormDB(session)

	var memberships []models.Membership
	result := db.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&memberships, "enabled = true")
	if result.Error != nil {
		return nil, application_specific.NewUnknownException("MEMBERSHIPS.INFRA.FAILED_TO_FIND_MEMBERSHIPS", result.Error.Error(), map[string]string{})
	}

	var domainMemberships []*domain.Membership
	for _, membership := range memberships {
		var t models.TrainingSession
		var currentTrainingSession = &t
		result = db.First(currentTrainingSession, "membership_id = ? AND ended_at IS NULL", membership.Id)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				currentTrainingSession = nil
			} else {
				return nil, application_specific.NewUnknownException("MEMBERSHIPS.INFRA.FAILED_TO_FIND_MEMBERSHIP", result.Error.Error(), map[string]string{})
			}
		}

		var unpaidBills []models.Bill
		result = db.Find(&unpaidBills, "membership_id = ? AND paid = false", membership.Id)
		if result.Error != nil {
			return nil, application_specific.NewUnknownException("MEMBERSHIPS.INFRA.FAILED_TO_FIND_MEMBERSHIP", result.Error.Error(), map[string]string{})
		}

		domainMemberships = append(domainMemberships, membershipToDomain(&membership, currentTrainingSession, unpaidBills))
	}

	return domainMemberships, nil
}

func (g *GormMembershipRepository) FindEnabledMembershipsByGymID(gymId string, session *application_specific.Session) ([]*domain.Membership, *application_specific.ApplicationException) {
	db := lib.GormDB(session)

	var memberships []models.Membership
	result := db.Clauses(clause.Locking{Strength: "UPDATE"}).Joins("Plan").Find(&memberships, "memberships.enabled = true AND plan.gym_id = ?", gymId)
	if result.Error != nil {
		return nil, application_specific.NewUnknownException("MEMBERSHIPS.INFRA.FAILED_TO_FIND_MEMBERSHIPS", result.Error.Error(), map[string]string{})
	}

	var domainMemberships []*domain.Membership
	for _, membership := range memberships {
		var t models.TrainingSession
		var currentTrainingSession *models.TrainingSession
		result = db.First(&t, "membership_id = ? AND ended_at IS NULL", membership.Id)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				currentTrainingSession = nil
			} else {
				return nil, application_specific.NewUnknownException("MEMBERSHIPS.INFRA.FAILED_TO_FIND_MEMBERSHIP", result.Error.Error(), map[string]string{})
			}
		}
		currentTrainingSession = &t

		var unpaidBills []models.Bill
		result = db.Find(&unpaidBills, "membership_id = ? AND paid = false", membership.Id)
		if result.Error != nil {
			return nil, application_specific.NewUnknownException("MEMBERSHIPS.INFRA.FAILED_TO_FIND_MEMBERSHIP", result.Error.Error(), map[string]string{})
		}

		domainMemberships = append(domainMemberships, membershipToDomain(&membership, currentTrainingSession, unpaidBills))
	}

	return domainMemberships, nil
}

func (g *GormMembershipRepository) CountTrainingSessionsThisWeek(membershipId string, session *application_specific.Session) (int, *application_specific.ApplicationException) {
	db := lib.GormDB(session)

	now := time.Now().UTC()
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	daysSinceSunday := int(now.Weekday()) - int(time.Sunday)

	startOfWeek := now.AddDate(0, 0, -daysSinceSunday)
	endOfWeek := startOfWeek.AddDate(0, 0, 6)

	var count int64
	result := db.Model(&models.TrainingSession{}).Where("membership_id = ? AND started_at >= ? AND started_at <= ?", membershipId, startOfWeek, endOfWeek).Count(&count)
	if result.Error != nil {
		return 0, application_specific.NewUnknownException("MEMBERSHIPS.INFRA.FAILED_TO_COUNT_TRAINING_SESSIONS", result.Error.Error(), map[string]string{})
	}

	return int(count), nil
}

func membershipToDB(membership *domain.Membership) (*models.Membership, *models.TrainingSession, []*models.Bill) {
	state := membership.State()

	membershipModel := &models.Membership{
		Id:              state.Id,
		Code:            state.Code,
		StartDate:       state.StartDate,
		EndDate:         state.EndDate,
		Enabled:         state.Enabled,
		DisabledFor:     state.DisabledFor,
		SessionsPerWeek: state.SessionsPerWeek,
		WithCoach:       state.WithCoach,
		MonthlyPrice:    state.MonthlyPrice,
		PlanId:          state.PlanId,
		CustomerId:      state.CustomerId,
		RenewedAt:       state.RenewedAt,
		CreatedBy:       state.CreatedBy,
		UpdatedBy:       state.UpdatedBy,
	}

	var trainingSessionModel *models.TrainingSession = nil
	if state.CurrentTrainingSession != nil {
		trainingSessionModel = &models.TrainingSession{
			Id:           state.CurrentTrainingSession.Id,
			MembershipId: state.Id,
			StartedAt:    state.CurrentTrainingSession.StartedAt,
			EndedAt:      state.CurrentTrainingSession.EndedAt,
		}
	}

	billsModels := make([]*models.Bill, len(state.UnpaidBills))
	for i, bill := range state.UnpaidBills {
		billsModels[i] = &models.Bill{
			Id:           bill.Id,
			Amount:       bill.Amount,
			Paid:         bill.Paid,
			PaidAt:       bill.PaidAt,
			DueTo:        bill.DueDate,
			MembershipId: state.Id,
		}
	}

	return membershipModel, trainingSessionModel, billsModels
}

func membershipToDomain(membership *models.Membership, currentTrainingSession *models.TrainingSession, unpaidBills []models.Bill) *domain.Membership {
	var currentTrainingSessionState *domain.TrainingSessionState = nil
	if currentTrainingSession != nil {
		currentTrainingSessionState = &domain.TrainingSessionState{
			Id:        currentTrainingSession.Id,
			StartedAt: currentTrainingSession.StartedAt,
			EndedAt:   currentTrainingSession.EndedAt,
		}
	}

	unpaidBillsState := make([]*domain.BillState, len(unpaidBills))
	for i, bill := range unpaidBills {
		unpaidBillsState[i] = &domain.BillState{
			Id:        bill.Id,
			Amount:    bill.Amount,
			Paid:      bill.Paid,
			PaidAt:    bill.PaidAt,
			DueDate:   bill.DueTo,
			CreatedAt: bill.CreatedAt,
		}
	}

	state := &domain.MembershipState{
		Id:                     membership.Id,
		Code:                   membership.Code,
		StartDate:              membership.StartDate,
		EndDate:                membership.EndDate,
		Enabled:                membership.Enabled,
		DisabledFor:            membership.DisabledFor,
		SessionsPerWeek:        membership.SessionsPerWeek,
		WithCoach:              membership.WithCoach,
		MonthlyPrice:           membership.MonthlyPrice,
		UnpaidBills:            unpaidBillsState,
		CurrentTrainingSession: currentTrainingSessionState,
		PlanId:                 membership.PlanId,
		RenewedAt:              membership.RenewedAt,
		CustomerId:             membership.CustomerId,
		CreatedBy:              membership.CreatedBy,
		UpdatedBy:              membership.UpdatedBy,
	}

	return domain.MembershipFromState(state)
}
