.PHONY: dev up down logs test build clean reset deploy help

# Default target
help:
	@echo "Available commands:"
	@echo "  make dev     - Start development environment"
	@echo "  make up      - Start all services"
	@echo "  make down    - Stop all services"
	@echo "  make logs    - Show logs for all services"
	@echo "  make test    - Run tests"
	@echo "  make build   - Build all services"
	@echo "  make clean   - Clean up containers and volumes"
	@echo "  make reset   - Reset entire development environment"
	@echo "  make deploy  - Deploy to production"

# Development environment
dev: up
	@echo "Development environment started!"
	@echo "Frontend: http://localhost:3005"
	@echo "Backend API: http://localhost:8802"
	@echo "Database Admin: http://localhost:8080"

# Start all services
up:
	docker compose up -d

# Stop all services
down:
	docker compose down

# Show logs
logs:
	docker compose logs -f

# Run tests
test:
	docker compose exec backend go test ./...
	@if [ -d "frontend" ] && [ -f "frontend/package.json" ]; then \
		docker compose exec frontend npm test; \
	fi

# Build all services
build:
	docker compose build

# Clean up
clean:
	docker compose down -v
	docker system prune -f

# Reset development environment
reset: clean
	docker compose build --no-cache
	docker compose up -d

# Deploy (placeholder)
deploy:
	@echo "Deploying to production..."
	@echo "Backend: fly deploy --config backend/fly.toml"
	@echo "Frontend: cd frontend && vercel --prod"