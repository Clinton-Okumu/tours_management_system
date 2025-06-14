package models

import (
	"time"
)

type Tour struct {
	ID           uint    `gorm:"primaryKey"`
	Name         string  `gorm:"type:varchar(100);not null"`
	Duration     int     `gorm:"not null"` // e.g., in days
	MaxGroupSize int     `gorm:"not null"`
	Difficulty   string  `gorm:"type:varchar(50);not null"`
	Summary      string  `gorm:"type:text"`
	Description  string  `gorm:"type:text"`
	Price        float64 `gorm:"not null"`
	ImageCover   string  `gorm:"type:varchar(255)"`

	StartDates []time.Time `gorm:"-"` // optionally handled in a separate table if needed

	CreatedAt time.Time
	UpdatedAt time.Time
}
