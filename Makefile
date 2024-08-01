clean:
	rm -r bin

build: gen
	go build -o bin/htemplx main.go

serve: build
	./bin/htemplx serve

gen:
	templ generate
