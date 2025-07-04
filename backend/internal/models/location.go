package models

type Location struct {
	BaseModel
	TourID  uint    `json:"tour_id" gorm:"not null"`
	Name    string  `json:"name" gorm:"type:varchar(100);not null"`
	Address string  `json:"address" gorm:"type:text"`
	Lat     float64 `json:"lat" gorm:"not null"`
	Lng     float64 `json:"lng" gorm:"not null"`
}
