package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	Venue       string
	Url         string
	Dates       []time.Time `gorm:"type:timestamptz[]"`
	StartDate   time.Time
	EndDate     time.Time
	Deadline    time.Time
}
