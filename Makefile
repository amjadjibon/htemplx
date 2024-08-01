clean:
	rm -r bin

build:
	go build -o bin/htemplx main.go

serve: build
	./bin/htemplx serve
