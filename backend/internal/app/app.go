package app

import (
	"backend/internal/api"
	"backend/internal/middleware"
	"backend/internal/store"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Application struct {
	DB              *gorm.DB
	Logger          *log.Logger
	TourHandler     *api.TourHandler
	BookingHandler  *api.BookingHandler
	LocationHandler *api.LocationHandler
	ReviewHandler   *api.ReviewHandler
	UserHandler     *api.UserHandler
	TokenHandler    *api.TokenHandler
	Middleware      middleware.UserMiddleware
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
	bookingStore := store.NewBookingStore(db)
	locationStore := store.NewLocationStore(db)
	reviewStore := store.NewReviewStore(db)
	userStore := store.NewUserStore(db)
	tokenStore := store.NewTokenStore(db)

	// handlers
	tourHandler := api.NewTourHandler(tourStore, logger)
	bookingHandler := api.NewBookingHandler(bookingStore, logger)
	locationHandler := api.NewLocationHandler(locationStore, logger)
	reviewHandler := api.NewReviewHandler(reviewStore, logger)
	userHandler := api.NewUserHandler(userStore, logger)
	tokenHandler := api.NewTokenHandler(tokenStore, userStore, logger)
	middlwareHandler := middleware.UserMiddleware{UserStore: userStore}

	app := &Application{
		DB:              db,
		Logger:          logger,
		TourHandler:     tourHandler,
		BookingHandler:  bookingHandler,
		LocationHandler: locationHandler,
		ReviewHandler:   reviewHandler,
		UserHandler:     userHandler,
		TokenHandler:    tokenHandler,
		Middleware:      middlwareHandler,
	}
	return app, nil
}

func (a *Application) HealthChecker(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available\n")
}
