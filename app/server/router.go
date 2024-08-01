package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"htemplx/app/handler"
	appMiddleware "htemplx/app/middleware"
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
	r.Get("/healthz", handler.Healthz)
	r.Get("/", handler.Index)
	r.Get("/about", handler.About)
	r.Get("/contact", handler.Contact)

	return r
}
