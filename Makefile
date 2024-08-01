clean:
	rm -rf bin tmp

build: templ-generate
	go build -o bin/htemplx main.go

serve: build
	./bin/htemplx serve

templ-install:
	go install github.com/a-h/templ/cmd/templ@latest

templ-generate:
	templ generate

air-install:
	go install github.com/air-verse/air@latest

run:
	air -c .air.toml
