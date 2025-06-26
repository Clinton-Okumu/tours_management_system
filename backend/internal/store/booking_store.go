package store

import (
	"backend/internal/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

var ErrBookingNotFound = errors.New("booking not found")

type BookingStore interface {
	CreateBooking(ctx context.Context, booking *models.Booking) (*models.Booking, error)
	GetBookingByID(ctx context.Context, bookingID uint) (*models.Booking, error)
	UpdateBooking(ctx context.Context, booking *models.Booking) error
	DeleteBooking(ctx context.Context, bookingID uint) error
}

type bookingStore struct {
	db *gorm.DB
}

func NewBookingStore(db *gorm.DB) BookingStore {
	return &bookingStore{db}
}

func (bs *bookingStore) CreateBooking(ctx context.Context, booking *models.Booking) (*models.Booking, error) {
	if err := bs.db.WithContext(ctx).Create(booking).Error; err != nil {
		return nil, err
	}
	return booking, nil
}

func (bs *bookingStore) GetBookingByID(ctx context.Context, bookingID uint) (*models.Booking, error) {
	var booking models.Booking
	if err := bs.db.WithContext(ctx).First(&booking, bookingID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrBookingNotFound
		}
		return nil, err
	}
	return &booking, nil
}

func (bs *bookingStore) UpdateBooking(ctx context.Context, booking *models.Booking) error {
	return bs.db.WithContext(ctx).
		Model(&models.Booking{}).
		Where("id = ?", booking.ID).
		Updates(booking).
		Error
}

func (bs *bookingStore) DeleteBooking(ctx context.Context, bookingID uint) error {
	return bs.db.WithContext(ctx).Delete(&models.Booking{}, bookingID).Error
}
