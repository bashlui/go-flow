.PHONY: migrate migrate-status run build clean test deps

# Run migrations
migrate:
	go run cmd/migrate/main.go

# Check migration status
migrate-status:
	psql "$(DATABASE_URL)" -c "SELECT version, applied_at FROM schema_migrations ORDER BY applied_at;"

# Run the server
run:
	go run cmd/server/main.go

# Build the application
build:
	go build -o bin/server cmd/server/main.go
	go build -o bin/migrate cmd/migrate/main.go

# Clean build artifacts
clean:
	rm -rf bin/

# Run tests
test:
	go test ./...

# Install dependencies
deps:
	go mod tidy
	go mod download