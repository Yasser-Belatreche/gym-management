package middlewares

import (
	"github.com/gin-gonic/gin"
	"gym-management/src/components"
	"gym-management/src/components/auth/core/usecases/get_session"
	"gym-management/src/web/gin/v1/utils"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		chunks := strings.Split(authHeader, " ")
		if len(chunks) != 2 || chunks[0] != "Bearer" {
			utils.HandleError(c, utils.NewNoTokenError())
			return
		}

		token := strings.Split(authHeader, " ")[1]

		session := utils.ExtractSession(c)

		res, err := components.Auth().GetUserSession(&get_session.GetSessionQuery{
			Token:   token,
			Session: session,
		})
		if err != nil {
			utils.HandleError(c, err)
			return
		}

		c.Set("session", res.Session)

		c.Next()
	}
}
