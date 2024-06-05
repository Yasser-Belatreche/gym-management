package main

import (
	"gym-management-gyms/src/components"
	"gym-management-gyms/src/web/gin"
)

func main() {
	components.Initialize()
	gin.StartWebServer()
}
