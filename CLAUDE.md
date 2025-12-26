# CLAUDE.md - AI Assistant Developer Guide

This document provides comprehensive guidance for AI assistants working with the htemplx codebase. It covers architecture, conventions, workflows, and best practices.

## Table of Contents

1. [Project Overview](#project-overview)
2. [Architecture](#architecture)
3. [Tech Stack](#tech-stack)
4. [Directory Structure](#directory-structure)
5. [Development Workflows](#development-workflows)
6. [Code Conventions](#code-conventions)
7. [Common Tasks](#common-tasks)
8. [Testing](#testing)
9. [Deployment](#deployment)
10. [Important Notes](#important-notes)

---

## Project Overview

**htemplx** is a modern Go web application using the "HTML over the wire" pattern with HTMX and Templ templates. It combines traditional server-side rendering with dynamic HTMX interactivity, eliminating the need for a separate frontend framework.

### Key Features
- Server-side rendering with Templ (type-safe Go templates)
- HTMX for dynamic client-side interactions
- Dual API support: Web (HTMX) and REST (JSON)
- PostgreSQL database with migrations
- Redis-backed sessions
- OAuth authentication (Google)
- Tailwind CSS for styling
- Swagger API documentation
- Docker and Kubernetes deployment ready

---

## Architecture

The project follows **Clean Architecture** principles with clear separation of concerns:

```
HTTP Request → Handler → Domain (Business Logic) → Repository → Database
                  ↓
               Response (Templ Template or JSON)
```

### Layers

1. **Handlers** (`app/handlers/`)
   - `WebHandler`: Renders HTMX pages using Templ templates
   - `ApiHandler`: Returns JSON responses for REST API
   - Both handle HTTP requests and delegate to domain layer

2. **Domain** (`app/domain/`)
   - Contains business logic
   - `UsersDomain`: User management, auth, password reset
   - `ContactsDomain`: Contact form processing
   - Orchestrates between handlers and repositories

3. **Repository** (`app/repo/`)
   - Data access layer
   - `UsersRepo`: Database operations for users
   - `ContactsRepo`: Database operations for contacts
   - Uses Squirrel for type-safe SQL building

4. **Models** (`app/models/`)
   - Database entity definitions
   - Matches database schema structure

5. **DTOs** (`app/dto/`)
   - Data Transfer Objects for request/response validation

6. **Views** (`app/views/`)
   - `layouts/`: Base HTML structure
   - `pages/`: Full page templates
   - `components/`: Reusable UI components

### Dependencies Flow

```
cmd/ → server/ → handlers/ → domain/ → repo/ → pkg/dbx/
                      ↓
                   views/
```

---

## Tech Stack

### Backend
- **Go 1.22**: Primary language
- **Chi v5**: HTTP router and middleware
- **Squirrel**: Type-safe SQL query builder
- **sqlx**: Database operations with pgx driver
- **PostgreSQL**: Primary database (pgx driver)
- **Redis**: Session storage
- **Goose v3**: Database migrations

### Frontend
- **HTMX v1.x**: Dynamic HTML updates
- **Templ v0.2.x**: Type-safe Go templates
- **Tailwind CSS**: Utility-first CSS framework
- **Hyperscript**: Declarative client-side scripting
- **Flowbite**: UI component library

### Authentication
- **gorilla/sessions**: Session management
- **redisstore**: Redis-backed session store
- **bcrypt**: Password hashing (cost factor 14)
- **markbates/goth**: OAuth authentication (Google)

### DevOps & Tools
- **Air**: Hot reloading for development
- **Swagger/Swag**: API documentation generation
- **Sentry**: Error tracking and monitoring
- **Docker**: Containerization
- **Kubernetes**: Orchestration

### Email & Utilities
- **go-mail**: SMTP email sending
- **slog**: Structured logging
- **uuid**: UUID generation
- **go-password**: Secure password generation

---

## Directory Structure

```
/home/user/htemplx/
├── app/                       # Main application code
│   ├── conf/                  # Environment configuration
│   │   └── env.go            # Config struct with env tags
│   ├── docs/                  # Generated Swagger docs
│   ├── domain/                # Business logic layer
│   │   ├── users.go          # User management logic
│   │   └── contacts.go       # Contact form logic
│   ├── dto/                   # Data Transfer Objects
│   │   ├── users.go          # User DTOs
│   │   └── contacts.go       # Contact DTOs
│   ├── handlers/              # HTTP request handlers
│   │   ├── index.go          # WebHandler (HTMX)
│   │   ├── api.go            # ApiHandler (JSON)
│   │   ├── health.go         # Health check endpoint
│   │   └── render.go         # Templ rendering utility
│   ├── models/                # Database models
│   │   ├── users.go          # User model
│   │   └── contacts.go       # Contact model
│   ├── repo/                  # Data access layer
│   │   ├── users.go          # User repository
│   │   └── contacts.go       # Contact repository
│   ├── server/                # HTTP server setup
│   │   ├── http.go           # Server initialization
│   │   ├── router.go         # Route definitions
│   │   └── static.go         # Static file serving
│   └── views/                 # Templ templates
│       ├── components/        # Reusable UI components
│       │   ├── navbar.templ
│       │   ├── login.templ
│       │   ├── register.templ
│       │   └── ...
│       ├── layouts/           # Page layouts
│       │   └── base.templ    # Base HTML structure
│       └── pages/             # Full page templates
│           ├── index.templ
│           ├── about.templ
│           └── ...
├── cmd/                       # CLI commands (Cobra)
│   ├── root.go               # Root command
│   ├── serve.go              # Serve HTTP server
│   ├── migrate.go            # Database migrations
│   └── version.go            # Version command
├── pkg/                       # Shared packages
│   ├── auth/                  # Authentication utilities
│   │   ├── password.go       # Bcrypt hashing
│   │   └── goth.go          # OAuth setup
│   ├── dbx/                   # Database wrapper
│   │   └── dbx.go           # sqlx + Squirrel setup
│   ├── logger/                # Logging utilities
│   │   └── logger.go        # slog configuration
│   ├── mailer/                # Email service
│   │   ├── mailer.go        # SMTP client
│   │   └── mailer_test.go   # Tests
│   └── middlewares/           # HTTP middlewares
│       ├── request_id.go    # Request ID injection
│       └── logger.go        # Request logging
├── migrations/                # Database migrations
│   └── postgres/
│       └── *.sql            # Goose migration files
├── distribution/              # Deployment configs
│   ├── docker/
│   │   └── Dockerfile       # Multi-stage Docker build
│   ├── compose/
│   │   ├── docker-compose.yml       # Production
│   │   └── docker-compose-dev.yml   # Development
│   └── k8s/                  # Kubernetes manifests
│       ├── htemplx.yaml     # App deployment
│       ├── postgres.yaml    # PostgreSQL
│       └── redis.yaml       # Redis
├── public/                    # Static assets (embedded)
│   └── assets/
│       ├── css/
│       │   ├── input.css    # Tailwind input
│       │   └── output.css   # Compiled CSS
│       ├── js/              # HTMX, Hyperscript, Flowbite
│       └── logo/            # Logo files
├── tools/                     # Development tools
├── docs/                      # Documentation
├── .air.toml                 # Air hot reload config
├── go.mod                    # Go dependencies
├── go.sum                    # Dependency checksums
├── Makefile                  # Build automation
├── main.go                   # Entry point
├── tailwind.config.js        # Tailwind configuration
└── postcss.config.js         # PostCSS configuration
```

### File Naming Conventions

- **Templ files**: `*.templ` (generate `*_templ.go`)
- **Test files**: `*_test.go`
- **Go files**: `snake_case.go`
- **Migrations**: `YYYYMMDDHHMMSS_description.sql`

---

## Development Workflows

### Initial Setup

1. **Install dependencies**:
   ```bash
   # Install Air for hot reloading
   make air-install

   # Install Templ CLI
   make templ-install

   # Install Node.js dependencies for Tailwind
   npm install
   ```

2. **Set up environment variables**:
   ```bash
   # Required
   export DB_URL="postgres://user:pass@localhost:5432/dbname?sslmode=disable"
   export REDIS_URL="localhost:6379"

   # Optional
   export SMTP_HOST="smtp.gmail.com"
   export SMTP_FROM="your-email@gmail.com"
   export SMTP_PASSWORD="your-password"
   export GOOGLE_CLIENT_ID="..."
   export GOOGLE_CLIENT_SECRET="..."
   export GOOGLE_CALLBACK="http://localhost:8080/auth/google/callback"
   export SENTRY_DSN="..."
   ```

3. **Start development services**:
   ```bash
   # Start PostgreSQL and Redis
   make docker-up-dev
   ```

4. **Run migrations**:
   ```bash
   make migrate-up
   ```

5. **Start development server**:
   ```bash
   make run  # Uses Air for hot reloading
   ```

### Development Loop (Air)

Air watches for changes and automatically:
1. Formats Swagger comments: `swag fmt`
2. Generates Swagger docs: `swag init`
3. Generates Templ templates: `templ generate`
4. Compiles Tailwind CSS
5. Builds and restarts the Go binary

**Watched extensions**: `.go`, `.templ`, `.html`
**Excluded**: `*_test.go`, `*_templ.go`, `tmp/`, `vendor/`

**Build command** (from `.air.toml:8`):
```bash
swag fmt && swag init --output ./app/docs && templ generate && \
npx tailwindcss -i ./public/assets/css/input.css -o ./public/assets/css/output.css && \
go build -o ./tmp/main main.go
```

### Manual Build

```bash
# Clean build artifacts
make clean

# Build binary
make build  # Output: bin/htemplx

# Run binary
make serve
```

### Working with Templates

1. **Create/edit `.templ` files** in `app/views/`
2. **Generate Go code**:
   ```bash
   make templ-generate
   # Or let Air do it automatically
   ```
3. **Use in handlers**:
   ```go
   render(w, r, pages.Index(loggedIn, "htemplx"))
   ```

### Working with Tailwind CSS

1. **Edit classes** in `.templ` files
2. **Compile CSS**:
   ```bash
   make tailwind-gen
   # Or let Air do it automatically
   ```
3. **Output**: `public/assets/css/output.css`

### Database Migrations

**Create new migration**:
```bash
cd migrations/postgres
goose create my_migration sql
```

**Run migrations**:
```bash
make migrate-up    # Apply all pending
make migrate-down  # Rollback last migration
```

**Or use CLI directly**:
```bash
go run main.go migrate up --db_url="postgres://..."
go run main.go migrate down --db_url="postgres://..."
```

---

## Code Conventions

### Handlers

**WebHandler pattern** (app/handlers/index.go):
```go
func (h *WebHandler) MyPage(w http.ResponseWriter, r *http.Request) {
    // 1. Check authentication if needed
    loggedIn, err := h.IsLoggedIn(r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // 2. Get/process data via domain layer
    data, err := h.someDomain.GetData(r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // 3. Render Templ template
    render(w, r, pages.MyPage(loggedIn, data))
}
```

**ApiHandler pattern** (app/handlers/api.go):
```go
// @Summary     Short description
// @Description Detailed description
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       user body dto.CreateUserRequest true "User data"
// @Success     200 {object} dto.CreateUserResponse
// @Router      /users [post]
func (h *ApiHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    // Similar pattern but return JSON
}
```

### Domain Layer

**Pattern** (app/domain/users.go):
```go
type UsersDomain struct {
    usersRepo *repo.UsersRepo
    mailer    *mailer.Mailer
}

func (d *UsersDomain) CreateUser(r *http.Request) error {
    // 1. Parse/validate input
    // 2. Business logic
    // 3. Call repository
    // 4. Additional actions (send email, etc.)
    return nil
}
```

### Repository Layer

**Pattern** (app/repo/users.go):
```go
func (r *UsersRepo) Create(user *models.User) error {
    query, args, err := squirrel.
        Insert("users").
        Columns("id", "email", "password", ...).
        Values(user.ID, user.Email, user.Password, ...).
        PlaceholderFormat(squirrel.Dollar).
        ToSql()

    if err != nil {
        return err
    }

    _, err = r.dbx.Exec(query, args...)
    return err
}
```

**Always use**:
- `squirrel.Dollar` for PostgreSQL placeholders
- Parameterized queries (prevent SQL injection)
- Proper error handling

### Templ Templates

**Component definition** (app/views/components/alert.templ):
```templ
package components

templ Alert(title string, message string) {
    <div class="bg-green-100 border border-green-400 text-green-700 px-4 py-3 rounded">
        <strong class="font-bold">{ title }</strong>
        <span class="block sm:inline">{ message }</span>
    </div>
}
```

**Using HTMX** (app/views/components/login.templ):
```templ
<form hx-post="/sign-in"
      hx-target="#response"
      hx-swap="innerHTML"
      hx-ext="response-targets">
    <input type="email" name="email" required />
    <input type="password" name="password" required />
    <button type="submit">Login</button>
</form>
<div id="response"></div>
```

**Key HTMX attributes**:
- `hx-post`, `hx-get`: HTTP requests
- `hx-target`: Where to insert response
- `hx-swap`: How to insert (innerHTML, outerHTML, etc.)
- `hx-ext="response-targets"`: Error handling extension

### Error Handling

**Pattern**:
```go
if err != nil {
    slog.Error("operation failed", "error", err, "context", value)
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}
```

**Use structured logging** with slog:
```go
slog.Info("user created", "id", user.ID, "email", user.Email)
slog.Warn("rate limit exceeded", "ip", r.RemoteAddr)
slog.Error("database error", "error", err, "query", query)
```

### Authentication

**Check login status**:
```go
loggedIn, err := h.IsLoggedIn(r)
```

**Session management** (app/handlers/index.go):
```go
session, _ := h.sessionStore.Get(r, "auth")
session.Values["authenticated"] = true
session.Save(r, w)
```

**OAuth flow** (app/handlers/index.go):
```go
// Initiate: /auth/google
// Callback: /auth/google/callback
user, err := auth.GothicLogin(w, r)
```

### Configuration

**Environment variables** (app/conf/env.go):
```go
type Config struct {
    HttpHost         string        `env:"HOST" envDefault:"0.0.0.0"`
    HttpPort         int           `env:"PORT" envDefault:"8080"`
    LogLevel         string        `env:"LOG_LEVEL" envDefault:"debug"`
    HttpReadTimeout  time.Duration `env:"HTTP_READ_TIMEOUT" envDefault:"10s"`
    HttpWriteTimeout time.Duration `env:"HTTP_WRITE_TIMEOUT" envDefault:"10s"`
}
```

**Load config**:
```go
cfg := conf.NewConfig()  // Automatically parses env vars
```

---

## Common Tasks

### Adding a New Page

1. **Create Templ template** (`app/views/pages/mypage.templ`):
   ```templ
   package pages

   import "htemplx/app/views/layouts"

   templ MyPage(loggedIn bool, title string) {
       @layouts.Base(loggedIn, title) {
           <h1>My Page</h1>
           <p>Content here</p>
       }
   }
   ```

2. **Add handler** (`app/handlers/index.go`):
   ```go
   func (h *WebHandler) MyPage(w http.ResponseWriter, r *http.Request) {
       loggedIn, err := h.IsLoggedIn(r)
       if err != nil {
           http.Error(w, err.Error(), http.StatusInternalServerError)
           return
       }
       render(w, r, pages.MyPage(loggedIn, "My Page"))
   }
   ```

3. **Add route** (`app/server/router.go`):
   ```go
   r.Get("/mypage", webHandler.MyPage)
   ```

### Adding a New API Endpoint

1. **Create DTO** (`app/dto/myresource.go`):
   ```go
   type CreateMyResourceRequest struct {
       Name string `json:"name" validate:"required"`
   }

   type CreateMyResourceResponse struct {
       ID   string `json:"id"`
       Name string `json:"name"`
   }
   ```

2. **Create model** (`app/models/myresource.go`):
   ```go
   type MyResource struct {
       ID        uuid.UUID  `db:"id"`
       Name      string     `db:"name"`
       CreatedAt time.Time  `db:"created_at"`
   }
   ```

3. **Create repository** (`app/repo/myresource.go`):
   ```go
   type MyResourceRepo struct {
       dbx *dbx.DBX
   }

   func NewMyResourceRepo(dbx *dbx.DBX) *MyResourceRepo {
       return &MyResourceRepo{dbx: dbx}
   }

   func (r *MyResourceRepo) Create(resource *models.MyResource) error {
       query, args, err := squirrel.
           Insert("my_resources").
           Columns("id", "name", "created_at").
           Values(resource.ID, resource.Name, resource.CreatedAt).
           PlaceholderFormat(squirrel.Dollar).
           ToSql()

       if err != nil {
           return err
       }

       _, err = r.dbx.Exec(query, args...)
       return err
   }
   ```

4. **Create domain** (`app/domain/myresource.go`):
   ```go
   type MyResourceDomain struct {
       repo *repo.MyResourceRepo
   }

   func NewMyResourceDomain(repo *repo.MyResourceRepo) *MyResourceDomain {
       return &MyResourceDomain{repo: repo}
   }

   func (d *MyResourceDomain) CreateResource(req dto.CreateMyResourceRequest) (*models.MyResource, error) {
       resource := &models.MyResource{
           ID:        uuid.New(),
           Name:      req.Name,
           CreatedAt: time.Now(),
       }

       err := d.repo.Create(resource)
       return resource, err
   }
   ```

5. **Add handler** (`app/handlers/api.go`):
   ```go
   // @Summary     Create resource
   // @Description Creates a new resource
   // @Tags        resources
   // @Accept      json
   // @Produce     json
   // @Param       resource body dto.CreateMyResourceRequest true "Resource data"
   // @Success     200 {object} dto.CreateMyResourceResponse
   // @Router      /resources [post]
   func (h *ApiHandler) CreateResource(w http.ResponseWriter, r *http.Request) {
       var req dto.CreateMyResourceRequest
       if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
           http.Error(w, err.Error(), http.StatusBadRequest)
           return
       }

       resource, err := h.resourceDomain.CreateResource(req)
       if err != nil {
           http.Error(w, err.Error(), http.StatusInternalServerError)
           return
       }

       resp := dto.CreateMyResourceResponse{
           ID:   resource.ID.String(),
           Name: resource.Name,
       }

       w.Header().Set("Content-Type", "application/json")
       json.NewEncoder(w).Encode(resp)
   }
   ```

6. **Wire up dependencies** (`app/server/router.go`):
   ```go
   resourceRepo := repo.NewMyResourceRepo(nDBX)
   resourceDomain := domain.NewMyResourceDomain(resourceRepo)
   apiHandler := handlers.NewApiHandler(..., resourceDomain)

   r.Route("/api/v1", func(r chi.Router) {
       r.Post("/resources", apiHandler.CreateResource)
   })
   ```

7. **Create migration**:
   ```bash
   cd migrations/postgres
   goose create create_my_resources_table sql
   ```

   Edit the SQL file:
   ```sql
   -- +goose Up
   CREATE TABLE my_resources (
       id UUID PRIMARY KEY,
       name VARCHAR(255) NOT NULL,
       created_at TIMESTAMP NOT NULL
   );

   -- +goose Down
   DROP TABLE my_resources;
   ```

8. **Run migration**:
   ```bash
   make migrate-up
   ```

### Adding a New Component

1. **Create component** (`app/views/components/mycomponent.templ`):
   ```templ
   package components

   templ MyComponent(title string, items []string) {
       <div class="p-4 bg-white rounded shadow">
           <h3 class="text-lg font-bold">{ title }</h3>
           <ul class="mt-2">
               for _, item := range items {
                   <li>{ item }</li>
               }
           </ul>
       </div>
   }
   ```

2. **Use in page**:
   ```templ
   @components.MyComponent("My List", []string{"Item 1", "Item 2"})
   ```

### Adding Middleware

1. **Create middleware** (`pkg/middlewares/mymiddleware.go`):
   ```go
   func MyMiddleware(next http.Handler) http.Handler {
       return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
           // Do something before
           next.ServeHTTP(w, r)
           // Do something after
       })
   }
   ```

2. **Register in router** (`app/server/router.go`):
   ```go
   r.Use(middlewares.MyMiddleware)
   ```

---

## Testing

### Current Test Coverage

Currently, the project has **minimal test coverage**:
- Only `pkg/mailer/mailer_test.go` exists
- Uses `testify/suite` for test organization

### Writing Tests

**Test file pattern** (`pkg/mypackage/myfile_test.go`):
```go
package mypackage

import (
    "testing"

    "github.com/stretchr/testify/suite"
)

type MyTestSuite struct {
    suite.Suite
    // Test dependencies
}

func (suite *MyTestSuite) SetupSuite() {
    // Setup once before all tests
}

func (suite *MyTestSuite) SetupTest() {
    // Setup before each test
}

func (suite *MyTestSuite) TestMyFunction() {
    result := MyFunction("input")
    suite.Equal("expected", result)
}

func TestMyTestSuite(t *testing.T) {
    suite.Run(t, new(MyTestSuite))
}
```

**Run tests**:
```bash
go test ./...
go test -v ./pkg/mailer
go test -cover ./...
```

### Testing Best Practices

1. **Unit tests**: Test individual functions/methods
2. **Integration tests**: Test interactions between layers
3. **Use testify assertions**: `suite.Equal()`, `suite.Nil()`, etc.
4. **Mock dependencies**: Use interfaces for mockability
5. **Test file naming**: `*_test.go`
6. **Exclude from Air**: Already configured in `.air.toml:12`

---

## Deployment

### Docker

**Build image**:
```bash
make docker-build  # Uses distribution/docker/Dockerfile
```

**Run production**:
```bash
make docker-up  # PostgreSQL + Redis + htemplx
```

**Run development** (services only):
```bash
make docker-up-dev  # PostgreSQL + Redis (app runs locally)
```

**Dockerfile structure** (distribution/docker/Dockerfile):
```dockerfile
# Stage 1: Build
FROM golang:1.22 AS builder
# ... build steps ...

# Stage 2: Runtime
FROM scratch
COPY --from=builder /usr/local/bin/htemplx /usr/local/bin/htemplx
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/htemplx", "serve"]
```

### Kubernetes

**Deploy**:
```bash
kubectl apply -f distribution/k8s/postgres.yaml
kubectl apply -f distribution/k8s/redis.yaml
kubectl apply -f distribution/k8s/htemplx.yaml
```

**Access**:
- NodePort: `http://<node-ip>:31000`

**Manifests**:
- `htemplx.yaml`: App deployment + NodePort service
- `postgres.yaml`: PostgreSQL deployment + ClusterIP
- `redis.yaml`: Redis deployment + ClusterIP

**Important**: Current K8s configs use `emptyDir` (ephemeral storage). For production, use PersistentVolumes.

### Environment Variables for Production

```bash
# Required
DB_URL=postgres://user:pass@postgres:5432/dbname?sslmode=disable
REDIS_URL=redis:6379

# Optional but recommended
SENTRY_DSN=https://...
LOG_LEVEL=info

# For OAuth
GOOGLE_CLIENT_ID=...
GOOGLE_CLIENT_SECRET=...
GOOGLE_CALLBACK=https://yourdomain.com/auth/google/callback

# For email
SMTP_HOST=smtp.sendgrid.net
SMTP_FROM=noreply@yourdomain.com
SMTP_PASSWORD=...

# Server config
HOST=0.0.0.0
PORT=8080
HTTP_READ_TIMEOUT=10s
HTTP_WRITE_TIMEOUT=10s
```

---

## Important Notes

### Security Considerations

1. **CSRF Protection**: Currently **NOT implemented**. Consider adding CSRF middleware.
2. **CORS**: Currently allows all origins (`*`). Restrict in production (app/server/router.go:40).
3. **Password Hashing**: Uses bcrypt with cost 14 (secure).
4. **SQL Injection**: Protected via Squirrel parameterized queries.
5. **Session Security**: Redis-backed sessions, consider adding secure cookie flags.
6. **Rate Limiting**: Not implemented. Consider adding for production.

### Performance Optimizations

1. **Embedded assets**: Static files embedded in binary (app/server/static.go).
2. **Database pooling**: Configured in dbx (5 idle, 10 max connections).
3. **Redis sessions**: Scalable session storage.
4. **Multi-stage Docker**: Minimal image size (~10MB).
5. **Graceful shutdown**: 30-second timeout for draining connections.

### Monitoring

1. **Sentry**: Error tracking enabled (100% trace sampling - adjust for production).
2. **Structured logging**: JSON logs via slog.
3. **Request IDs**: Automatically injected via middleware.
4. **Health check**: `/healthz` endpoint.

### Common Pitfalls

1. **Forgetting to generate Templ**: Always run `templ generate` or use Air.
2. **Database connection string**: Must include `sslmode=disable` for local dev.
3. **Redis connection**: Format is `host:port`, not a full URL.
4. **Static assets**: Must be embedded; changes require rebuild.
5. **HTMX responses**: Return HTML fragments, not full pages.
6. **Migration direction**: `migrate down` only rolls back **one** migration.

### Key Files to Monitor

- `.air.toml`: Hot reload configuration
- `app/server/router.go`: All routes and middleware
- `app/conf/env.go`: Environment configuration
- `go.mod`: Dependencies
- `migrations/postgres/*`: Database schema

### Development Tips

1. **Use Air**: It's configured to handle everything automatically.
2. **Check build logs**: `tmp/build-errors.log` for Air build failures.
3. **Swagger UI**: Available at `/swagger/index.html` after running.
4. **Request logging**: All requests logged with request ID.
5. **Database debugging**: Enable SQL logging in dbx if needed.

### Code Generation

Several files are **auto-generated** and should not be edited directly:

- `*_templ.go`: Generated from `*.templ` files
- `app/docs/*`: Generated from Swagger comments
- `public/assets/css/output.css`: Generated from Tailwind input

### Recommended Additions

For a production-ready application, consider:

1. **Testing**: Add comprehensive unit and integration tests
2. **CSRF Protection**: Add CSRF middleware
3. **Rate Limiting**: Prevent abuse
4. **Input Validation**: Add validation middleware/library
5. **Email Templates**: Use Templ for HTML emails
6. **Admin Panel**: User management interface
7. **Monitoring Dashboard**: Grafana/Prometheus integration
8. **CI/CD Pipeline**: GitHub Actions or similar
9. **Environment Configs**: Separate dev/staging/prod configs
10. **Database Backups**: Automated backup strategy

---

## Quick Reference

### Makefile Commands
```bash
make clean          # Remove build artifacts
make build          # Build binary
make serve          # Run binary
make run            # Run with Air (hot reload)
make migrate-up     # Apply migrations
make migrate-down   # Rollback last migration
make templ-generate # Generate Templ templates
make tailwind-gen   # Compile Tailwind CSS
make docker-build   # Build Docker image
make docker-up      # Run production compose
make docker-up-dev  # Run dev compose
```

### CLI Commands
```bash
go run main.go serve                    # Start server
go run main.go migrate up --db_url=...  # Run migrations
go run main.go migrate down --db_url=...# Rollback migration
go run main.go version                  # Show version
```

### Common URLs
- App: http://localhost:8080
- Swagger: http://localhost:8080/swagger/index.html
- Health check: http://localhost:8080/healthz

### Key Dependencies
- Router: `github.com/go-chi/chi/v5`
- Templates: `github.com/a-h/templ`
- Database: `github.com/jmoiron/sqlx` + `github.com/jackc/pgx/v5`
- Queries: `github.com/Masterminds/squirrel`
- Migrations: `github.com/pressly/goose/v3`
- Sessions: `github.com/gorilla/sessions` + `github.com/rbcervilla/redisstore/v9`
- OAuth: `github.com/markbates/goth`
- Email: `github.com/wneessen/go-mail`
- CLI: `github.com/spf13/cobra`
- Logging: `log/slog` (standard library)

---

## Summary for AI Assistants

When working with this codebase:

1. **Follow the layered architecture**: Handlers → Domain → Repository
2. **Use Templ for all templates**: Never write raw HTML in handlers
3. **Use HTMX for interactivity**: Return HTML fragments, not JSON
4. **Use Squirrel for queries**: Type-safe SQL building with Dollar placeholders
5. **Use structured logging**: slog with context fields
6. **Let Air handle builds**: It orchestrates Swagger, Templ, Tailwind, and Go
7. **Follow naming conventions**: Files, packages, and functions use snake_case
8. **Test before deploying**: Though test coverage is currently low
9. **Check security**: Review CORS, CSRF, authentication for each change
10. **Document APIs**: Add Swagger comments for all API endpoints

This codebase is a **modern, well-structured Go web application** that effectively combines server-side rendering with HTMX for a seamless developer and user experience. The clean architecture makes it easy to extend and maintain.
