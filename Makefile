all: build run

build:
	go build server.go

run:
	go run server.go

clean:
	go mod tidy