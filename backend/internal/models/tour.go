package models

type Tour struct {
	BaseModel
	Name         string  `gorm:"type:varchar(100);not null"`
	Duration     int     `gorm:"not null"`
	MaxGroupSize int     `gorm:"not null"`
	Difficulty   string  `gorm:"type:varchar(50);not null"`
	Summary      string  `gorm:"type:text"`
	Description  string  `gorm:"type:text"`
	Price        float64 `gorm:"not null"`
	ImageCover   string  `gorm:"type:varchar(255)"`
}
