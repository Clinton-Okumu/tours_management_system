package routes

import (
	"backend/internal/app"

	"github.com/go-chi/chi/v5"
)

func SetUpRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Get("/tour/{id}", app.TourHandler.GetTourByID)
		r.Post("/tour", app.TourHandler.CreateTour)
		r.Put("/tour/{id}", app.TourHandler.UpdateTour)
		r.Delete("/tour/{id}", app.TourHandler.DeleteTour)
	})
	r.Get("/health", app.HealthChecker)

	return r
}
