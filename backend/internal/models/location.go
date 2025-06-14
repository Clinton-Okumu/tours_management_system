package models

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	TourID  uint    `gorm:"not null"` // foreign key
	Name    string  `gorm:"type:varchar(100);not null"`
	Address string  `gorm:"type:text"`
	Lat     float64 `gorm:"not null"`
	Lng     float64 `gorm:"not null"`
}
