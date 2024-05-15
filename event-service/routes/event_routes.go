package routes

import (
	"event-service/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.RouterGroup) {
	router.GET("/events", controllers.GetAllEvents)
	router.POST("/events", controllers.CreateEvent)
	router.GET("/events/:id", controllers.GetEvent)
	router.PATCH("/events/:id", controllers.UpdateEvent)
	router.DELETE("/events/:id", controllers.DeleteEvent)
}
