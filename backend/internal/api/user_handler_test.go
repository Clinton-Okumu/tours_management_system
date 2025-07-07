package api_test

import (
	"backend/internal/api"
	"backend/internal/models"
	"backend/internal/store"
	"backend/internal/testutils"
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupUserHandler(t *testing.T) (*chi.Mux, store.UserStore) {
	db := testutils.SetupTestDB(t)
	userStore := store.NewUserStore(db)
	logger := log.New(nil, "", 0)
	handler := api.NewUserHandler(userStore, logger)

	r := chi.NewRouter()
	r.Post("/register", handler.Register)
	r.Post("/login", handler.Login)

	return r, userStore
}

func TestRegisterHandler_Success(t *testing.T) {
	r, _ := setupUserHandler(t)

	input := map[string]string{
		"name":     "Test User",
		"email":    "test@example.com",
		"password": "securepass",
	}
	body, _ := json.Marshal(input)

	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)

	var resp map[string]any
	err := json.NewDecoder(rec.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Contains(t, resp, "user")
}

func TestLoginHandler_Success(t *testing.T) {
	r, userStore := setupUserHandler(t)

	// Register test user manually
	user := &models.User{
		Name:  "Clint",
		Email: "clint@example.com",
		Role:  "user",
	}
	err := user.SetPassword("testpass123")
	require.NoError(t, err)
	require.NoError(t, userStore.CreateUser(context.Background(), user))

	// Login request
	login := map[string]string{
		"email":    "clint@example.com",
		"password": "testpass123",
	}
	body, _ := json.Marshal(login)

	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var resp map[string]any
	err = json.NewDecoder(rec.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Contains(t, resp, "user")
}

func TestLoginHandler_InvalidPassword(t *testing.T) {
	r, userStore := setupUserHandler(t)

	user := &models.User{
		Name:  "Clint",
		Email: "wrongpass@example.com",
		Role:  "user",
	}
	_ = user.SetPassword("correctpass")
	_ = userStore.CreateUser(context.Background(), user)

	login := map[string]string{
		"email":    "wrongpass@example.com",
		"password": "wrongpass",
	}
	body, _ := json.Marshal(login)

	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}
