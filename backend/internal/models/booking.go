package models

import (
	"time"
)

type Booking struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"not null"` // foreign key
	TourID    uint `gorm:"not null"` // foreign key
	CreatedAt time.Time
	UpdatedAt time.Time
}
