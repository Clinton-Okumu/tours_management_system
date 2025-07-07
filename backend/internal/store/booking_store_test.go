package store

import (
	"backend/internal/models"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func createTestUserAndTour(t *testing.T, db *gorm.DB) (*models.User, *models.Tour) {
	user := &models.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "hashedpassword", // simulate hashed pw
	}
	require.NoError(t, db.Create(user).Error)

	tour := &models.Tour{
		Name:         "Test Tour",
		Duration:     5,
		MaxGroupSize: 20,
		Difficulty:   "easy",
		Summary:      "test summary",
		Description:  "detailed description",
		Price:        500.0,
		ImageCover:   "cover.jpg",
	}
	require.NoError(t, db.Create(tour).Error)

	return user, tour
}

func TestCreateAndGetBooking(t *testing.T) {
	db := setupTestDB(t)
	store := NewBookingStore(db)

	user, tour := createTestUserAndTour(t, db)

	booking := &models.Booking{
		TourID:    tour.ID,
		UserID:    user.ID,
		StartDate: time.Date(2025, 7, 10, 0, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2025, 12, 10, 0, 0, 0, 0, time.UTC),
	}

	ctx := context.Background()

	created, err := store.CreateBooking(ctx, booking)
	require.NoError(t, err)

	fetched, err := store.GetBookingByID(ctx, created.ID)
	require.NoError(t, err)

	assert.Equal(t, booking.TourID, fetched.TourID)
	assert.Equal(t, booking.UserID, fetched.UserID)
	assert.WithinDuration(t, booking.CreatedAt, fetched.CreatedAt, time.Second)
	assert.WithinDuration(t, booking.UpdatedAt, fetched.UpdatedAt, time.Second)
}

func TestUpdateBooking(t *testing.T) {
	db := setupTestDB(t)
	store := NewBookingStore(db)

	user, tour := createTestUserAndTour(t, db)

	booking := &models.Booking{
		TourID:    tour.ID,
		UserID:    user.ID,
		StartDate: time.Date(2025, 8, 1, 0, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2025, 8, 10, 0, 0, 0, 0, time.UTC),
	}

	ctx := context.Background()
	created, err := store.CreateBooking(ctx, booking)
	require.NoError(t, err)

	created.EndDate = time.Date(2025, 8, 15, 0, 0, 0, 0, time.UTC)
	created.UpdatedAt = time.Now()

	err = store.UpdateBooking(ctx, created)
	require.NoError(t, err)

	updated, err := store.GetBookingByID(ctx, created.ID)
	require.NoError(t, err)
	assert.WithinDuration(t, created.EndDate, updated.EndDate, time.Second)
}

func TestDeleteBooking(t *testing.T) {
	db := setupTestDB(t)
	store := NewBookingStore(db)

	user, tour := createTestUserAndTour(t, db)

	booking := &models.Booking{
		TourID:    tour.ID,
		UserID:    user.ID,
		StartDate: time.Date(2025, 9, 1, 0, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2025, 9, 5, 0, 0, 0, 0, time.UTC),
	}

	ctx := context.Background()
	created, err := store.CreateBooking(ctx, booking)
	require.NoError(t, err)

	err = store.DeleteBooking(ctx, created.ID)
	require.NoError(t, err)

	_, err = store.GetBookingByID(ctx, created.ID)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, ErrBookingNotFound))
}
