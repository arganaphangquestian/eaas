compose-up:
	@docker-compose up -d
compose-down:
	@docker-compose down
run:
	@go run cmd/main.go