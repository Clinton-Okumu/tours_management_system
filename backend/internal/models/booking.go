package models

import "time"

type Booking struct {
	BaseModel
	UserID    uint      `json:"user_id" gorm:"not null"`
	TourID    uint      `json:"tour_id" gorm:"not null"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}
