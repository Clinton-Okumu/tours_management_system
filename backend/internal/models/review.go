package models

import (
	"time"
)

type Review struct {
	ID        uint    `gorm:"primaryKey"`
	UserID    uint    `gorm:"not null"` // foreign key
	TourID    uint    `gorm:"not null"` // foreign key
	Rating    float64 `gorm:"not null;check:rating >= 1 AND rating <= 5"`
	Review    string  `gorm:"type:text;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
