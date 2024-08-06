package server

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/rbcervilla/redisstore/v9"
	"github.com/redis/go-redis/v9"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	_ "htemplx/app/docs"
	"htemplx/app/domain"
	"htemplx/app/handlers"
	"htemplx/app/repo"
	"htemplx/pkg/dbx"
	"htemplx/pkg/middlewares"
	"htemplx/public"
)

func setupRouter() http.Handler {
	r := chi.NewRouter()

	// setup middlewares
	r.Use(middlewares.RequestID)
	r.Use(middlewares.Logger)
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

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // The url pointing to API definition
	))

	// Serve embedded files from the "public" directory
	fileServer(r, "/public", http.FS(public.AssetsFS))

	// setup routers
	r.Get("/healthz", handlers.Healthz)

	nDBX := dbx.NewDBX(
		"postgres://rootuser:rootpassword@localhost:5432/htemplx_db?sslmode=disable",
		5,
		10,
		20*time.Minute,
		30*time.Minute,
	)
	usersRepo := repo.NewUsersRepo(nDBX)
	usersDomain := domain.NewUsersDomain(usersRepo)

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "rootpassword",
	})

	// New default RedisStore
	store, err := redisstore.NewRedisStore(context.Background(), client)
	if err != nil {
		panic(err)
	}

	// Example changing configuration for sessions
	store.KeyPrefix("session_")
	// store.Options(sessions.Options{
	// 	Path:   "/path",
	// 	Domain: "example.com",
	// 	MaxAge: 86400 * 60,
	// })

	webHandler := handlers.NewWebHandler(usersDomain, store)
	r.Get("/", webHandler.Index)
	r.Get("/services", webHandler.Services)
	r.Get("/about", webHandler.About)
	r.Get("/contact", webHandler.Contact)
	r.Get("/login", webHandler.Login)
	r.Post("/sign-in", webHandler.SignIn)
	r.Get("/register", webHandler.Register)
	r.Post("/sign-up", webHandler.SignUp)
	r.Get("/forgot-password", webHandler.ForgotPassword)
	r.Get("/under-construction", webHandler.UnderConstruction)
	r.Get("/under-construction", webHandler.UnderConstruction)
	r.Get("/terms-and-conditions", webHandler.TermsAndConditions)
	r.Get("/privacy-policy", webHandler.Privacy)

	// default not found page
	r.NotFound(webHandler.NotFound)

	apiHandler := handlers.NewApiHandler(usersDomain)
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/users", apiHandler.CreateUser)
		r.Get("/users", apiHandler.GetUserList)
		r.Get("/users/{id}", apiHandler.GetUserByID)
		r.Put("/users/{id}", apiHandler.UpdateUser)
		r.Delete("/users/{id}", apiHandler.DeleteUser)
	})

	return r
}

func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("fileServer does not permit any URL parameters")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
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
