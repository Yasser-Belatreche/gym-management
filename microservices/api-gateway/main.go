package main

import (
	"gym-management-api-gateway/src/components"
	"gym-management-api-gateway/src/web/gin"
)

func main() {
	components.Initialize()
	gin.StartWebServer()
}
