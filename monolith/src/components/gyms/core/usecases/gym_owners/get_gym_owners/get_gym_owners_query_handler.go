package get_gym_owners

import (
	"gym-management/src/components/gyms/core/usecases/gym_owners"
	"gym-management/src/lib"
	"gym-management/src/lib/persistence/psql/gorm/models"
	"gym-management/src/lib/primitives/application_specific"
)

type GetGymOwnersQueryHandler struct{}

func (h *GetGymOwnersQueryHandler) Handle(query *GetGymOwnersQuery) (*GetGymOwnersQueryResponse, *application_specific.ApplicationException) {
	db := lib.GormDB(query.Session)

	options := application_specific.NewPaginationOptions(&query.PaginatedQuery)

	var gymOwners []models.GymOwner

	dbQuery := db.Model(&models.GymOwner{})

	if query.Id != nil {
		dbQuery.Where("id IN ?", query.Id)
	}

	if query.Search != "" {
		dbQuery.Where("name ILIKE ?", "%"+query.Search+"%")
	}

	if query.Restricted != nil {
		dbQuery.Where("restricted = ?", *query.Restricted)
	}

	if query.Deleted {
		dbQuery.Where("deleted_at IS NOT NULL")
	} else {
		dbQuery.Where("deleted_at IS NULL")
	}

	result := dbQuery.Offset(options.Skip).Limit(options.PerPage).Find(&gymOwners)

	if result.Error != nil {
		return nil, application_specific.NewUnknownException("FAILED_TO_GET_GYM_OWNERS", result.Error.Error(), nil)
	}

	var total int64
	countResult := dbQuery.Count(&total)
	if countResult.Error != nil {
		return nil, application_specific.NewUnknownException("FAILED_TO_GET_GYM_OWNERS", countResult.Error.Error(), nil)
	}

	response := GetGymOwnersQueryResponse(application_specific.NewPaginatedResponse(options, total, gymOwners, func(item models.GymOwner) gym_owners.GymOwnerToReturn {
		return gym_owners.GymOwnerToReturn{
			Id:          item.Id,
			Name:        item.Name,
			PhoneNumber: item.PhoneNumber,
			Email:       item.Email,
			Restricted:  item.Restricted,
			CreatedBy:   item.CreatedBy,
			CreatedAt:   item.CreatedAt,
			UpdatedBy:   item.UpdatedBy,
			UpdatedAt:   item.UpdatedAt,
			DeletedBy:   item.DeletedBy,
			DeletedAt:   item.DeletedAt,
		}
	}))

	return &response, nil
}
