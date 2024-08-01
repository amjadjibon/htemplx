package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	appMiddleware "htemplx/app/middlewares"

	"htemplx/app/handlers"
)

func setupRouter() http.Handler {
	r := chi.NewRouter()

	// setup middlewares
	r.Use(appMiddleware.RequestID)
	r.Use(appMiddleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// setup routers
	r.Get("/healthz", handlers.Healthz)

	r.Get("/", handlers.Index)
	r.Get("/services", handlers.Services)
	r.Get("/about", handlers.About)
	r.Get("/contact", handlers.Contact)

	return r
}
