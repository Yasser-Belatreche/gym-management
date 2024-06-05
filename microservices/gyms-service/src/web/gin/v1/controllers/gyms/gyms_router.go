package gyms

import (
	"github.com/gin-gonic/gin"
	"gym-management-gyms/src/web/gin/v1/middlewares"
)

func GymsRouter(router *gin.RouterGroup) {
	g := router.Group("/owners")

	g.Use(middlewares.AuthMiddleware())
	{
		g.GET("/:ownerId", GetGymOwnerHandler)
		g.GET("/", GetGymOwnersHandler)
		g.POST("/", middlewares.TransactionMiddleware(), CreateGymOwnerHandler)
		g.PUT("/:ownerId", middlewares.TransactionMiddleware(), UpdateGymOwnerHandler)
		g.DELETE("/:ownerId", middlewares.TransactionMiddleware(), DeleteGymOwnerHandler)
		g.PATCH("/:ownerId/restrict", middlewares.TransactionMiddleware(), RestrictGymOwnerHandler)
		g.PATCH("/:ownerId/unrestrict", middlewares.TransactionMiddleware(), UnrestrictGymOwnerHandler)
	}

	gyms := g.Group("/:ownerId/gyms")
	{
		gyms.GET("/:gymId", GetGymHandler)
		gyms.GET("/", GetGymsHandler)
		gyms.POST("/", middlewares.TransactionMiddleware(), CreateGymHandler)
		gyms.PUT("/:gymId", middlewares.TransactionMiddleware(), UpdateGymHandler)
		gyms.DELETE("/:gymId", middlewares.TransactionMiddleware(), DeleteGymHandler)
		gyms.PATCH("/:gymId/enable", middlewares.TransactionMiddleware(), EnableGymHandler)
		gyms.PATCH("/:gymId/disable", middlewares.TransactionMiddleware(), DisableGymHandler)
	}
}
