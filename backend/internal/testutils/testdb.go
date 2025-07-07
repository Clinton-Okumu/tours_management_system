package testutils

import (
	"backend/internal/store"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func SetupTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	err := godotenv.Load("../../.env.test")
	require.NoError(t, err, "failed to load .env.test")

	db, err := store.Open()
	require.NoError(t, err, "failed to open DB connection")

	err = store.RunMigrations(db)
	require.NoError(t, err, "failed to run migrations")

	err = db.Exec(`TRUNCATE users, tours, bookings, reviews, locations, tokens RESTART IDENTITY CASCADE`).Error
	require.NoError(t, err, "failed to truncate tables")

	return db
}
