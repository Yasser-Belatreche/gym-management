package gyms

import (
	"github.com/gin-gonic/gin"
	"gym-management/src/web/gin/v1/middlewares"
)

func GymsRouter(router *gin.RouterGroup) {
	g := router.Group("/owners")

	g.Use(middlewares.AuthMiddleware())

	g.GET("/:ownerId", GetGymOwnerHandler)
	g.GET("/", GetGymOwnersHandler)
	g.POST("/", middlewares.TransactionMiddleware(), CreateGymOwnerHandler)
	g.PUT("/:ownerId", middlewares.TransactionMiddleware(), UpdateGymOwnerHandler)
	g.DELETE("/:ownerId", middlewares.TransactionMiddleware(), DeleteGymOwnerHandler)
	g.PATCH("/:ownerId/restrict", RestrictGymOwnerHandler)
	g.PATCH("/:ownerId/unrestrict", UnrestrictGymOwnerHandler)

	gg := g.Group("/:ownerId/gyms")

	gg.GET("/:gymId", GetGymHandler)
	gg.GET("/", GetGymsHandler)
	gg.POST("/", CreateGymHandler)
	gg.PUT("/:gymId", UpdateGymHandler)
	gg.DELETE("/:gymId", DeleteGymHandler)
	gg.PATCH("/:gymId/enable", EnableGymHandler)
	gg.PATCH("/:gymId/disable", DisableGymHandler)
}
