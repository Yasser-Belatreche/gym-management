package gyms

import (
	"github.com/gin-gonic/gin"
	"gym-management/src/components"
	"gym-management/src/components/gyms/core/usecases/gym_owners"
	"gym-management/src/components/gyms/core/usecases/gym_owners/create_gym_owner"
	"gym-management/src/components/gyms/core/usecases/gym_owners/delete_gym_owner"
	"gym-management/src/components/gyms/core/usecases/gym_owners/get_gym_owner"
	"gym-management/src/components/gyms/core/usecases/gym_owners/get_gym_owners"
	"gym-management/src/components/gyms/core/usecases/gym_owners/restrict_gym_owner"
	"gym-management/src/components/gyms/core/usecases/gym_owners/unrestrict_gym_owner"
	"gym-management/src/components/gyms/core/usecases/gym_owners/update_gym_owner"
	gyms2 "gym-management/src/components/gyms/core/usecases/gyms"
	"gym-management/src/components/gyms/core/usecases/gyms/create_gym"
	"gym-management/src/components/gyms/core/usecases/gyms/delete_gym"
	"gym-management/src/components/gyms/core/usecases/gyms/disable_gym"
	"gym-management/src/components/gyms/core/usecases/gyms/enable_gym"
	"gym-management/src/components/gyms/core/usecases/gyms/get_gym"
	"gym-management/src/components/gyms/core/usecases/gyms/get_gyms"
	"gym-management/src/components/gyms/core/usecases/gyms/update_gym"
	"gym-management/src/lib/primitives/application_specific"
	"gym-management/src/web/gin/v1/controllers/gyms/contracts"
	"gym-management/src/web/gin/v1/controllers/gyms/contracts/base"
	"gym-management/src/web/gin/v1/utils"
	"net/http"
)

