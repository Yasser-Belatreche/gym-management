package get_gym_owner

import (
	"errors"
	"gorm.io/gorm"
	"gym-management-gyms/src/components/gyms/core/usecases/gym_owners"
	"gym-management-gyms/src/lib"
	"gym-management-gyms/src/lib/persistence/psql/gorm/models"
	"gym-management-gyms/src/lib/primitives/application_specific"
)

type GetGymOwnerQueryHandler struct{}

func (h *GetGymOwnerQueryHandler) Handle(query *GetGymOwnerQuery) (*GetGymOwnerQueryResponse, *application_specific.ApplicationException) {
	db := lib.GormDB(query.Session)

	var owner models.GymOwner
	results := db.Where("id = ?", query.Id).First(&owner)
	if results.Error != nil {
		if errors.Is(results.Error, gorm.ErrRecordNotFound) {
			return nil, application_specific.NewNotFoundException("GYM.OWNERS.NOT_FOUND", "Gym owner not found", map[string]string{
				"id": query.Id,
			})
		}

		return nil, application_specific.NewUnknownException("GYMS.OWNERS.QUERY_FAILED", results.Error.Error(), nil)
	}

	res := GetGymOwnerQueryResponse(
		gym_owners.GymOwnerToReturn{
			Id:          owner.Id,
			Name:        owner.Name,
			PhoneNumber: owner.PhoneNumber,
			Email:       owner.Email,
			Restricted:  owner.Restricted,
			CreatedBy:   owner.CreatedBy,
			CreatedAt:   owner.CreatedAt,
			UpdatedBy:   owner.UpdatedBy,
			UpdatedAt:   owner.UpdatedAt,
			DeletedBy:   owner.DeletedBy,
			DeletedAt:   owner.DeletedAt,
		},
	)

	return &res, nil

}
