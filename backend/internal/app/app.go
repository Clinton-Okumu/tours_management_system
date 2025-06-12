package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Application struct {
	Logger *log.Logger
	DB     *gorm.DB
}

func NewApplication(dbConnStr string) (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	gormDB, err := gorm.Open(postgres.Open(dbConnStr), &gorm.Config{})
	if err != nil {
		logger.Printf("Error connecting to database with GORM: %v", err)
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		logger.Printf("Error getting underlying sql.DB: %v", err)
		return nil, fmt.Errorf("error getting underlying sql.DB: %w", err)
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	if err := sqlDB.Ping(); err != nil {
		logger.Printf("Error pinging database: %v", err)
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	logger.Println("Successfully connected to database")

	app := &Application{
		Logger: logger,
		DB:     gormDB,
	}

	return app, nil
}

func (a *Application) HealthChecker(w http.ResponseWriter, r *http.Request) {
	sqlDB, err := a.DB.DB()
	if err != nil {
		a.Logger.Printf("Health check: Failed to get underlying DB for ping: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if err := sqlDB.Ping(); err != nil {
		a.Logger.Printf("Health check: Database ping failed: %v", err)
		http.Error(w, "Database not healthy", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Status is available (DB Healthy)\n")
}
