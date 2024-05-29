package auth

import (
	"github.com/gin-gonic/gin"
	"gym-management/src/web/gin/v1/middlewares"
)

func AuthRouter(router *gin.RouterGroup) {
	g := router.Group("/auth")

	g.POST("/login", middlewares.TransactionMiddleware(), LoginHandler)

	g.GET("/me", middlewares.AuthMiddleware(), GetCurrentUser)
}