func GetGymOwnerHandler(c *gin.Context) {
	var request contracts.GetGymOwnerUrl
	if err := c.ShouldBindUri(&request); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	response, err := components.Gyms().GetGymOwner(&get_gym_owner.GetGymOwnerQuery{
		Id:      request.Id,
		Session: utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, contracts.GetGymOwnerResponse(
		base.GymOwner{
			Id:          response.Id,
			Name:        response.Name,
			PhoneNumber: response.PhoneNumber,
			Email:       response.Email,
			Restricted:  response.Restricted,
			CreatedBy:   response.CreatedBy,
			CreatedAt:   response.CreatedAt,
			UpdatedBy:   response.UpdatedBy,
			UpdatedAt:   response.UpdatedAt,
			DeletedBy:   response.DeletedBy,
			DeletedAt:   response.DeletedAt,
		},
	))
}

func GetGymOwnersHandler(c *gin.Context) {
	var request contracts.GetGymOwnersRequest
	if err := c.ShouldBind(&request); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	response, err := components.Gyms().GetGymOwners(&get_gym_owners.GetGymOwnersQuery{
		PaginatedQuery: application_specific.PaginatedQuery{
			Page:    request.Page,
			PerPage: request.PerPage,
		},
		Id:         request.Id,
		Search:     request.Search,
		Restricted: request.Restricted,
		Deleted:    request.Deleted,
		Session:    utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, utils.GetHttpPaginatedResponse(response, func(item gym_owners.GymOwnerToReturn) base.GymOwner {
		return base.GymOwner{
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
}

func CreateGymOwnerHandler(c *gin.Context) {
	var request contracts.CreateGymOwnerRequest
	if err := c.ShouldBind(&request); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	response, err := components.Gyms().CreateGymOwner(&create_gym_owner.CreateGymOwnerCommand{
		Name:        request.Name,
		PhoneNumber: request.PhoneNumber,
		Email:       request.Email,
		Password:    request.Password,
		Session:     utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, contracts.CreateGymOwnerResponse{Id: response.Id})
}

func UpdateGymOwnerHandler(c *gin.Context) {
	var url contracts.UpdateGymOwnerUrl
	if err := c.ShouldBindUri(&url); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	var request contracts.UpdateGymOwnerRequest
	if err := c.ShouldBind(&request); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	response, err := components.Gyms().UpdateGymOwner(&update_gym_owner.UpdateGymOwnerCommand{
		Id:          url.Id,
		Name:        request.Name,
		PhoneNumber: request.PhoneNumber,
		Email:       request.Email,
		NewPassword: request.NewPassword,
		Session:     utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, contracts.UpdateGymOwnerResponse{Id: response.Id})
}

func DeleteGymOwnerHandler(c *gin.Context) {
	var request contracts.DeleteGymOwnerUrl
	if err := c.ShouldBindUri(&request); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	_, err := components.Gyms().DeleteGymOwner(&delete_gym_owner.DeleteGymOwnerCommand{
		Id:      request.Id,
		Session: utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func RestrictGymOwnerHandler(c *gin.Context) {
	var request contracts.RestrictGymOwnerUrl
	if err := c.ShouldBindUri(&request); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	res, err := components.Gyms().RestrictGymOwner(&restrict_gym_owner.RestrictGymOwnerCommand{
		Id:      request.Id,
		Session: utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, contracts.RestrictGymOwnerResponse{Id: res.Id})
}

func UnrestrictGymOwnerHandler(c *gin.Context) {
	var request contracts.UnrestrictGymOwnerUrl
	if err := c.ShouldBindUri(&request); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	res, err := components.Gyms().UnrestrictGymOwner(&unrestrict_gym_owner.UnrestrictGymOwnerCommand{
		Id:      request.Id,
		Session: utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, contracts.UnrestrictGymOwnerResponse{Id: res.Id})
}

func GetGymHandler(c *gin.Context) {
	var request contracts.GetGymUrl
	if err := c.ShouldBindUri(&request); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	response, err := components.Gyms().GetGym(&get_gym.GetGymQuery{
		Id:      request.Id,
		Session: utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, contracts.GetGymResponse(
		base.Gym{
			Id:          response.Id,
			Name:        response.Name,
			Address:     response.Address,
			Enabled:     response.Enabled,
			OwnerId:     response.OwnerId,
			DisabledFor: response.DisabledFor,
			CreatedBy:   response.CreatedBy,
			CreatedAt:   response.CreatedAt,
			UpdatedBy:   response.UpdatedBy,
			UpdatedAt:   response.UpdatedAt,
			DeletedBy:   response.DeletedBy,
			DeletedAt:   response.DeletedAt,
		},
	))
}

func GetGymsHandler(c *gin.Context) {
	var url contracts.GetGymsUrl
	if err := c.ShouldBindUri(&url); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	var request contracts.GetGymsRequest
	if err := c.ShouldBind(&request); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	response, err := components.Gyms().GetGyms(&get_gyms.GetGymsQuery{
		PaginatedQuery: application_specific.PaginatedQuery{
			Page:    request.Page,
			PerPage: request.PerPage,
		},
		Id:      request.Id,
		OwnerId: []string{url.OwnerId},
		Search:  request.Search,
		Enabled: request.Enabled,
		Deleted: request.Deleted,
		Session: utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, utils.GetHttpPaginatedResponse(response, func(item gyms2.GymToReturn) base.Gym {
		return base.Gym{
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
}

func CreateGymHandler(c *gin.Context) {
	var url contracts.CreateGymUrl
	if err := c.ShouldBindUri(&url); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	var request contracts.CreateGymRequest
	if err := c.ShouldBind(&request); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	response, err := components.Gyms().CreateGym(&create_gym.CreateGymCommand{
		OwnerId: url.OwnerId,
		Name:    request.Name,
		Address: request.Address,
		Session: utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, contracts.CreateGymResponse{Id: response.Id})
}

func UpdateGymHandler(c *gin.Context) {
	var url contracts.UpdateGymUrl
	if err := c.ShouldBindUri(&url); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	var request contracts.UpdateGymRequest
	if err := c.ShouldBind(&request); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	response, err := components.Gyms().UpdateGym(&update_gym.UpdateGymCommand{
		Name:    request.Name,
		Address: request.Address,
		GymId:   url.GymId,
		OwnerId: url.OwnerId,
		Session: utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, contracts.UpdateGymResponse{Id: response.Id})
}

func DeleteGymHandler(c *gin.Context) {
	var url contracts.DeleteGymUrl
	if err := c.ShouldBindUri(&url); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	_, err := components.Gyms().DeleteGym(&delete_gym.DeleteGymCommand{
		GymId:   url.GymId,
		OwnerId: url.OwnerId,
		Session: utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func EnableGymHandler(c *gin.Context) {
	var url contracts.EnableGymUrl
	if err := c.ShouldBindUri(&url); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	res, err := components.Gyms().EnableGym(&enable_gym.EnableGymCommand{
		GymId:   url.GymId,
		OwnerId: url.OwnerId,
		Session: utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, contracts.EnableGymResponse{Id: res.Id})
}

func DisableGymHandler(c *gin.Context) {
	var url contracts.DisableGymUrl
	if err := c.ShouldBindUri(&url); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	res, err := components.Gyms().DisableGym(&disable_gym.DisableGymCommand{
		GymId:   url.GymId,
		OwnerId: url.OwnerId,
		Session: utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, contracts.DisableGymResponse{Id: res.Id})
}
