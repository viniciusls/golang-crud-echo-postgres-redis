all: build run

build:
	go build main.go

run:
	go run main.go

clean:
	go mod tidy