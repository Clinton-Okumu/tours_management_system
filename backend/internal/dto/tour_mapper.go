package dto

import (
	"backend/internal/models"
	"time"
)

func ToTourModel(input *TourCreateRequest) *models.Tour {
	return &models.Tour{
		Name:         input.Name,
		Duration:     input.Duration,
		MaxGroupSize: input.MaxGroupSize,
		Difficulty:   input.Difficulty,
		Summary:      input.Summary,
		Description:  input.Description,
		Price:        input.Price,
		ImageCover:   input.ImageCover,
	}
}

func ToTourResponse(tour *models.Tour) TourResponse {
	return TourResponse{
		ID:           tour.ID,
		Name:         tour.Name,
		Duration:     tour.Duration,
		MaxGroupSize: tour.MaxGroupSize,
		Difficulty:   tour.Difficulty,
		Summary:      tour.Summary,
		Description:  tour.Description,
		Price:        tour.Price,
		ImageCover:   tour.ImageCover,
		CreatedAt:    tour.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    tour.UpdatedAt.Format(time.RFC3339),
	}
}
