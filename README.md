# 🎬 映画・TV番組API

Go言語とTMDb APIを活用した映画・TV番組情報提供APIです。作品検索、詳細情報、キャスト情報、レビュー、トレンド情報など包括的な機能を提供します。

![Go](https://img.shields.io/badge/Go-1.24+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![React](https://img.shields.io/badge/React-19-61DAFB?style=for-the-badge&logo=react&logoColor=black)
![TypeScript](https://img.shields.io/badge/TypeScript-5.0+-3178C6?style=for-the-badge&logo=typescript&logoColor=white)
![Tailwind](https://img.shields.io/badge/Tailwind-4.0-06B6D4?style=for-the-badge&logo=tailwindcss&logoColor=white)
![MUI](https://img.shields.io/badge/MUI-5.0-007FFF?style=for-the-badge&logo=mui&logoColor=white)
![Redux](https://img.shields.io/badge/Redux-5.0-764ABC?style=for-the-badge&logo=redux&logoColor=white)
![TMDb](https://img.shields.io/badge/TMDb-API-01b4e4?style=for-the-badge&logo=themoviedatabase&logoColor=white)

## 🌟 主な機能

- **🔍 検索機能**: 映画・TV番組のキーワード検索、フィルタリング、ページネーション
- **📋 詳細情報**: 作品の包括的な詳細データ（あらすじ、評価、制作情報等）
- **⭐ レビュー・評価**: TMDbレビューと独自レビュー投稿機能
- **🎭 キャスト・スタッフ**: 出演者・制作陣の詳細情報と出演作品履歴
- **🔥 トレンド・おすすめ**: 人気作品、類似作品、パーソナライズ推薦
- **💻 デモアプリ**: React19ベースのモダンWebアプリケーション

## 🚀 クイックスタート

### 必要な環境

- **Go**: 1.24
- **Node.js**: 20.0以上
- **TMDb API**: Read Access Token（無料取得可能）
- **Git**: バージョン管理

### インストール

```bash
# リポジトリのクローン
git clone https://github.com/takeshi-arihori/movie-api.git
cd movie-api

# バックエンド依存関係のインストール
go mod download

# フロントエンド依存関係のインストール
cd frontend
npm install

# 環境変数の設定
cp .env.example .env
# .envファイルにTMDb APIキーを設定してください

# サーバー起動（開発環境）
cd ..
make dev

# または個別に起動
# バックエンド
go run main.go

# フロントエンド（別ターミナル）
cd frontend
npm run dev
```

### 基本的な使用例

```bash
# ヘルスチェック
curl http://localhost:8080/api/v1/health

# 映画検索
curl "http://localhost:8080/api/v1/search?query=君の名は&type=movie"

# 映画詳細取得
curl http://localhost:8080/api/v1/movies/372058

# トレンド作品取得
curl "http://localhost:8080/api/v1/trending?media_type=movie&time_window=week"
```

### Webアプリケーション

フロントエンド開発サーバー起動後、ブラウザで `http://localhost:3000` を開いてModern Reactアプリをお試しください。

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

### 🏥 システム系
