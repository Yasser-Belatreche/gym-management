package get_gym

import (
	"errors"
	"gorm.io/gorm"
	"gym-management-gyms/src/components/gyms/core/usecases/gyms"
	"gym-management-gyms/src/lib"
	"gym-management-gyms/src/lib/persistence/psql/gorm/models"
	"gym-management-gyms/src/lib/primitives/application_specific"
)

type GetGymQueryHandler struct {
}

func (h GetGymQueryHandler) Handle(query *GetGymQuery) (*GetGymQueryResponse, *application_specific.ApplicationException) {
	db := lib.GormDB(query.Session)

	var gym models.Gym
	if err := db.Where("id = ?", query.Id).First(&gym).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, application_specific.NewNotFoundException("GYMS.NOT_FOUND", "gym not found", map[string]string{
				"id": query.Id,
			})
		}

		return nil, application_specific.NewUnknownException("GYMS.FAILED_TO_GET_GYM", err.Error(), nil)
	}

	response := GetGymQueryResponse(
		gyms.GymToReturn{
			Id:          gym.Id,
			Name:        gym.Name,
			Address:     gym.Address,
			Enabled:     gym.Enabled,
			OwnerId:     gym.OwnerId,
			DisabledFor: gym.DisabledFor,
			CreatedBy:   gym.CreatedBy,
			CreatedAt:   gym.CreatedAt,
			UpdatedBy:   gym.UpdatedBy,
			UpdatedAt:   gym.UpdatedAt,
			DeletedBy:   gym.DeletedBy,
			DeletedAt:   gym.DeletedAt,
		},
	)

	return &response, nil
}
