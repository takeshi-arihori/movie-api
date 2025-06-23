# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**Movie & TV Shows API** - Full-stack application combining Go backend API with React 19 frontend for modern movie and TV show information service.

### Goal
Build a modern movie and TV show information application combining Go backend API with React 19 frontend.

### Tech Stack

#### Backend
- **Go 1.24** - Main programming language
- **net/http + gorilla/mux** - HTTP server and routing
- **TMDb API v3** - Movie and TV show data source
- **JSON** - Data format

#### Frontend
- **React 19** - UI framework
- **TypeScript 5.0+** - Type-safe development
- **Tailwind CSS 4.0** - Utility-first CSS
- **Material-UI (MUI) 5.0** - UI component library
- **Redux Toolkit + RTK Query** - State management and API calls
- **Vite** - Build tool
- **React Router v6** - Routing

## Architecture Design

### Design Principles
- **Clean Architecture** - Layer separation for maintainability
- **bulletproof-react** - React best practices
- **Feature-based structure** - Scalable directory organization  
- **Type safety** - TypeScript strict mode development

### Directory Structure

```
movie-api/
├── backend/              # Go backend API
│   ├── main.go          # Application entry point
│   ├── go.mod           # Go module definition
│   ├── Dockerfile       # Backend container config
│   └── internal/        # Internal packages
│       ├── config/      # Configuration management
│       ├── models/      # Data models
│       ├── services/    # Business logic (TMDb client)
│       ├── handlers/    # HTTP handlers
│       ├── middleware/  # HTTP middleware
│       └── utils/       # Utilities
├── frontend/            # React frontend
│   ├── Dockerfile       # Frontend container config
│   └── src/
│       ├── app/         # Application layer (store, router)
│       ├── features/    # Feature-based modules
│       ├── components/  # Shared components
│       ├── hooks/       # Custom hooks
│       ├── stores/      # Redux state management
│       ├── types/       # TypeScript definitions
│       └── utils/       # Utilities
├── compose.yaml         # Docker Compose configuration
├── Makefile            # Development commands
└── .env.example        # Environment variables template
```

## Development Commands

### Project Setup
```bash
# Clone and setup environment
git clone https://github.com/takeshi-arihori/movie-api.git
cd movie-api
cp .env.example .env
# Edit .env with your TMDb API key
```

### Docker Development Environment
```bash
# Start all services (recommended)
make dev
# or
docker compose up -d

# View logs
make logs
# or
docker compose logs -f

# Stop all services
make down
# or
docker compose down
```

### Individual Service Management
```bash
# Start specific services
docker compose up backend postgres redis
docker compose up frontend

# Restart services
docker compose restart backend
docker compose restart frontend

# Access containers
docker compose exec backend sh
docker compose exec frontend sh
docker compose exec postgres psql -U developer -d movieapi
```

### Testing
```bash
# Backend tests (in container)
docker compose exec backend go test ./...
docker compose exec backend go test -cover ./...

# Frontend tests (in container)
docker compose exec frontend npm test
docker compose exec frontend npm run test:e2e

# Local testing (if Go/Node installed)
cd backend && go test ./...
cd frontend && npm test
```

### Build & Deploy
```bash
# Build all services
make build
# or
docker compose build

# Deploy to production
make deploy
# Deploys backend to Fly.io and frontend to Vercel
```

## Environment Configuration

### Root Environment Variables (.env)
```bash
# TMDb API (required)
TMDB_API_KEY=your_tmdb_api_key_here

# Database credentials
POSTGRES_DB=movieapi
POSTGRES_USER=developer
POSTGRES_PASSWORD=password

# Security
JWT_SECRET=your_jwt_secret_here

# Environment
ENV=development
```

### Docker Service URLs
```bash
# Frontend: http://localhost:3005
# Backend API: http://localhost:8802/api/v1
# Database Admin: http://localhost:8080
# PostgreSQL: localhost:5435
# Redis: localhost:6379
```

## Key API Endpoints

