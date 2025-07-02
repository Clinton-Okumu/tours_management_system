package store

import (
	"backend/internal/models"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateAndGetBooking(t *testing.T) {
	db := setupTestDB(t)
	store := NewBookingStore(db)

	// You must have TourID and UserID already present in DB
	// For now, assume IDs 1 and 1 exist
	booking := &models.Booking{
		TourID:     1,
		UserID:     1,
		BookedAt:   time.Now(),
		TotalPrice: 1200.00,
	}

	ctx := context.Background()

	created, err := store.CreateBooking(ctx, booking)
	require.NoError(t, err)

	fetched, err := store.GetBookingByID(ctx, created.ID)
	require.NoError(t, err)

	assert.Equal(t, booking.TourID, fetched.TourID)
	assert.Equal(t, booking.UserID, fetched.UserID)
	assert.Equal(t, booking.TotalPrice, fetched.TotalPrice)
}
