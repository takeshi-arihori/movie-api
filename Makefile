# Movie API Makefile
# Provides convenient commands for development, testing, and debugging

.PHONY: help dev debug build test clean logs status

# Default target
help: ## Show this help message
	@echo "Movie API Development Commands"
	@echo "============================="
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development
dev: ## Start development environment (production mode)
	@echo "🚀 Starting development environment..."
	docker compose up -d postgres redis
	@echo "⏳ Waiting for services to be ready..."
	sleep 5
	docker compose up backend frontend

dev-detached: ## Start development environment in background
	@echo "🚀 Starting development environment in detached mode..."
	docker compose up -d

##@ Debugging
debug: ## Start debug environment with Delve debugger
	@echo "🐛 Starting debug environment with Delve..."
	docker compose --profile debug up -d postgres redis
	@echo "⏳ Waiting for services to be ready..."
	sleep 5
	docker compose --profile debug up backend-debug
	@echo "📍 Delve debugger available at localhost:2345"
	@echo "📊 pprof available at http://localhost:6060/debug/pprof/"

debug-detached: ## Start debug environment in background
	@echo "🐛 Starting debug environment in detached mode..."
	docker compose --profile debug up -d
	@echo "📍 Delve debugger available at localhost:2345"
	@echo "📊 pprof available at http://localhost:6060/debug/pprof/"

debug-attach: ## Attach to running debug container
	@echo "🔗 Attaching to debug container..."
	docker compose --profile debug exec backend-debug /bin/sh

##@ Building
build: ## Build all services
	@echo "🔨 Building all services..."
	docker compose build

build-backend: ## Build backend service only
	@echo "🔨 Building backend..."
	docker compose build backend

build-frontend: ## Build frontend service only
	@echo "🔨 Building frontend..."
	docker compose build frontend

build-debug: ## Build debug version of backend
	@echo "🔨 Building debug backend..."
	docker compose --profile debug build backend-debug

##@ Testing
test: ## Run all tests
	@echo "🧪 Running backend tests..."
	cd backend && go test ./...
	@echo "🧪 Running frontend tests..."
	cd frontend && npm test

test-backend: ## Run backend tests only
	@echo "🧪 Running backend tests..."
	cd backend && go test ./...

test-backend-coverage: ## Run backend tests with coverage
	@echo "🧪 Running backend tests with coverage..."
	cd backend && go test -cover ./...

test-frontend: ## Run frontend tests only
	@echo "🧪 Running frontend tests..."
	cd frontend && npm test

##@ Database
db-tools: ## Start database tools (pgAdmin)
	@echo "🗄️  Starting database tools..."
	docker compose --profile tools up -d pgadmin
	@echo "🌐 pgAdmin available at http://localhost:8081"

db-migrate: ## Run database migrations (placeholder)
	@echo "📋 Running database migrations..."
	@echo "⚠️  Migrations not implemented yet"

db-seed: ## Seed database with sample data (placeholder)
	@echo "🌱 Seeding database..."
	@echo "⚠️  Seeding not implemented yet"

##@ Monitoring
logs: ## View logs from all services
	docker compose logs -f

logs-backend: ## View backend logs
	docker compose logs -f backend

logs-frontend: ## View frontend logs
	docker compose logs -f frontend

logs-debug: ## View debug backend logs
	docker compose --profile debug logs -f backend-debug

status: ## Show status of all services
	@echo "📊 Service Status:"
	docker compose ps

health: ## Check health of all services
	@echo "🏥 Health Check:"
	@echo "Backend: $$(curl -s http://localhost:8080/health | jq -r '.status // "unhealthy"' 2>/dev/null || echo "unreachable")"
	@echo "Frontend: $$(curl -s http://localhost:3000/health 2>/dev/null | head -n1 || echo "unreachable")"
	@echo "Postgres: $$(docker compose exec postgres pg_isready -U movieapi 2>/dev/null || echo "unreachable")"
	@echo "Redis: $$(docker compose exec redis redis-cli ping 2>/dev/null || echo "unreachable")"

##@ Cleanup
clean: ## Stop and remove all containers, networks, and volumes
	@echo "🧹 Cleaning up..."
	docker compose --profile debug --profile tools down -v
	docker compose down -v
	docker system prune -f

clean-images: ## Remove all built images
	@echo "🧹 Removing built images..."
	docker compose --profile debug --profile tools down --rmi all
	docker compose down --rmi all

clean-volumes: ## Remove all volumes (⚠️  This will delete all data!)
	@echo "⚠️  This will delete all database data!"
	@read -p "Are you sure? [y/N] " -n 1 -r; \
	if [[ $$REPLY =~ ^[Yy]$$ ]]; then \
		echo "🧹 Removing all volumes..."; \
		docker compose down -v; \
		docker volume prune -f; \
	else \
		echo "Cancelled."; \
	fi

##@ Local Development
local-backend: ## Run backend locally (requires Go)
	@echo "🏃 Running backend locally..."
	cd backend && go run main.go

local-frontend: ## Run frontend locally (requires Node.js)
	@echo "🏃 Running frontend locally..."
	cd frontend && npm run dev

local-install: ## Install local dependencies
	@echo "📦 Installing Go dependencies..."
	cd backend && go mod tidy
	@echo "📦 Installing Node.js dependencies..."
	cd frontend && npm install

##@ Security
security-scan: ## Run security scan (placeholder)
	@echo "🔒 Running security scan..."
	@echo "⚠️  Security scanning not implemented yet"

##@ Utilities
reset: clean build dev-detached ## Reset environment (clean, build, start)

format: ## Format code
	@echo "🎨 Formatting Go code..."
	cd backend && gofmt -w .
	@echo "🎨 Formatting TypeScript code..."
	cd frontend && npm run format

lint: ## Run linters
	@echo "🔍 Linting Go code..."
	cd backend && golangci-lint run
	@echo "🔍 Linting TypeScript code..."
	cd frontend && npm run lint