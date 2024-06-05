package middlewares

import (
	"github.com/gin-gonic/gin"
	"gym-management-memberships/src/web/gin/v1/utils"
	"os"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		exists := utils.CheckUserSession(c)
		if !exists {
			utils.HandleError(c, utils.NewNoUserSessionError())
			return
		}

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
