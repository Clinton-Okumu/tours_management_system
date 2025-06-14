package models

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Role      string `gorm:"type:varchar(20);default:'user'"` // e.g. admin, guide, user
	CreatedAt time.Time
	UpdatedAt time.Time
}
