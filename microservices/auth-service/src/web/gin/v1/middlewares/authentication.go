package middlewares

import (
	"github.com/gin-gonic/gin"
	"gym-management-auth/src/components"
	"gym-management-auth/src/components/auth/core/usecases/get_session"
	"gym-management-auth/src/web/gin/v1/utils"
	"os"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		exists := utils.CheckUserSession(c)
		if exists { // means that the user session is passed in the header from the gateway
			c.Next()
			return
		}

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

func ServiceAuthMiddleware() gin.HandlerFunc {
	apiSecret, ok := os.LookupEnv("API_SECRET")
	if !ok {
		panic("API_SECRET env var is required")
	}

	return func(c *gin.Context) {
		secret := c.GetHeader("X-Api-Secret")

		if secret == "" {
			utils.HandleError(c, utils.NewNoApiSecretError())
			return
		}

		if secret != apiSecret {
			utils.HandleError(c, utils.NewWrongApiSecretError())
			return
		}

		c.Next()
	}
}
