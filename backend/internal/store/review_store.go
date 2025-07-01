package store

import (
	"backend/internal/models"
	"context"

	"gorm.io/gorm"
)

type ReviewStore interface {
	CreateReview(ctx context.Context, review *models.Review) (*models.Review, error)
	DeleteReview(ctx context.Context, reviewID uint) error
}

type reviewStore struct {
	db *gorm.DB
}

func NewReviewStore(db *gorm.DB) ReviewStore {
	return &reviewStore{db}
}

func (rs *reviewStore) CreateReview(ctx context.Context, review *models.Review) (*models.Review, error) {
	if err := rs.db.WithContext(ctx).Create(review).Error; err != nil {
		return nil, err
	}
	return review, nil
}

func (rs *reviewStore) DeleteReview(ctx context.Context, reviewID uint) error {
	return rs.db.WithContext(ctx).Delete(&models.Review{}, reviewID).Error
}
