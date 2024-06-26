DB_HOST = localhost
DB_USER = postgres
DB_NAME = postgres
DB_PASSWORD = postgres
DB_PORT = 5432

dev:
	@echo "Running locally..."
	@nodemon --exec DB_PORT=$(DB_PORT) DB_HOST=$(DB_HOST) DB_USER=$(DB_USER) DB_NAME=$(DB_NAME) DB_PASSWORD=$(DB_PASSWORD) go run main.go --signal SIGTERM
up:
	@echo "Setting up containeres..."
	@docker compose down 
	@docker compose up
generate:
	@echo "Generating..."
	@go generate ./...
test:
	@go run github.com/onsi/ginkgo/v2/ginkgo -v -p --race ./tests --output-interceptor-mode=none
