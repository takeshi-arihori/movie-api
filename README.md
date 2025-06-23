# 🎬 映画・TV番組API

Go言語とTMDb APIを活用した映画・TV番組情報提供APIです。作品検索、詳細情報、キャスト情報、レビュー、トレンド情報など包括的な機能を提供します。

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![TMDb](https://img.shields.io/badge/TMDb-API-01b4e4?style=for-the-badge&logo=themoviedatabase&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)

## 🌟 主な機能

- **🔍 検索機能**: 映画・TV番組のキーワード検索、フィルタリング、ページネーション
- **📋 詳細情報**: 作品の包括的な詳細データ（あらすじ、評価、制作情報等）
- **⭐ レビュー・評価**: TMDbレビューと独自レビュー投稿機能
- **🎭 キャスト・スタッフ**: 出演者・制作陣の詳細情報と出演作品履歴
- **🔥 トレンド・おすすめ**: 人気作品、類似作品、パーソナライズ推薦
- **💻 デモアプリ**: API機能を体験できるレスポンシブWebアプリケーション

## 🚀 クイックスタート

### 必要な環境

- **Go**: 1.21以上
- **TMDb API**: Read Access Token（無料取得可能）
- **Git**: バージョン管理

### インストール

```bash
# リポジトリのクローン
git clone https://github.com/takeshi-arihori/movie-api.git
cd movie-api

# 依存関係のインストール
go mod download

# 環境変数の設定
cp .env.example .env
# .envファイルにTMDb APIキーを設定してください

# サーバー起動
go run main.go
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

### デモアプリケーション

サーバー起動後、ブラウザで `http://localhost:8080` を開いてデモアプリをお試しください。

## 📖 API仕様

### 🔍 検索系エンドポイント

| エンドポイント | メソッド | 説明 |
|---|---|---|
| `/api/v1/search` | GET | 映画・TV番組の統合検索 |
| `/api/v1/discover` | GET | 条件指定による作品探索 |

### 📋 詳細情報系エンドポイント

| エンドポイント | メソッド | 説明 |
|---|---|---|
| `/api/v1/movies/{id}` | GET | 映画の詳細情報 |
| `/api/v1/tv/{id}` | GET | TV番組の詳細情報 |
| `/api/v1/movies/{id}/credits` | GET | 映画のキャスト・スタッフ情報 |
| `/api/v1/person/{id}` | GET | 人物の詳細情報 |

### ⭐ レビュー・評価系エンドポイント

| エンドポイント | メソッド | 説明 |
|---|---|---|
| `/api/v1/movies/{id}/reviews` | GET | 映画のレビュー一覧取得 |
| `/api/v1/movies/{id}/reviews` | POST | 映画のレビュー投稿 |
| `/api/v1/movies/{id}/rating` | GET | 映画の評価統計 |

### 🔥 トレンド・おすすめ系エンドポイント

| エンドポイント | メソッド | 説明 |
|---|---|---|
| `/api/v1/trending` | GET | トレンド作品一覧 |
| `/api/v1/movies/{id}/similar` | GET | 類似映画の推薦 |
| `/api/v1/popular` | GET | 人気作品ランキング |
| `/api/v1/top-rated` | GET | 高評価作品ランキング |

### 🏥 システム系エンドポイント

| エンドポイント | メソッド | 説明 |
|---|---|---|
| `/api/v1/health` | GET | ヘルスチェック |

> 📚 詳細な仕様は [API仕様書](docs/api/README.md) をご覧ください。

## 🛠️ 技術スタック

### バックエンド
- **言語**: Go 1.21
- **フレームワーク**: net/http, gorilla/mux
- **外部API**: The Movie Database (TMDb) API v3
- **ミドルウェア**: CORS, ログ機能, エラーハンドリング

### フロントエンド
- **基本技術**: HTML5, CSS3, JavaScript (ES6+)
- **UI Framework**: Bootstrap 5
- **アーキテクチャ**: SPA (Single Page Application)

### アーキテクチャ
- **設計パターン**: Clean Architecture, Dependency Injection
- **API設計**: RESTful API
- **データ形式**: JSON
- **キャッシュ**: インメモリキャッシュ
- **ログ**: 構造化ログ

## 📁 プロジェクト構成

```
movie-api/
├── 📄 main.go                  # アプリケーションエントリーポイント
├── 📄 go.mod                   # Go モジュール定義
├── 📄 go.sum                   # 依存関係のハッシュ
├── 📄 .env.example             # 環境変数設定例
├── 📄 README.md                # プロジェクト概要
├── 📄 README.github.md         # 開発フロー・ブランチ戦略
├── 📄 LICENSE                  # ライセンス
│
├── 📂 cmd/                     # コマンドラインアプリケーション
│   └── 📂 server/              # サーバー起動用
│       └── 📄 main.go          # サーバーエントリーポイント
│
├── 📂 internal/                # 内部パッケージ（外部利用不可）
│   ├── 📂 config/              # 設定管理
│   │   ├── 📄 config.go
│   │   └── 📄 config_test.go
│   │
│   ├── 📂 models/              # データモデル
│   │   ├── 📄 movie.go         # 映画構造体
│   │   ├── 📄 tv.go            # TV番組構造体
│   │   ├── 📄 person.go        # 人物構造体
│   │   ├── 📄 review.go        # レビュー構造体
│   │   ├── 📄 common.go        # 共通構造体
│   │   └── 📄 models_test.go
│   │
│   ├── 📂 services/            # ビジネスロジック
│   │   ├── 📄 tmdb_client.go   # TMDb API クライアント
│   │   ├── 📄 cache.go         # キャッシュサービス
│   │   ├── 📄 rate_limiter.go  # レート制限管理
│   │   └── 📄 *_test.go
│   │
│   ├── 📂 handlers/            # HTTPハンドラー
│   │   ├── 📄 search.go        # 🔍 検索機能
│   │   ├── 📄 details.go       # 📋 詳細情報
│   │   ├── 📄 reviews.go       # ⭐ 評価・レビュー
│   │   ├── 📄 credits.go       # 🎭 キャスト・スタッフ
│   │   ├── 📄 trending.go      # 🔥 トレンド・おすすめ
│   │   ├── 📄 common.go        # 共通ハンドラー機能
│   │   └── 📄 *_test.go
│   │
│   ├── 📂 middleware/          # ミドルウェア
│   │   ├── 📄 cors.go          # CORS設定
│   │   ├── 📄 logging.go       # ログ機能
│   │   ├── 📄 auth.go          # 認証（将来拡張用）
│   │   └── 📄 middleware_test.go
│   │
│   └── 📂 utils/               # ユーティリティ
│       ├── 📄 response.go      # レスポンス共通処理
│       ├── 📄 errors.go        # エラーハンドリング
│       ├── 📄 validation.go    # バリデーション
│       └── 📄 utils_test.go
│
├── 📂 web/                     # フロントエンド
│   ├── 📂 static/              # 静的ファイル
│   │   ├── 📂 css/
│   │   │   └── 📄 style.css
│   │   ├── 📂 js/
│   │   │   ├── 📄 app.js       # メインアプリケーション
│   │   │   └── 📄 api.js       # APIクライアント
│   │   └── 📂 images/
│   │
│   └── 📂 templates/           # HTMLテンプレート
│       └── 📄 index.html       # メインページ
│
├── 📂 docs/                    # ドキュメント
│   ├── 📂 api/                 # API仕様書
│   │   ├── 📄 README.md        # API概要
│   │   └── 📄 endpoints.md     # エンドポイント詳細
│   ├── 📄 setup.md             # セットアップガイド
│   ├── 📄 examples.md          # 使用例
│   └── 📄 architecture.md      # システム構成
│
└── 📂 scripts/                 # スクリプト類
    ├── 📄 build.sh             # ビルドスクリプト
    ├── 📄 test.sh              # テストスクリプト
    └── 📄 deploy.sh            # デプロイスクリプト
```

## 🧪 テスト

```bash
# 全テスト実行
go test ./...

# カバレッジ付きテスト
go test -cover ./...

# 特定パッケージのテスト
go test ./internal/handlers/

# ベンチマークテスト
go test -bench=. ./...
```

## 🚀 デプロイ

### ローカル開発環境

```bash
# 開発サーバー起動（ホットリロード）
go run main.go

# ビルド
go build -o movie-api main.go

# 実行
./movie-api
```

### Docker使用

```bash
# イメージビルド
docker build -t movie-api .

# コンテナ起動
docker run -p 8080:8080 --env-file .env movie-api
```

### 本番環境

```bash
# 本番用ビルド
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o movie-api main.go

# systemdサービス登録
sudo systemctl enable movie-api
sudo systemctl start movie-api
```

## 📚 ドキュメント

| ドキュメント | 説明 |
|---|---|
| [API仕様書](docs/api/README.md) | REST APIの詳細仕様 |
| [セットアップガイド](docs/setup.md) | 環境構築・インストール手順 |
| [使用例](docs/examples.md) | 実践的なAPI利用例 |
| [システム構成](docs/architecture.md) | アーキテクチャ設計 |
| [開発フロー](README.github.md) | ブランチ戦略・コミット規約 |

## 🤝 コントリビューション

プロジェクトへの貢献を歓迎します！コントリビューションの前に [開発フロー](README.github.md) をご確認ください。

### 開発フロー

1. **Issues作成**: 実装する機能や修正内容をIssueで管理
2. **ブランチ作成**: `feature/機能名` または `fix/修正内容` 形式
3. **実装・コミット**: [コミットメッセージ規約](README.github.md#-コミットメッセージ)に従う
4. **プルリクエスト**: レビュー依頼とコードレビュー
5. **マージ**: レビュー完了後にmainブランチへマージ

### コミット例

```bash
# 新機能追加
feat: ✨ #15 映画検索エンドポイントを実装

# バグ修正  
fix: 🐛 #20 レスポンス形式のエラーを修正

# ドキュメント更新
docs: 📝 #23 API仕様書を更新
```

## 🛡️ ライセンス

このプロジェクトは [MIT License](LICENSE) の下で公開されています。

## 📞 お問い合わせ・サポート

- **作者**: [@takeshi-arihori](https://github.com/takeshi-arihori)
- **Issues**: [GitHub Issues](https://github.com/takeshi-arihori/movie-api/issues)
- **Discussion**: [GitHub Discussions](https://github.com/takeshi-arihori/movie-api/discussions)

## 🙏 謝辞

- **[The Movie Database (TMDb)](https://www.themoviedb.org/)**: 豊富な映画・TV番組データの提供
- **[Go Team](https://golang.org/)**: 優れたプログラミング言語の開発
- **オープンソースコミュニティ**: 多くのライブラリとツールの提供

---

**📌 注意**: このプロジェクトは教育・学習目的で作成されています。TMDb APIの利用規約を遵守してご使用ください。
