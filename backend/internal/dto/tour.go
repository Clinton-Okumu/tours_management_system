package dto

type TourCreateRequest struct {
	Name         string  `json:"name" example:"Safari Tour"`
	Duration     int     `json:"duration" example:"7"`
	MaxGroupSize int     `json:"maxGroupSize" example:"10"`
	Difficulty   string  `json:"difficulty" example:"easy"`
	Summary      string  `json:"summary" example:"A short summary of the tour"`
	Description  string  `json:"description" example:"Detailed description of the tour"`
	Price        float64 `json:"price" example:"299.99"`
	ImageCover   string  `json:"imageCover" example:"cover.jpg"`
}

type TourResponse struct {
	ID           uint    `json:"id" example:"1"`
	Name         string  `json:"name"`
	Duration     int     `json:"duration"`
	MaxGroupSize int     `json:"maxGroupSize"`
	Difficulty   string  `json:"difficulty"`
	Summary      string  `json:"summary"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	ImageCover   string  `json:"imageCover"`
	CreatedAt    string  `json:"created_at" example:"2025-07-11T14:00:00Z"`
	UpdatedAt    string  `json:"updated_at" example:"2025-07-11T14:05:00Z"`
}
