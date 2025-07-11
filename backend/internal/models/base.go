package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `gorm:"primaryKey" json:"id" example:"1"`
	CreatedAt time.Time      `json:"created_at" example:"2025-07-11T14:00:00Z"`
	UpdatedAt time.Time      `json:"updated_at" example:"2025-07-11T14:05:00Z"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-" swaggerignore:"true"`
}
