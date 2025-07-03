package store

import (
	"backend/internal/models"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateAndGetUser(t *testing.T) {
	db := setupTestDB(t)
	store := NewUserStore(db)

	user := &models.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "hashedpassword", // simulate hashed pw
		Username: "testuser",
	}

	ctx := context.Background()

	err := store.CreateUser(ctx, user)
	require.NoError(t, err)

	fetchedByEmail, err := store.GetUserByEmail(ctx, user.Email)
	require.NoError(t, err)
	assert.Equal(t, user.Email, fetchedByEmail.Email)

	fetchedByUsername, err := store.GetUserByUsername(ctx, user.Username)
	require.NoError(t, err)
	assert.Equal(t, user.Username, fetchedByUsername.Username)
}

func TestGetUserNotFound(t *testing.T) {
	db := setupTestDB(t)
	store := NewUserStore(db)

	ctx := context.Background()

	_, err := store.GetUserByEmail(ctx, "nonexistent@example.com")
	assert.Error(t, err)
	assert.True(t, errors.Is(err, ErrUserNotFound))

	_, err = store.GetUserByUsername(ctx, "nonexistentuser")
	assert.Error(t, err)
	assert.True(t, errors.Is(err, ErrUserNotFound))
}
