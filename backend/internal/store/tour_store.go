package store

import (
	"backend/internal/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

var ErrTourNotFound = errors.New("tour not found")

type TourStore interface {
	CreateTour(ctx context.Context, tour *models.Tour) (*models.Tour, error)
	GetTourByID(ctx context.Context, tourID uint) (*models.Tour, error)
	UpdateTour(ctx context.Context, tour *models.Tour) error
	DeleteTour(ctx context.Context, tourID uint) error
}

type tourStore struct {
	db *gorm.DB
}

func NewTourStore(db *gorm.DB) TourStore {
	return &tourStore{db}
}

func (ts *tourStore) CreateTour(ctx context.Context, tour *models.Tour) (*models.Tour, error) {
	if err := ts.db.WithContext(ctx).Create(tour).Error; err != nil {
		return nil, err
	}
	return tour, nil
}

func (ts *tourStore) GetTourByID(ctx context.Context, tourID uint) (*models.Tour, error) {
	var tour models.Tour

	err := ts.db.WithContext(ctx).First(&tour, tourID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrTourNotFound
	}
	if err != nil {
		return nil, err
	}

	return &tour, nil
}

func (ts *tourStore) UpdateTour(ctx context.Context, tour *models.Tour) error {
	result := ts.db.WithContext(ctx).Model(&models.Tour{}).Where("id = ?", tour.ID).Updates(map[string]interface{}{
		"name":           tour.Name,
		"duration":       tour.Duration,
		"max_group_size": tour.MaxGroupSize,
		"difficulty":     tour.Difficulty,
		"summary":        tour.Summary,
		"description":    tour.Description,
		"price":          tour.Price,
		"image_cover":    tour.ImageCover,
	})
	if result.RowsAffected == 0 {
		return ErrTourNotFound
	}
	return result.Error
}

func (ts *tourStore) DeleteTour(ctx context.Context, tourID uint) error {
	result := ts.db.WithContext(ctx).Delete(&models.Tour{}, tourID)
	if result.RowsAffected == 0 {
		return ErrTourNotFound
	}
	return result.Error
}
