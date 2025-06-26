package store

import (
	"backend/internal/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

var ErrLocationNotFound = errors.New("location not found")

type LocationStore interface {
	CreateLocation(ctx context.Context, location *models.Location) (*models.Location, error)
	GetLocationByID(ctx context.Context, locationID uint) (*models.Location, error)
	UpdateLocation(ctx context.Context, location *models.Location) error
	DeleteLocation(ctx context.Context, locationID uint) error
}

type locationStore struct {
	db *gorm.DB
}

func NewLocationStore(db *gorm.DB) LocationStore {
	return &locationStore{db}
}

func (ls *locationStore) CreateLocation(ctx context.Context, location *models.Location) (*models.Location, error) {
	if err := ls.db.WithContext(ctx).Create(location).Error; err != nil {
		return nil, err
	}
	return location, nil
}

func (ls *locationStore) GetLocationByID(ctx context.Context, locationID uint) (*models.Location, error) {
	var location models.Location
	if err := ls.db.WithContext(ctx).First(&location, locationID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrLocationNotFound
		}
		return nil, err
	}
	return &location, nil
}

func (ls *locationStore) UpdateLocation(ctx context.Context, location *models.Location) error {
	return ls.db.WithContext(ctx).
		Model(&models.Location{}).
		Where("id = ?", location.ID).
		Updates(location).
		Error
}

func (ls *locationStore) DeleteLocation(ctx context.Context, locationID uint) error {
	return ls.db.WithContext(ctx).Delete(&models.Location{}, locationID).Error
}
