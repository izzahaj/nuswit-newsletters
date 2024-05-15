package main

import (
	"event-service/initializers"
	"event-service/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	fmt.Println("Hello")

	router := gin.Default()
	routerGroup := router.Group("/api/event-service")
	routes.SetUpRoutes(routerGroup)

	router.Run(":5000")
}
