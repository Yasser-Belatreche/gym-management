package memberships

import (
	"github.com/gin-gonic/gin"
	"gym-management/src/web/gin/v1/middlewares"
)

func MembershipsRouter(r *gin.RouterGroup) {
	g := r.Group("/gym-owners/:ownerId/gyms/:gymId/memberships")

	g.Use(middlewares.AuthMiddleware())
}
