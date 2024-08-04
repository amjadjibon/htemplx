package server

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"htemplx/app/handlers"
	appMiddleware "htemplx/app/middlewares"
	"htemplx/public"
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

	// Serve embedded files from the "public" directory
	fileServer(r, "/public", http.FS(public.AssetsFS))

	// setup routers
	r.Get("/healthz", handlers.Healthz)

	r.Get("/", handlers.Index)
	r.Get("/services", handlers.Services)
	r.Get("/about", handlers.About)
	r.Get("/contact", handlers.Contact)

	r.Get("/login", handlers.Login)
	r.Get("/register", handlers.Register)
	r.Get("/forgot-password", handlers.ForgotPassword)
	r.Get("/under-construction", handlers.UnderConstruction)
	r.Get("/under-construction", handlers.UnderConstruction)
	r.Get("/terms-and-conditions", handlers.TermsAndConditions)
	r.Get("/privacy-policy", handlers.Privacy)

	// default not found page
	r.NotFound(handlers.NotFound)

	return r
}

func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("fileServer does not permit any URL parameters")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rCtx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rCtx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
