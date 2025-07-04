//go:generate swag init
package main

import (
	_ "backend/docs"
	"backend/internal/app"
	"backend/internal/routes"
	"flag"
	"fmt"
	"net/http"
	"time"
)

// @title Tours API
// @version 1.0
// @description REST API for managing tours and bookings
// @termsOfService http://example.com/terms/
// @contact.name Clint Okumu
// @contact.email clint@example.com
// @host localhost:8080
// @BasePath /
func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Go backend server port")
	flag.Parse()

	// initialize the Application struct
	application, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	// Close the underlying *sql.DB used by GORM
	sqlDB, err := application.DB.DB()
	if err != nil {
		application.Logger.Fatalf("failed to get raw DB: %v", err)
	}
	defer sqlDB.Close()

	// setup routes
	r := routes.SetUpRoutes(application)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	application.Logger.Printf("🚀 Starting server on http://localhost:%d", port)
	if err := server.ListenAndServe(); err != nil {
		application.Logger.Fatal(err)
	}
}
