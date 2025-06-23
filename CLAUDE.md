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
- **Test-Driven Development (TDD)** - Write tests first, then implement

### Development Methodology: TDD (Test-Driven Development)

**TDD Process (Red-Green-Refactor):**
1. **Red** - Write a failing test that describes the desired functionality
2. **Green** - Write minimal code to make the test pass
3. **Refactor** - Clean up code while keeping tests passing

**TDD Guidelines:**
- Always write tests BEFORE implementing functionality
- Start with the simplest failing test case
- Write only enough code to make the test pass
- Refactor for better code quality after tests pass
- Maintain high test coverage (aim for 80%+)

**Testing Strategy:**
- **Unit Tests** - Test individual functions/methods in isolation
- **Integration Tests** - Test component interactions
- **API Tests** - Test HTTP endpoints and responses
- **Contract Tests** - Verify API contracts between services

**Test Organization:**
```
backend/
├── internal/
│   ├── config/
│   │   ├── config.go
│   │   └── config_test.go
│   ├── services/
│   │   ├── tmdb_client.go
│   │   └── tmdb_client_test.go
│   └── handlers/
│       ├── movie_handler.go
│       └── movie_handler_test.go
└── tests/
    ├── integration/     # Integration tests
    └── fixtures/        # Test data fixtures

frontend/
├── src/
│   ├── features/
│   │   └── search/
│   │       ├── components/
│   │       │   ├── SearchForm.tsx
│   │       │   └── SearchForm.test.tsx
│   │       └── hooks/
│   │           ├── useMovieSearch.ts
│   │           └── useMovieSearch.test.ts
└── tests/
    ├── __mocks__/       # Mock implementations
    └── utils/           # Test utilities
```

**Test Naming Convention:**
- Go: `TestFunctionName_Scenario_ExpectedResult`
- TypeScript: `describe('Component/Function') { it('should do something when condition') }`

**Example TDD Workflow:**
1. Create failing test for new feature
2. Run test to confirm it fails (Red)
3. Implement minimal code to pass test (Green)
4. Refactor for better design (Refactor)
5. Repeat for next requirement

**Test Requirements:**
- Every new function/method must have unit tests
- API endpoints must have integration tests
- Critical business logic must have comprehensive test coverage
- Tests should be fast, isolated, and deterministic
- Use descriptive test names that explain the behavior being tested

### Directory Structure

```
movie-api/
├── internal/             # Go backend
│   ├── config/          # Configuration management
│   ├── models/          # Data models
│   ├── services/        # Business logic (TMDb client)
│   ├── handlers/        # HTTP handlers
│   ├── middleware/      # HTTP middleware
│   └── utils/          # Utilities
├── cmd/server/         # Main application entry
└── frontend/           # React frontend
    └── src/
        ├── app/        # Application layer (store, router)
        ├── features/   # Feature-based modules
        ├── components/ # Shared components
        ├── hooks/      # Custom hooks
        ├── stores/     # Redux state management
        ├── types/      # TypeScript definitions
        └── utils/      # Utilities
```

## Development Commands

### Project Setup
```bash
# Initialize Go module
go mod init github.com/takeshi-arihori/movie-api

# Create backend structure
mkdir -p internal/{config,models,services,handlers,middleware,utils}
mkdir -p cmd/server

# Setup frontend
cd frontend
npm create vite@latest . -- --template react-ts
npm install @mui/material @emotion/react @emotion/styled @reduxjs/toolkit react-redux react-router-dom @tanstack/react-query react-hook-form @hookform/resolvers zod axios
npm install -D tailwindcss postcss autoprefixer @types/react @types/react-dom eslint @typescript-eslint/eslint-plugin prettier eslint-config-prettier
```

### Development Server
```bash
# Backend
go run main.go

# Frontend
cd frontend && npm run dev

# Both (with Makefile)
make dev
```

### Testing
```bash
# Go backend tests
go test ./...
go test -cover ./...
go test -bench=. ./...

# Frontend tests
cd frontend && npm run test
cd frontend && npm run test:e2e
cd frontend && npm run type-check
```

### Build & Deploy
```bash
# Frontend build
cd frontend && npm run build

# Backend build
go build -o movie-api main.go

# Docker (development)
docker-compose up -d
```

## Environment Configuration

### Environment Variables
```bash
# Backend (.env)
PORT=8080
TMDB_API_KEY=your_tmdb_api_key_here
TMDB_BASE_URL=https://api.themoviedb.org/3
CACHE_ENABLED=true
LOG_LEVEL=info

# Frontend (frontend/.env)
VITE_API_BASE_URL=http://localhost:8080/api/v1
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
- Configuration management (`internal/config/`)
- TMDb API client (`internal/services/tmdb_client.go`)
- Data models (`internal/models/`)
- Search endpoint (`/api/v1/search`)
- Detail endpoints (`/api/v1/movies/{id}`, `/api/v1/tv/{id}`)

### Phase 2: Frontend Foundation (Medium Priority)
- React environment setup (Vite + TypeScript)
- Redux store configuration (`src/app/store.ts`)
- API client (`src/lib/api.ts`)
- Layout components (`src/components/layout/`)
- Search feature (`src/features/search/`)

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

## Issue Memories

### GitHub Issue Tracking
- 29 fixのissueをお願い。README.github.mdの手順通りに