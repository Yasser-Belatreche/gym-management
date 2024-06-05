package main

import (
	"gym-management-memberships/src/components"
	"gym-management-memberships/src/web/gin"
)

func main() {
	components.Initialize()
	gin.StartWebServer()
}
