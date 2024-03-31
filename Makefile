DB_HOST = localhost
DB_USER = postgres
DB_NAME = postgres
DB_PASSWORD = postgres

dev:
	@echo "Running locally..."
	@nodemon --exec DB_HOST=$(DB_HOST) DB_USER=$(DB_USER) DB_NAME=$(DB_NAME) DB_PASSWORD=$(DB_PASSWORD) go run main.go --signal SIGTERM
up:
	@echo "Setting up containeres..."
	@docker compose down 
	@docker compose up
generate:
	@echo "Generating..."
	@go generate ./...
