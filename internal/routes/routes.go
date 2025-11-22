// Package routes follows all the routing logic of the backend
package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/rshdhere/health/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)
	return r
}
