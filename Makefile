start:
	go run main.go

up:
	docker-compose up

down:
	docker-compose down

build:
	docker-compose up --build

test:
	go test -v ./...

.PHONY: start up down build