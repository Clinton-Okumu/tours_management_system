package main

import (
	"backend/internal/app"
	"backend/internal/routes"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Go backend server")
	flag.Parse()

	// Read DB connection string from environment
	godotenv.Load()
	dbConnStr := os.Getenv("POSTGRES_URL")
	if dbConnStr == "" {
		panic("POSTGRES_URL environment variable is not set")
	}

	// Initialize the application
	app, err := app.NewApplication(dbConnStr)
	if err != nil {
		panic(err)
	}

	// Close the DB connection on shutdown
	sqlDB, _ := app.DB.DB()
	defer sqlDB.Close()

	// Setup routes
	r := routes.SetupRoutes(app)

	// Start HTTP server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	app.Logger.Printf("Starting server on port %d", port)
	if err := server.ListenAndServe(); err != nil {
		app.Logger.Fatal(err)
	}
}
