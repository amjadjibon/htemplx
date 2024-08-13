# Your Makefile rules
all:
	@echo "This is the main Makefile."

clean:
	rm -rf bin tmp

build:
	CGO_ENABLED=0 GOOS=linux go build -o bin/htemplx main.go

serve: build
	./bin/htemplx serve

DB_URL=postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable
migrate-up:
	go run main.go migrate up --db_url=${DB_URL}

migrate-down:
	go run main.go migrate down --db_url=${DB_URL}

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

# Docker Compose File
COMPOSE_FILE=distribution/compose/docker-compose.yml
COMPOSE_FILE_DEV=distribution/compose/docker-compose-dev.yml

# Build the Go binary and Docker image
docker-build:
	docker build -f ./distribution/docker/Dockerfile -t htemplx .

# Start the services
docker-up:
	docker-compose -f $(COMPOSE_FILE) up

# Start the services
docker-up-dev:
	docker-compose -f $(COMPOSE_FILE_DEV) up