package main

import (
	"gym-management-auth/src/components"
	"gym-management-auth/src/web/gin"
)

func main() {
	components.Initialize()
	gin.StartWebServer()
}
