package auth

import (
	"github.com/gin-gonic/gin"
	"gym-management/src/components"
	"gym-management/src/components/auth/core/usecases/get_user"
	"gym-management/src/components/auth/core/usecases/login"
	"gym-management/src/web/gin/v1/controllers/auth/contracts"
	"gym-management/src/web/gin/v1/utils"
	"net/http"
)

func LoginHandler(c *gin.Context) {
	var request contracts.LoginRequest

	if err := c.ShouldBind(&request); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	res, err := components.Auth().Login(&login.LoginCommand{
		Username: request.Username,
		Password: request.Password,
		Session:  utils.ExtractSession(c),
	})

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, contracts.LoginResponse{Token: res.Token})
}

func GetCurrentUser(c *gin.Context) {
	session := utils.ExtractUserSession(c)

	res, err := components.Auth().GetUser(&get_user.GetUserQuery{
		Id:      session.UserId(),
		Session: session,
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, contracts.GetCurrentUserResponse{
		Id:        res.Id,
		Role:      res.Role,
		Usernames: res.Usernames,
		FirstName: res.FirstName,
		LastName:  res.LastName,
		Email:     res.Email,
		Phone:     res.Phone,
		LastLogin: res.LastLogin,
	})
}
