clean:
	rm -rf bin tmp

build: templ-generate tailwind-gen
	go build -o bin/htemplx main.go

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