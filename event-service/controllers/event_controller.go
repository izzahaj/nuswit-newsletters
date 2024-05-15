package controllers

import (
	"event-service/initializers"
	"event-service/models"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllEvents(c *gin.Context) {
	var events []models.Event
	result := initializers.DB.Find(&events)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve events"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"events": events})
}

func GetEvent(c *gin.Context) {
	id := c.Param("id")

	var event models.Event
	result := initializers.DB.First(&event, id)

	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"event": event})
}

func CreateEvent(c *gin.Context) {
	var body struct {
		Name        string
		Description string
		Venue       string
		Url         string
		Dates       []time.Time
		StartDate   time.Time
		EndDate     time.Time
		Deadline    time.Time
	}

	// TODO: implement authentication and authorisation checks (using middleware)

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid fields provided"})
		return
	}

	// TODO: fix bug where a request body with only a blank description gives unexpected error message
	if len(strings.TrimSpace(body.Name)) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name cannot be blank"})
		return
	}

	if len(strings.TrimSpace(body.Description)) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "description cannot be blank"})
		return
	}

	var existingEvent models.Event

	r := initializers.DB.Where(&models.Event{
		Name:        body.Name,
		Description: body.Description,
		Venue:       body.Venue,
		Url:         body.Url,
		StartDate:   body.StartDate,
		EndDate:     body.EndDate,
		Deadline:    body.Deadline,
	}).First(&existingEvent)

	if r.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "event with the same details already exists"})
		return
	}

	if r.Error != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create event"})
		return
	}

	event := models.Event{
		Name:        body.Name,
		Description: body.Description,
		Venue:       body.Venue,
		Url:         body.Url,
		Dates:       body.Dates,
		StartDate:   body.StartDate,
		EndDate:     body.EndDate,
		Deadline:    body.Deadline,
	}

	result := initializers.DB.Create(&event)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create event"})
		return
	}

	c.JSON(http.StatusCreated, event)
	return
}

func UpdateEvent(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Name        string
		Description string
		Venue       string
		Url         string
		Dates       []time.Time
		StartDate   time.Time
		EndDate     time.Time
		Deadline    time.Time
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid fields provided"})
		return
	}

	if len(strings.TrimSpace(body.Name)) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name cannot be blank"})
		return
	}

	if len(strings.TrimSpace(body.Description)) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "description cannot be blank"})
		return
	}

	var event models.Event
	r := initializers.DB.First(&event, id)

	if r.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}

	if r.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update event"})
		return
	}

	result := initializers.DB.Model(&event).Updates(models.Event{
		Name:        body.Name,
		Description: body.Description,
		Venue:       body.Venue,
		Url:         body.Url,
		Dates:       body.Dates,
		StartDate:   body.StartDate,
		EndDate:     body.EndDate,
		Deadline:    body.Deadline,
	})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"event": event})
}

func DeleteEvent(c *gin.Context) {
	id := c.Param("id")

	result := initializers.DB.Delete(&models.Event{}, id)

	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete event"})
		return
	}

	c.Status(http.StatusNoContent)
}
