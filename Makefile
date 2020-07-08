
wire:
	wire ./app/container

run:
	@go run cmd/main.go

up:
	docker-compose up --build
down:
	docker-compose down --remove-orphans
