package get_gyms

import (
	gyms2 "gym-management/src/components/gyms/core/usecases/gyms"
	"gym-management/src/lib"
	"gym-management/src/lib/persistence/psql/gorm/models"
	"gym-management/src/lib/primitives/application_specific"
)

type GetGymsQueryHandler struct {
}

func (h *GetGymsQueryHandler) Handle(query *GetGymsQuery) (*GetGymsQueryResponse, *application_specific.ApplicationException) {
	db := lib.GormDB(query.Session)

	options := application_specific.NewPaginationOptions(&query.PaginatedQuery)

	var gyms []models.Gym

	dbQuery := db.Model(&models.Gym{})

	if len(query.Id) > 0 {
		dbQuery = dbQuery.Where("id IN (?)", query.Id)
	}

	if len(query.OwnerId) > 0 {
		dbQuery = dbQuery.Where("owner_id IN (?)", query.OwnerId)
	}

	if query.Search != "" {
		dbQuery = dbQuery.Where("name ILIKE ?", "%"+query.Search+"%")
	}

	if query.Enabled != nil {
		dbQuery = dbQuery.Where("enabled = ?", *query.Enabled)
	}

	if query.Deleted {
		dbQuery = dbQuery.Where("deleted_at IS NOT NULL")
	} else {
		dbQuery = dbQuery.Where("deleted_at IS NULL")
	}

	result := dbQuery.Offset(options.Skip).Limit(options.PerPage).Find(&gyms)
	if result.Error != nil {
		return nil, application_specific.NewUnknownException("GYMS.FAILED_TO_GET_GYMS", result.Error.Error(), nil)
	}

	var total int64
	result = dbQuery.Count(&total)
	if result.Error != nil {
		return nil, application_specific.NewUnknownException("GYMS.FAILED_TO_GET_GYMS", result.Error.Error(), nil)
	}

	response := GetGymsQueryResponse(application_specific.NewPaginatedResponse(options, total, gyms, func(item models.Gym) gyms2.GymToReturn {
		return gyms2.GymToReturn{
			Id:          item.Id,
			Name:        item.Name,
			Address:     item.Address,
			Enabled:     item.Enabled,
			OwnerId:     item.OwnerId,
			DisabledFor: item.DisabledFor,
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
