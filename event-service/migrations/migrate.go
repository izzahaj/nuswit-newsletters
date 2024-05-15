package main

import (
	"event-service/initializers"
	"event-service/models"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Event{})
}
