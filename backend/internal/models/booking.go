package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"not null"`
	TourID    uint `gorm:"not null"`
	StartDate time.Time
	EndDate   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
