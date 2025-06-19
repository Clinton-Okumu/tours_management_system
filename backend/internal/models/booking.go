package models

import (
	"time"
)

type Booking struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"not null"`
	TourID    uint `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
