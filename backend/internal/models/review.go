package models

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	ID        uint    `gorm:"primaryKey"`
	UserID    uint    `gorm:"not null"`
	TourID    uint    `gorm:"not null"`
	Rating    float64 `gorm:"not null;check:rating >= 1 AND rating <= 5"`
	Review    string  `gorm:"type:text;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
