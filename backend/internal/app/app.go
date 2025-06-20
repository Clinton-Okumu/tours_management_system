package app

import (
	"backend/internal/api"
	"backend/internal/store"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Application struct {
	DB          *gorm.DB
	Logger      *log.Logger
	TourHandler *api.TourHandler
}

func NewApplication() (*Application, error) {
	// load environment variables
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	// database connection
	db, err := store.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// automigrate models
	if err := store.RunMigrations(db); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	// setup Logger
	logger := log.New(os.Stdout, "[Tours] ", log.Ldate|log.Ltime)

	// stores
	tourStore := store.NewTourStore(db)

	// handlers
	tourHandler := api.NewTourHandler(tourStore, logger)

	app := &Application{
		DB:          db,
		Logger:      logger,
		TourHandler: tourHandler,
	}
	return app, nil
}

func (a *Application) HealthChecker(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available\n")
}
