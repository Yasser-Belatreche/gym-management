package middlewares

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gym-management-gyms/src/lib/primitives/application_specific"
	"gym-management-gyms/src/web/gin/v1/utils"
)

func SessionExtractorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionBase64 := c.GetHeader("X-Session")
		if sessionBase64 == "" {
			utils.HandleError(c, utils.NewNoSessionError())
			return
		}

		sessionJsonString, err := base64.StdEncoding.DecodeString(sessionBase64)
		if err != nil {
			utils.HandleError(c, utils.NewInvalidSessionError())
			return
		}

		var userSession application_specific.UserSession
		err = json.Unmarshal(sessionJsonString, &userSession)
		if err == nil && userSession.User != nil {
			c.Set("session", &userSession)
			c.Next()
			return
		}

		var session application_specific.Session
		err = json.Unmarshal(sessionJsonString, &session)
		if err == nil {
			c.Set("session", &session)
			c.Next()
			return
		}

		utils.HandleError(c, utils.NewInvalidSessionError())
	}
}
