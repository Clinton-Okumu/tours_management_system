package routes

import (
	"backend/internal/app"

	"github.com/go-chi/chi/v5"
)

func SetUpRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(app.Middleware.Authenticate)
		// Tours
		r.Get("/tour/{id}", app.TourHandler.GetTourByID)
		r.Post("/tour", app.TourHandler.CreateTour)
		r.Put("/tour/{id}", app.TourHandler.UpdateTour)
		r.Delete("/tour/{id}", app.TourHandler.DeleteTour)
		// Bookings
		r.Get("/booking/{id}", app.BookingHandler.GetBookingByID)
		r.Post("/booking", app.BookingHandler.CreateBooking)
		r.Put("/booking/{id}", app.BookingHandler.UpdateBooking)
		r.Delete("/booking/{id}", app.BookingHandler.DeleteBooking)
		// Locations
		r.Get("/location/{id}", app.LocationHandler.GetLocationByID)
		r.Post("/location", app.LocationHandler.CreateLocation)
		r.Put("/location/{id}", app.LocationHandler.UpdateLocation)
		r.Delete("/location/{id}", app.LocationHandler.DeleteLocation)
		// Reviews
		r.Post("/review", app.ReviewHandler.CreateReview)
		r.Delete("/review/{id}", app.ReviewHandler.DeleteReview)
		// User
		r.Post("/login", app.UserHandler.Login)
	})
	r.Get("/health", app.HealthChecker)
	r.Post("/register", app.UserHandler.Register)
	r.Post("/tokens/authentication", app.TokenHandler.HandleCreateToken)

	return r
}
