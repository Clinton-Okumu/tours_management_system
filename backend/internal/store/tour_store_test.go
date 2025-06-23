package store

import (
	"backend/internal/models"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	dsn := "host=localhost user=testuser password=postgres dbname=postgres port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	require.NoError(t, err)

	// Run GORM auto migrations
	err = db.AutoMigrate(&models.Tour{})
	require.NoError(t, err)

	// Clean the table
	err = db.Exec(`TRUNCATE tours RESTART IDENTITY CASCADE`).Error
	require.NoError(t, err)

	return db
}

func TestCreateAndGetTour(t *testing.T) {
	db := setupTestDB(t)
	store := NewTourStore(db)

	tour := &models.Tour{
		Name:         "Beach Getaway",
		Duration:     3,
		MaxGroupSize: 15,
		Difficulty:   "easy",
		Summary:      "Relax at the beach",
		Description:  "A calming and sunny experience by the ocean.",
		Price:        599.99,
		ImageCover:   "beach.jpg",
		StartDate:    time.Now(),
		EndDate:      time.Now().AddDate(0, 0, 3),
	}

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
	assert.Equal(t, tour.Price, fetchedTour.Price)
	assert.Equal(t, tour.StartDate, fetchedTour.StartDate)
	assert.Equal(t, tour.EndDate, fetchedTour.EndDate)
}

func TestUpdateTour(t *testing.T) {
	db := setupTestDB(t)
	store := NewTourStore(db)

	ctx := context.Background()

	originalTour := &models.Tour{
		Name:         "Beach Getaway",
		Duration:     3,
		MaxGroupSize: 15,
		Difficulty:   "easy",
		Summary:      "Relax at the beach",
		Description:  "A calming and sunny experience by the ocean.",
		Price:        599.99,
		ImageCover:   "beach.jpg",
		StartDate:    time.Now(),
		EndDate:      time.Now().AddDate(0, 0, 3),
	}

	createdTour, err := store.CreateTour(ctx, originalTour)
	require.NoError(t, err)

	createdTour.Name = "Beach Getaway 2"
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

	tour := &models.Tour{
		Name:         "City Break",
		Duration:     1,
		MaxGroupSize: 20,
		Difficulty:   "easy",
		Summary:      "Quick city tour",
		Description:  "See the city highlights in a day.",
		Price:        199.99,
		ImageCover:   "city.jpg",
		StartDate:    time.Now(),
		EndDate:      time.Now().AddDate(0, 0, 1),
	}

	created, err := store.CreateTour(ctx, tour)
	require.NoError(t, err)

	err = store.DeleteTour(ctx, created.ID)
	require.NoError(t, err)

	_, err = store.GetTourByID(ctx, created.ID)
	require.Error(t, err)
	assert.Equal(t, ErrTourNotFound, err)
}

func TestGetTourByID(t *testing.T) {
	db := setupTestDB(t)
	store := NewTourStore(db)

	ctx := context.Background()

	tour := &models.Tour{
		Name:         "Rainforest Retreat",
		Duration:     4,
		MaxGroupSize: 10,
		Difficulty:   "medium",
		Summary:      "Explore the lush green forest",
		Description:  "A guided tour deep into the rainforest.",
		Price:        749.50,
		ImageCover:   "rainforest.jpg",
		StartDate:    time.Now(),
		EndDate:      time.Now().AddDate(0, 0, 4),
	}

	created, err := store.CreateTour(ctx, tour)
	require.NoError(t, err)

	fetched, err := store.GetTourByID(ctx, created.ID)
	require.NoError(t, err)

	assert.Equal(t, created.ID, fetched.ID)
	assert.Equal(t, created.Name, fetched.Name)
	assert.Equal(t, created.Duration, fetched.Duration)
	assert.Equal(t, created.Price, fetched.Price)
}
