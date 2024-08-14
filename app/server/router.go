package server

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
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
	"htemplx/pkg/auth"
	"htemplx/pkg/dbx"
	"htemplx/pkg/mailer"
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

	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:           os.Getenv("SENTRY_DSN"),
		EnableTracing: true,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for tracing.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	}); err != nil {
		slog.Warn("failed to initialize sentry", "err", err)
	}

	defer sentry.Flush(time.Minute)

	sentryMiddleware := sentryhttp.New(sentryhttp.Options{
		Repanic: true,
	})

	r.Use(sentryMiddleware.Handle)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // The url pointing to API definition
	))

	// Serve embedded files from the "public" directory
	fileServer(r, "/public", http.FS(public.AssetsFS))

	// setup routers
	r.Get("/healthz", handlers.Healthz)

	nDBX := dbx.NewDBX(
		os.Getenv("DB_URL"),
		5,
		10,
		20*time.Minute,
		30*time.Minute,
	)
	usersRepo := repo.NewUsersRepo(nDBX)
	newMailer, err := mailer.NewMailer(
		os.Getenv("SMTP_HOST"),
		587,
		os.Getenv("SMTP_FROM"),
		os.Getenv("SMTP_FROM"),
		os.Getenv("SMTP_PASSWORD"),
	)

	if err != nil {
		panic(err)
	}

	usersDomain := domain.NewUsersDomain(usersRepo, newMailer)

	contactsRepo := repo.NewContactsRepo(nDBX)
	contactsDomain := domain.NewContactsDomain(contactsRepo)

	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
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

	auth.SetupGoth(store)

	webHandler := handlers.NewWebHandler(usersDomain, contactsDomain, store)
	r.Get("/", webHandler.Index)
	r.Get("/components", webHandler.Services)
	r.Get("/about", webHandler.About)
	r.Get("/contact", webHandler.Contact)
	r.Post("/contact-submit", webHandler.ContactSubmit)
	r.Get("/login", webHandler.Login)
	r.Post("/sign-in", webHandler.SignIn)
	r.Get("/sign-out", webHandler.SignOut)
	r.Get("/register", webHandler.Register)
	r.Post("/sign-up", webHandler.SignUp)
	r.Get("/forgot-password", webHandler.ForgotPassword)
	r.Get("/forgot-password-submit", webHandler.ForgotPasswordSubmit)
	r.Get("/under-construction", webHandler.UnderConstruction)
	r.Get("/under-construction", webHandler.UnderConstruction)
	r.Get("/terms-and-conditions", webHandler.TermsAndConditions)
	r.Get("/privacy-policy", webHandler.Privacy)

	r.Get("/auth/{provider}", webHandler.GothLogin)
	r.Get("/auth/{provider}/callback", webHandler.GothCallback)

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