### Search & Discovery
- `GET /api/v1/search` - Multi-search movies/TV shows
- `GET /api/v1/discover` - Discover content with filters

### Content Details
- `GET /api/v1/movies/{id}` - Movie details
- `GET /api/v1/tv/{id}` - TV show details
- `GET /api/v1/movies/{id}/credits` - Cast & crew
- `GET /api/v1/person/{id}` - Person details

### Reviews & Ratings
- `GET /api/v1/movies/{id}/reviews` - Get reviews
- `POST /api/v1/movies/{id}/reviews` - Submit review

### Trending & Recommendations
- `GET /api/v1/trending` - Trending content
- `GET /api/v1/movies/{id}/similar` - Similar movies
- `GET /api/v1/popular` - Popular content
- `GET /api/v1/top-rated` - Top rated content

## Implementation Priority

### Phase 1: Backend Core (High Priority)
- Configuration management (`backend/internal/config/`)
- TMDb API client (`backend/internal/services/tmdb_client.go`)
- Data models (`backend/internal/models/`)
- Search endpoint (`/api/v1/search`)
- Detail endpoints (`/api/v1/movies/{id}`, `/api/v1/tv/{id}`)

### Phase 2: Frontend Foundation (Medium Priority)
- React environment setup (`frontend/` with Vite + TypeScript)
- Redux store configuration (`frontend/src/app/store.ts`)
- API client (`frontend/src/lib/api.ts`)
- Layout components (`frontend/src/components/layout/`)
- Search feature (`frontend/src/features/search/`)

### Phase 3: Extended Features (Low Priority)
- Review/rating system (`/api/v1/reviews`)
- Cast/crew information (`/api/v1/credits`)
- Trending/recommendation features (`/api/v1/trending`)
- User authentication (future)

## 🎨 コーディング規約

### Go (バックエンド)
```go
// パッケージ命名: 小文字、単語
package handlers

// 構造体: PascalCase
type MovieHandler struct {}

// 関数: PascalCase (public), camelCase (private)
func (h *MovieHandler) GetMovieDetails() {}
func (h *MovieHandler) validateRequest() {}

// エラーハンドリング
if err != nil {
    return fmt.Errorf("failed to get movie: %w", err)
}
```

### TypeScript (フロントエンド)
```typescript
// ファイル名: kebab-case
// movie-search.tsx

// コンポーネント: PascalCase
export const MovieSearch: React.FC = () => {}

// フック: use + PascalCase
export const useMovieSearch = () => {}

// 型定義: PascalCase
export interface MovieSearchProps {}

// 定数: UPPER_SNAKE_CASE
export const API_ENDPOINTS = {}
```




## Common Issues & Troubleshooting

1. **CORS errors**: Check `internal/middleware/cors.go` configuration
2. **API key errors**: Verify TMDb API key in `.env` file
3. **Port conflicts**: Ensure ports 8080 (backend) and 3000 (frontend) are available
4. **Log monitoring**: Use `tail -f /var/log/movie-api.log` for backend logs

## Performance Considerations

### Backend
- Implement caching (Redis or in-memory) for TMDb API responses
- Handle TMDb API rate limiting
- Consider PostgreSQL for future data persistence

### Frontend
- Use React.lazy for route-based code splitting
- Optimize images (WebP format, lazy loading)
- Monitor bundle size with Vite Bundle Analyzer

## Git Workflow

- Main branch: `main` (always deployable)
- Feature branches: `feature/issue-number-description`
- Bug fixes: `fix/issue-number-description`
- Commit format: `type: emoji #issue description` (e.g., `feat: ✨ #15 映画検索エンドポイントを実装`)
- One feature/issue per branch and PR

## Language & Documentation

### CLAUDE.md Language
- **English is recommended** for CLAUDE.md to ensure compatibility with Claude Code instances globally
- Japanese project documentation should be maintained in separate files (README.md, README.github.md)
- Code comments and commit messages can be in Japanese as per project team preference
- This approach balances international accessibility with local team communication needs
