package api

import (
	"backend/internal/models"
	"backend/internal/store"
	"backend/internal/testutils"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setUpTourHandler(t *testing.T) (*chi.Mux, store.TourStore) {
	db := testutils.SetupTestDB(t)
	tourStore := store.NewTourStore(db)
	logger := log.New(io.Discard, "", 0)
	handler := NewTourHandler(tourStore, logger)

	r := chi.NewRouter()

	r.Get("/tour/{id}", handler.GetTourByID)
	r.Post("/tour", handler.CreateTour)
	r.Put("/tour/{id}", handler.UpdateTour)
	r.Delete("/tour/{id}", handler.DeleteTour)

	return r, tourStore
}

func createTestTour() models.Tour {
	return models.Tour{
		Name:         "Safari Adventure",
		Duration:     7,
		MaxGroupSize: 20,
		Difficulty:   "medium",
		Summary:      "Exciting wildlife tour",
		Description:  "Explore the savannah with expert guides",
		Price:        1500,
		ImageCover:   "safari.jpg",
	}
}

func TestCreateTour_InvalidInput(t *testing.T) {
	r, _ := setUpTourHandler(t)

	// Missing required fields
	invalidTour := map[string]any{
		"name": "Invalid Tour",
	}
	body, _ := json.Marshal(invalidTour)

	req := httptest.NewRequest(http.MethodPost, "/tour", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var resp map[string]any
	err := json.NewDecoder(rec.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Contains(t, resp, "error")
}

func TestCreateTour_Success(t *testing.T) {
	r, _ := setUpTourHandler(t)

	tour := createTestTour()
	body, _ := json.Marshal(tour)

	req := httptest.NewRequest(http.MethodPost, "/tour", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)

	var resp map[string]any
	err := json.NewDecoder(rec.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Contains(t, resp, "tour")
}

func TestGetTourByID_NotFound(t *testing.T) {
	r, _ := setUpTourHandler(t)

	req := httptest.NewRequest(http.MethodGet, "/tour/9999", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)

	var resp map[string]any
	err := json.NewDecoder(rec.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Contains(t, resp, "error")
}

func TestGetTourByID_Success(t *testing.T) {
	r, store := setUpTourHandler(t)

	tour := createTestTour()
	created, err := store.CreateTour(context.Background(), &tour)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, "/tour/"+itoa(created.ID), nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var resp map[string]any
	err = json.NewDecoder(rec.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Contains(t, resp, "tour")
}

func TestUpdateTour_NotFound(t *testing.T) {
	r, _ := setUpTourHandler(t)

	// Attempt to update a non-existent tour
	updatedTour := models.Tour{
		Name:         "NonExistent Tour",
		Duration:     10,
		MaxGroupSize: 5,
		Difficulty:   "easy",
		Summary:      "Summary",
		Description:  "Description",
		Price:        100,
		ImageCover:   "image.jpg",
	}
	body, _ := json.Marshal(updatedTour)

	req := httptest.NewRequest(http.MethodPut, "/tour/9999", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)

	var resp map[string]any
	err := json.NewDecoder(rec.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Contains(t, resp, "error")
}

func TestUpdateTour_Success(t *testing.T) {
	r, store := setUpTourHandler(t)

	tour := createTestTour()
	created, err := store.CreateTour(context.Background(), &tour)
	require.NoError(t, err)

	// Modify fields
	created.Name = "Updated Safari Adventure"
	body, _ := json.Marshal(created)

	req := httptest.NewRequest(http.MethodPut, "/tour/"+itoa(created.ID), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var resp map[string]any
	err = json.NewDecoder(rec.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Equal(t, "tour updated", resp["message"])
}

func TestDeleteTour_NotFound(t *testing.T) {
	r, _ := setUpTourHandler(t)

	req := httptest.NewRequest(http.MethodDelete, "/tour/9999", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)

	var resp map[string]any
	err := json.NewDecoder(rec.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Contains(t, resp, "error")
}

func TestDeleteTour_Success(t *testing.T) {
	r, store := setUpTourHandler(t)

	tour := createTestTour()
	created, err := store.CreateTour(context.Background(), &tour)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodDelete, "/tour/"+itoa(created.ID), nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var resp map[string]any
	err = json.NewDecoder(rec.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Equal(t, "tour deleted", resp["message"])
}

func itoa(n uint) string {
	return fmt.Sprintf("%d", n)
}
