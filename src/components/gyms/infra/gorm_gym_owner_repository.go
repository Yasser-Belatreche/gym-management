package infra

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gym-management/src/components/gyms/core/domain"
	"gym-management/src/lib"
	"gym-management/src/lib/persistence/psql/gorm/models"
	"gym-management/src/lib/primitives/application_specific"
)

type GormGymOwnerRepository struct {
}

func (g *GormGymOwnerRepository) FindByID(id string, session *application_specific.Session) (*domain.GymOwner, *application_specific.ApplicationException) {
	db := lib.GormDB(session)

	var owner models.GymOwner
	result := db.Clauses(clause.Locking{Strength: "UPDATE"}).Preload("Gyms").First(&owner, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, application_specific.NewNotFoundException("GYMS.OWNERS.NOT_FOUND", "Gym owner not found", map[string]string{
				"id": id,
			})
		}
		return nil, application_specific.NewUnknownException("GYMS.INFRA.FAILED_TO_FIND_GYM_OWNER", result.Error.Error(), nil)
	}

	domainOwner, err := toDomain(&owner)
	if err != nil {
		return nil, err
	}

	return domainOwner, nil
}

func (g *GormGymOwnerRepository) Create(owner *domain.GymOwner, session *application_specific.Session) *application_specific.ApplicationException {
	db := lib.GormDB(session)

	ownerModel, gyms := toDB(owner)

	result := db.Create(ownerModel)
	if result.Error != nil {
		return application_specific.NewUnknownException("GYMS.INFRA.FAILED_TO_CREATE_GYM_OWNER", result.Error.Error(), nil)
	}

	if len(gyms) == 0 {
		return nil
	}
	result = db.Create(gyms)
	if result.Error != nil {
		return application_specific.NewUnknownException("GYMS.INFRA.FAILED_TO_CREATE_GYM", result.Error.Error(), nil)
	}

	return nil
}

func (g *GormGymOwnerRepository) Update(owner *domain.GymOwner, session *application_specific.Session) *application_specific.ApplicationException {
	db := lib.GormDB(session)

	ownerModel, gyms := toDB(owner)

	res := db.Save(ownerModel)
	if res.Error != nil {
		return application_specific.NewUnknownException("GYMS.INFRA.FAILED_TO_UPDATE_GYM_OWNER", res.Error.Error(), nil)
	}

	if len(gyms) == 0 {
		return nil
	}
	res = db.Save(gyms)
	if res.Error != nil {
		return application_specific.NewUnknownException("GYMS.INFRA.FAILED_TO_UPDATE_GYM", res.Error.Error(), nil)
	}

	return nil
}

func toDomain(owner *models.GymOwner) (*domain.GymOwner, *application_specific.ApplicationException) {
	domainGyms := make([]domain.GymState, len(owner.Gyms))

	for i, gym := range owner.Gyms {
		domainGyms[i] = domain.GymState{
			Id:          gym.Id,
			Name:        gym.Name,
			Address:     gym.Address,
			Enabled:     gym.Enabled,
			DisabledFor: gym.DisabledFor,
			CreatedBy:   gym.CreatedBy,
			UpdatedBy:   gym.UpdatedBy,
			DeletedBy:   gym.DeletedBy,
			DeletedAt:   gym.DeletedAt,
		}
	}

	ownerState := domain.GymOwnerState{
		Id:          owner.Id,
		Name:        owner.Name,
		PhoneNumber: owner.PhoneNumber,
		Email:       owner.Email,
		Restricted:  owner.Restricted,
		CreatedBy:   owner.CreatedBy,
		UpdatedBy:   owner.UpdatedBy,
		DeletedBy:   owner.DeletedBy,
		DeletedAt:   owner.DeletedAt,
		Gyms:        domainGyms,
	}

	return domain.GymOwnerFromState(ownerState), nil
}

func toDB(owner *domain.GymOwner) (*models.GymOwner, []*models.Gym) {
	state := owner.State()
	gyms := make([]*models.Gym, len(owner.State().Gyms))

	for i, gym := range owner.State().Gyms {
		gyms[i] = &models.Gym{
			Id:          gym.Id,
			Name:        gym.Name,
			Address:     gym.Address,
			Enabled:     gym.Enabled,
			DisabledFor: gym.DisabledFor,
			OwnerId:     state.Id,
			CreatedBy:   gym.CreatedBy,
			UpdatedBy:   gym.UpdatedBy,
			DeletedBy:   gym.DeletedBy,
			DeletedAt:   gym.DeletedAt,
		}
	}

	ownerModel := &models.GymOwner{
		Id:          state.Id,
		Name:        state.Name,
		PhoneNumber: state.PhoneNumber,
		Email:       state.Email,
		Restricted:  state.Restricted,
		CreatedBy:   state.CreatedBy,
		UpdatedBy:   state.UpdatedBy,
		DeletedBy:   state.DeletedBy,
		DeletedAt:   state.DeletedAt,
	}

	return ownerModel, gyms
}
