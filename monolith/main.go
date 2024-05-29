package main

import (
	"gym-management/src/components"
	"gym-management/src/web/gin"
)

func main() {
	components.Initialize()
	gin.StartWebServer()
}
