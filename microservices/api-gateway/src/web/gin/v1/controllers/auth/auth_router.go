package auth

import (
	"github.com/gin-gonic/gin"
	"gym-management-api-gateway/src/web/gin/v1/middlewares"
	"gym-management-api-gateway/src/web/gin/v1/utils"
)

func AuthRouter(router *gin.RouterGroup) {
	g := router.Group("/auth")

	g.POST("/login", utils.RedirectToAuthService)

	g.GET("/me", middlewares.AuthMiddleware(), utils.RedirectToAuthService)

	g.GET("/session", middlewares.AuthMiddleware(), utils.RedirectToAuthService)
}
