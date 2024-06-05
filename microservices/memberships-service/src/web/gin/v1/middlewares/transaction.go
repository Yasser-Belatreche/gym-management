package middlewares

import (
	"github.com/gin-gonic/gin"
	"gym-management-memberships/src/lib"
	"gym-management-memberships/src/lib/primitives/application_specific"
	"gym-management-memberships/src/web/gin/v1/utils"
)

func TransactionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := utils.ExtractSession(c)

		err := lib.Persistence().WithTransaction(session, func() *application_specific.ApplicationException {
			c.Next()

			if c.IsAborted() {
				return application_specific.NewUnknownException("REQUEST_ABORTED", "request aborted", nil)
			}
			return nil
		})

		if err != nil {
			utils.HandleError(c, err)
			return
		}
	}
}
