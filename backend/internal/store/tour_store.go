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
	if err := ts.db.WithContext(ctx).First(&tour, tourID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrTourNotFound
		}
		return nil, err
	}
	return &tour, nil
}

func (ts *tourStore) UpdateTour(ctx context.Context, tour *models.Tour) error {
	return ts.db.WithContext(ctx).
		Model(&models.Tour{}).
		Where("id = ?", tour.ID).
		Updates(tour).
		Error
}

func (ts *tourStore) DeleteTour(ctx context.Context, tourID uint) error {
	return ts.db.WithContext(ctx).Delete(&models.Tour{}, tourID).Error
}
