# 🎬 映画・TV番組API

Go言語とTMDb APIを活用した映画・TV番組情報提供APIです。作品検索、詳細情報、キャスト情報、レビュー、トレンド情報など包括的な機能を提供します。

![Go](https://img.shields.io/badge/Go-1.24+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![React](https://img.shields.io/badge/React-19-61DAFB?style=for-the-badge&logo=react&logoColor=black)
![TypeScript](https://img.shields.io/badge/TypeScript-5.0+-3178C6?style=for-the-badge&logo=typescript&logoColor=white)
![Tailwind](https://img.shields.io/badge/Tailwind-4.0-06B6D4?style=for-the-badge&logo=tailwindcss&logoColor=white)
![MUI](https://img.shields.io/badge/MUI-5.0-007FFF?style=for-the-badge&logo=mui&logoColor=white)
![Redux](https://img.shields.io/badge/Redux-5.0-764ABC?style=for-the-badge&logo=redux&logoColor=white)
![TMDb](https://img.shields.io/badge/TMDb-API-01b4e4?style=for-the-badge&logo=themoviedatabase&logoColor=white)
![Vercel](https://img.shields.io/badge/Vercel-000000?style=for-the-badge&logo=vercel&logoColor=white)
![Fly.io](https://img.shields.io/badge/Fly.io-7C4DFF?style=for-the-badge&logo=fly&logoColor=white)

## 🌟 主な機能

- **🔍 検索機能**: 映画・TV番組のキーワード検索、フィルタリング、ページネーション
- **📋 詳細情報**: 作品の包括的な詳細データ（あらすじ、評価、制作情報等）
- **⭐ レビュー・評価**: TMDbレビューと独自レビュー投稿機能
- **🎭 キャスト・スタッフ**: 出演者・制作陣の詳細情報と出演作品履歴
- **🔥 トレンド・おすすめ**: 人気作品、類似作品、パーソナライズ推薦
- **💻 モダンUI**: React19ベースのレスポンシブWebアプリケーション

## 🚀 クイックスタート

### 必要な環境

- **Docker & Docker Compose**: 必須（全サービス用）
- **TMDb API**: Read Access Token（無料取得可能）

### ローカル開発環境セットアップ

```bash
# 1. リポジトリのクローン
git clone https://github.com/takeshi-arihori/movie-api.git
cd movie-api

# 2. 環境変数の設定
cp .env.example .env
# .envファイルにTMDb APIキーを設定

# 3. 全サービス起動（フロント + バック + DB）
docker compose up -d

# 4. アプリケーションアクセス
# フロントエンド: http://localhost:3005
# バックエンドAPI: http://localhost:8802
```

### 簡単起動（推奨）

```bash
# 開発環境一括起動
make dev

# ログ確認しながら起動
docker compose up

# 特定サービスのみ起動
docker compose up frontend backend
```

## 🌐 デプロイメント

### 本番環境構成
- **フロントエンド**: [Vercel](https://vercel.com/) でホスティング
- **バックエンドAPI**: [Fly.io](https://fly.io/) でホスティング
- **データベース**: Fly.io PostgreSQL または外部サービス

### デプロイ手順

#### 🚀 バックエンド (Fly.io)

```bash
# 1. Fly CLI インストール
curl -L https://fly.io/install.sh | sh

# 2. ログイン
fly auth login

# 3. アプリケーション作成
fly apps create movie-api-backend

# 4. 設定ファイル生成
fly launch

# 5. 環境変数設定
fly secrets set TMDB_API_KEY=your_api_key_here
fly secrets set DATABASE_URL=your_database_url

# 6. デプロイ
fly deploy
```

#### ⚡ フロントエンド (Vercel)

```bash
# 1. Vercel CLI インストール
npm i -g vercel

# 2. プロジェクトリンク
cd frontend
vercel link

# 3. 環境変数設定
vercel env add VITE_API_BASE_URL

# 4. デプロイ
vercel --prod

# または GitHub連携で自動デプロイ
```

## 📖 API仕様

### 🔍 検索系エンドポイント

| エンドポイント     | メソッド | 説明                   |
| ------------------ | -------- | ---------------------- |
| `/api/v1/search`   | GET      | 映画・TV番組の統合検索 |
| `/api/v1/discover` | GET      | 条件指定による作品探索 |

### 📋 詳細情報系エンドポイント

| エンドポイント                | メソッド | 説明                         |
| ----------------------------- | -------- | ---------------------------- |
| `/api/v1/movies/{id}`         | GET      | 映画の詳細情報               |
| `/api/v1/tv/{id}`             | GET      | TV番組の詳細情報             |
| `/api/v1/movies/{id}/credits` | GET      | 映画のキャスト・スタッフ情報 |
| `/api/v1/person/{id}`         | GET      | 人物の詳細情報               |

### ⭐ レビュー・評価系エンドポイント

| エンドポイント                | メソッド | 説明                   |
| ----------------------------- | -------- | ---------------------- |
| `/api/v1/movies/{id}/reviews` | GET      | 映画のレビュー一覧取得 |
| `/api/v1/movies/{id}/reviews` | POST     | 映画のレビュー投稿     |
| `/api/v1/movies/{id}/rating`  | GET      | 映画の評価統計         |

### 🔥 トレンド・おすすめ系エンドポイント

| エンドポイント                | メソッド | 説明                 |
| ----------------------------- | -------- | -------------------- |
| `/api/v1/trending`            | GET      | トレンド作品一覧     |
| `/api/v1/movies/{id}/similar` | GET      | 類似映画の推薦       |
| `/api/v1/popular`             | GET      | 人気作品ランキング   |
| `/api/v1/top-rated`           | GET      | 高評価作品ランキング |

### 🏥 システム系エンドポイント

| エンドポイント   | メソッド | 説明           |
| ---------------- | -------- | -------------- |
| `/api/v1/health` | GET      | ヘルスチェック |

> 📚 詳細な仕様は [API仕様書](docs/api/README.md) をご覧ください。

## 🛠️ 技術スタック

### バックエンド
- **言語**: Go 1.24
- **フレームワーク**: net/http, gorilla/mux
- **外部API**: The Movie Database (TMDb) API v3
- **データベース**: PostgreSQL
- **キャッシュ**: Redis
- **開発環境**: Docker & Docker Compose
- **デプロイ**: Fly.io

### フロントエンド
- **言語**: TypeScript 5.0+
- **フレームワーク**: React 19
- **スタイリング**: Tailwind CSS 4.0
- **UIライブラリ**: Material-UI (MUI) 5.0
- **状態管理**: Redux Toolkit + RTK Query
- **ルーティング**: React Router v6
- **フォーム**: React Hook Form + Zod
- **ビルドツール**: Vite
- **デプロイ**: Vercel

### 開発・運用
- **コンテナ**: Docker & Docker Compose（フルスタック）
- **CI/CD**: GitHub Actions
- **モニタリング**: Fly.io metrics + Vercel Analytics
- **バージョン管理**: Git + GitHub

## 📁 プロジェクト構成

```
movie-api/
├── 📄 compose.yaml             # Docker Compose設定（全サービス）
├── 📄 Makefile                 # ビルド・開発用コマンド
├── 📄 .env.example             # 環境変数設定例
├── 📄 .gitignore               # Git除外設定
├── 📄 README.md                # プロジェクト概要
├── 📄 CLAUDE.md                # Claude Code開発ガイド
│
├── 📂 backend/                 # Go バックエンドAPI
│   ├── 📄 main.go              # Go アプリケーションエントリーポイント
│   ├── 📄 go.mod               # Go モジュール定義
│   ├── 📄 go.sum               # Go 依存関係ロック
│   ├── 📄 Dockerfile           # Go バックエンド用
│   ├── 📄 fly.toml             # Fly.io 設定
│   │
│   └── 📂 internal/            # Go 内部パッケージ
│       ├── 📂 config/          # 設定管理
│       ├── 📂 models/          # データモデル
│       ├── 📂 services/        # ビジネスロジック
│       ├── 📂 handlers/        # HTTPハンドラー
│       ├── 📂 middleware/      # ミドルウェア
│       └── 📂 utils/           # ユーティリティ
│
├── 📂 frontend/                # React フロントエンド
│   ├── 📄 package.json         # npm 依存関係
│   ├── 📄 Dockerfile           # React フロントエンド用
│   ├── 📄 vite.config.ts       # Vite設定
│   ├── 📄 vercel.json          # Vercel設定
│   ├── 📄 tailwind.config.js   # Tailwind CSS設定
│   ├── 📄 tsconfig.json        # TypeScript設定
│   │
│   ├── 📂 public/              # 静的ファイル
│   └── 📂 src/                 # ソースコード
│       ├── 📄 main.tsx         # React エントリーポイント
│       ├── 📂 app/             # アプリケーションレイヤー
│       ├── 📂 features/        # 機能別コード
│       ├── 📂 components/      # 共通コンポーネント
│       ├── 📂 hooks/           # カスタムフック
│       ├── 📂 stores/          # Redux状態管理
│       ├── 📂 types/           # TypeScript型定義
│       └── 📂 utils/           # ユーティリティ
│
├── 📂 docs/                    # ドキュメント
│   ├── 📂 api/                 # API仕様書
│   ├── 📄 deployment.md        # デプロイガイド
│   └── 📄 development.md       # 開発ガイド
│
└── 📂 scripts/                 # 自動化スクリプト
    ├── 📄 setup.sh             # 初期セットアップ
    ├── 📄 build.sh             # ビルド
    └── 📄 deploy.sh            # デプロイ
```

## 🐳 Docker構成

### フルスタック Docker構成

```yaml
# compose.yaml
services:
  # React フロントエンド
  frontend:
    build:
      context: ./frontend
      dockerfile: Dockerfile
    ports:
      - "3005:3000"
    environment:
      - VITE_API_BASE_URL=http://localhost:8802/api/v1
    volumes:
      - ./frontend:/app
      - /app/node_modules
    depends_on:
      - backend

  # Go バックエンドAPI
  backend:
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "8802:8080"
    depends_on:
      - postgres
      - redis
    environment:
      - PORT=8080
      - DATABASE_URL=postgres://developer:password@postgres:5432/movieapi
      - REDIS_URL=redis://redis:6379
      - TMDB_API_KEY=${TMDB_API_KEY}
      - CORS_ORIGINS=http://localhost:3005
    volumes:
      - .:/app
      - /app/tmp
    develop:
      watch:
        - action: rebuild
          path: .
          ignore:
            - frontend/
            - docs/
  
  # PostgreSQL データベース
  postgres:
    image: postgres:15-alpine
    environment:
      POSTGRES_DB: movieapi
      POSTGRES_USER: developer
      POSTGRES_PASSWORD: password
    ports:
      - "5435:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U developer -d movieapi"]
      interval: 30s
      timeout: 10s
      retries: 3
  
  # Redis キャッシュ
  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data
    healthcheck:
      test: ["CMD", "redis-cli", "ping"]
      interval: 30s
      timeout: 10s
      retries: 3

  # Adminer (データベース管理ツール)
  adminer:
    image: adminer:latest
    ports:
      - "8080:8080"
    depends_on:
      - postgres

volumes:
  postgres_data:
  redis_data:
```

### フロントエンド Dockerfile

```dockerfile
# frontend/Dockerfile
FROM node:20-alpine

WORKDIR /app

# 依存関係ファイルをコピーしてインストール
COPY package*.json ./
RUN npm ci

# ソースコードをコピー
COPY . .

# 開発サーバー起動
EXPOSE 3000
CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]
```

### バックエンド Dockerfile

```dockerfile
# Dockerfile (ルート)
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/main .
EXPOSE 8080

CMD ["./main"]
```

## 🧪 テスト

### バックエンドテスト
```bash
# Dockerコンテナ内でテスト実行
docker compose exec backend go test ./...

# カバレッジ付きテスト
docker compose exec backend go test -cover ./...

# 統合テスト
docker compose exec backend go test -tags=integration ./...

# ローカル環境でのテスト（Dockerなし）
go test ./...
```

### フロントエンドテスト
```bash
# Dockerコンテナ内でテスト実行
docker compose exec frontend npm test

# E2Eテスト
docker compose exec frontend npm run test:e2e

# カバレッジ確認
docker compose exec frontend npm run test:coverage

# ローカル環境でのテスト
cd frontend && npm test
```

## 🔧 開発コマンド

### Makefileコマンド
```bash
# 開発環境起動（全サービス）
make dev

# 全サービス起動
make up

# 全サービス停止
make down

# テスト実行
make test

# ビルド
make build

# ログ確認
make logs

# 環境リセット
make reset

# 本番デプロイ
make deploy
```

### Docker Composeコマンド
```bash
# 全サービス起動
docker compose up -d

# ログ確認しながら起動
docker compose up

# 特定サービスのみ起動
docker compose up frontend backend
docker compose up postgres redis

# サービス停止
docker compose down

# ボリューム含めて削除
docker compose down -v

# ログ確認
docker compose logs frontend
docker compose logs backend
docker compose logs -f  # リアルタイム

# コンテナ内に入る
docker compose exec frontend sh
docker compose exec backend sh
docker compose exec postgres psql -U developer -d movieapi

# サービス再起動
docker compose restart backend
docker compose restart frontend
```

### 開発用ショートカット
```bash
# フロントエンド開発（ホットリロード）
docker compose up frontend

# バックエンド開発（ホットリロード）
docker compose up backend

# データベースのみ起動
docker compose up postgres

# データベース接続
docker compose exec postgres psql -U developer -d movieapi

# Redis接続
docker compose exec redis redis-cli
```

## 🌍 環境変数

### ルート環境変数 (.env)
```bash
# TMDb API
TMDB_API_KEY=your_tmdb_api_key_here

# データベース
POSTGRES_DB=movieapi
POSTGRES_USER=developer
POSTGRES_PASSWORD=password

# セキュリティ
JWT_SECRET=your_jwt_secret_here

# 環境設定
ENV=development
```

### バックエンド環境変数（Docker内）
```bash
# サーバー設定
PORT=8080

# TMDb API
TMDB_API_KEY=${TMDB_API_KEY}
TMDB_BASE_URL=https://api.themoviedb.org/3

# データベース（Docker内部通信）
DATABASE_URL=postgres://developer:password@postgres:5432/movieapi

# Redis（Docker内部通信）
REDIS_URL=redis://redis:6379

# CORS
CORS_ORIGINS=http://localhost:3005
```

### フロントエンド環境変数（Docker内）
```bash
# API設定（Docker内部通信）
VITE_API_BASE_URL=http://localhost:8802/api/v1

# 本番用
VITE_API_BASE_URL=https://your-api-domain.fly.dev/api/v1
```

## 📈 パフォーマンス最適化

### バックエンド
- **Redis キャッシュ**: TMDb APIレスポンスのキャッシュ
- **接続プール**: データベース接続の効率化
- **レート制限**: API使用量の制御
- **圧縮**: gzip圧縮でレスポンスサイズ削減

### フロントエンド
- **コード分割**: React.lazy でルート分割
- **画像最適化**: Next.js Image最適化（将来）
- **Bundle最適化**: Vite の自動最適化
- **CDN配信**: Vercel Edge Network活用

## 🚀 CI/CD パイプライン

### GitHub Actions
```yaml
# .github/workflows/deploy.yml
name: Deploy
on:
  push:
    branches: [main]

jobs:
  backend:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: superfly/flyctl-actions/setup-flyctl@master
      - run: flyctl deploy --remote-only

  frontend:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - run: cd frontend && npm install && npm run build
      # Vercelは自動デプロイ
```

## 🤝 コントリビューション

### 開発フロー
1. **Issues作成**: 機能追加・バグ修正の提案
2. **ブランチ作成**: `feature/機能名` または `fix/修正内容`
3. **実装・テスト**: 単体テスト・E2Eテストの実行
4. **プルリクエスト**: コードレビュー依頼
5. **マージ**: レビュー完了後に本番反映

### コミット規約
```bash
# 新機能
feat: ✨ 映画検索APIエンドポイント追加

# バグ修正
fix: 🐛 検索結果のページネーション修正

# ドキュメント
docs: 📝 API仕様書更新

# リファクタリング
refactor: ♻️ コンポーネント構造改善

# パフォーマンス
perf: ⚡ キャッシュ機能追加

# デプロイ
deploy: 🚀 本番環境デプロイ
```

## 📞 サポート・お問い合わせ

- **開発者**: [@takeshi-arihori](https://github.com/takeshi-arihori)
- **Issues**: [GitHub Issues](https://github.com/takeshi-arihori/movie-api/issues)
- **API仕様**: [Swagger/OpenAPI](https://localhost:8802/swagger)
- **ライブデモ**: [https://your-app.vercel.app](https://your-app.vercel.app)
- **データベース管理**: [http://localhost:8080](http://localhost:8080) (Adminer)

## 🙏 謝辞

- **[The Movie Database (TMDb)](https://www.themoviedb.org/)**: 豊富な映画・TV番組データの提供
- **[Vercel](https://vercel.com/)**: 優れたフロントエンドホスティング
- **[Fly.io](https://fly.io/)**: 高速なバックエンドデプロイメント
- **[bulletproof-react](https://github.com/alan2207/bulletproof-react)**: Reactベストプラクティスの指針
- **オープンソースコミュニティ**: 多くのライブラリとツールの提供

---

**📌 注意**: このプロジェクトは学習・ポートフォリオ目的で作成されています。TMDb APIの利用規約を遵守してご使用ください。
