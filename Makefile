include: ./scripts/makefiles/docker.mk

# Your Makefile rules
all:
	@echo "This is the main Makefile."

clean:
	rm -rf bin tmp

build:
	CGO_ENABLED=0 GOOS=linux go build -o bin/htemplx main.go

serve: build
	./bin/htemplx serve

migrate-up:
	go run main.go migrate up --db_url=postgres://rootuser:rootpassword@localhost:5432/htemplx_db?sslmode=disable

migrate-down:
	go run main.go migrate down --db_url=postgres://rootuser:rootpassword@localhost:5432/htemplx_db?sslmode=disable

templ-install:
	go install github.com/a-h/templ/cmd/templ@latest

templ-generate:
	templ generate

air-install:
	go install github.com/air-verse/air@latest

run:
	air -c .air.toml

tailwind-gen:
	npx tailwindcss -i ./public/assets/css/input.css -o ./public/assets/css/output.css

# Environment Variables
export POSTGRES_USER=postgres
export POSTGRES_PASSWORD=postgres
export POSTGRES_DB=htemplx

# Docker Compose File
COMPOSE_FILE=docker-compose.yml

# Database URL
export DB_URL=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@postgres:5432/$(POSTGRES_DB)?sslmode=disable

# Build the Go binary and Docker image
docker-build:
	docker build -t htemplx .

# Start the services
docker-up:
	docker-compose -f $(COMPOSE_FILE) up