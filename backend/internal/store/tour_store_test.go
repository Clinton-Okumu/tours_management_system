package store

import (
	"backend/internal/models"
	"context"
	"errors"
	"testing"

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

func newTestTour() *models.Tour {
	return &models.Tour{
		Name:         "Sample Tour",
		Duration:     3,
		MaxGroupSize: 15,
		Difficulty:   "easy",
		Summary:      "A sample summary",
		Description:  "A detailed tour description.",
		Price:        599.99,
		ImageCover:   "sample.jpg",
	}
}

func TestCreateAndGetTour(t *testing.T) {
	db := setupTestDB(t)
	store := NewTourStore(db)

	tour := newTestTour()
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
}

func TestUpdateTour(t *testing.T) {
	db := setupTestDB(t)
	store := NewTourStore(db)

	ctx := context.Background()

	tour := newTestTour()
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

	tour := newTestTour()
	tour.Name = "City Break"

	createdTour, err := store.CreateTour(ctx, tour)
	require.NoError(t, err)

	err = store.DeleteTour(ctx, createdTour.ID)
	require.NoError(t, err)

	_, err = store.GetTourByID(ctx, createdTour.ID)
	require.Error(t, err)
	assert.True(t, errors.Is(err, ErrTourNotFound))
}
