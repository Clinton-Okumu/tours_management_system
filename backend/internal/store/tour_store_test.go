package store

import (
	"backend/internal/models"
	"context"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	err := godotenv.Load("../../.env.test")
	require.NoError(t, err, "failed to load .env.test")

	db, err := Open()
	require.NoError(t, err, "failed to open DB connection")

	err = RunMigrations(db)
	require.NoError(t, err, "failed to run migrations")

	err = db.Exec(`TRUNCATE users, tours, bookings, reviews, locations, tokens RESTART IDENTITY CASCADE`).Error
	require.NoError(t, err, "failed to truncate tables")

	return db
}

func truncateToMicros(t time.Time) time.Time {
	return t.Round(time.Microsecond)
}

func newTestTour(start, end time.Time) *models.Tour {
	return &models.Tour{
		Name:         "Sample Tour",
		Duration:     3,
		MaxGroupSize: 15,
		Difficulty:   "easy",
		Summary:      "A sample summary",
		Description:  "A detailed tour description.",
		Price:        599.99,
		ImageCover:   "sample.jpg",
		StartDate:    start,
		EndDate:      end,
	}
}

func TestCreateAndGetTour(t *testing.T) {
	db := setupTestDB(t)
	store := NewTourStore(db)

	start := truncateToMicros(time.Now())
	end := truncateToMicros(start.AddDate(0, 0, 3))
	tour := newTestTour(start, end)

	ctx := context.Background()

	createdTour, err := store.CreateTour(ctx, tour)
	require.NoError(t, err)
	assert.NotZero(t, createdTour.ID)

	fetchedTour, err := store.GetTourByID(ctx, createdTour.ID)
	require.NoError(t, err)

	assert.Equal(t, tour.Name, fetchedTour.Name)
	assert.Equal(t, tour.Duration, fetchedTour.Duration)
	assert.Equal(t, tour.Price, fetchedTour.Price)
	assert.Equal(t, tour.Difficulty, fetchedTour.Difficulty)
	assert.Equal(t, tour.Summary, fetchedTour.Summary)
	assert.Equal(t, tour.Description, fetchedTour.Description)
	assert.Equal(t, tour.StartDate, fetchedTour.StartDate)
	assert.Equal(t, tour.EndDate, fetchedTour.EndDate)
}

func TestUpdateTour(t *testing.T) {
	db := setupTestDB(t)
	store := NewTourStore(db)

	ctx := context.Background()

	start := truncateToMicros(time.Now())
	end := truncateToMicros(start.AddDate(0, 0, 3))
	tour := newTestTour(start, end)

	createdTour, err := store.CreateTour(ctx, tour)
	require.NoError(t, err)

	createdTour.Name = "Updated Tour"
	createdTour.Price = 499.99
	createdTour.Duration = 2
	createdTour.MaxGroupSize = 10
	createdTour.Difficulty = "medium"

	err = store.UpdateTour(ctx, createdTour)
	require.NoError(t, err)

	updatedTour, err := store.GetTourByID(ctx, createdTour.ID)
	require.NoError(t, err)

	assert.Equal(t, createdTour.Name, updatedTour.Name)
	assert.Equal(t, createdTour.Price, updatedTour.Price)
	assert.Equal(t, createdTour.Duration, updatedTour.Duration)
	assert.Equal(t, createdTour.MaxGroupSize, updatedTour.MaxGroupSize)
	assert.Equal(t, createdTour.Difficulty, updatedTour.Difficulty)
}

func TestDeleteTour(t *testing.T) {
	db := setupTestDB(t)
	store := NewTourStore(db)

	ctx := context.Background()

	start := truncateToMicros(time.Now())
	end := truncateToMicros(start.AddDate(0, 0, 1))
	tour := newTestTour(start, end)
	tour.Name = "City Break"

	created, err := store.CreateTour(ctx, tour)
	require.NoError(t, err)

	err = store.DeleteTour(ctx, created.ID)
	require.NoError(t, err)

	_, err = store.GetTourByID(ctx, created.ID)
	require.Error(t, err)
	assert.Equal(t, ErrTourNotFound, err)
}

func TestGetTourByID_NotFound(t *testing.T) {
	db := setupTestDB(t)
	store := NewTourStore(db)

	ctx := context.Background()

	_, err := store.GetTourByID(ctx, 9999)
	require.Error(t, err)
	assert.Equal(t, ErrTourNotFound, err)
}
