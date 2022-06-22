all: build run

build:
	go build -tags dynamic main.go

run:
	go run -tags dynamic main.go

clean:
	go mod tidy