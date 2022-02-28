compose-up:
	@docker-compose up -d
compose-down:
	@docker-compose down
b-add: # Build Add Binary
	@go build -tags add -o ./commander ./cmd
b-list: # Build List Binary
	@go build -tags=list -o ./commander ./cmd