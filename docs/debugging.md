# 🐛 デバッグガイド - Movie API

このドキュメントでは、Movie APIプロジェクトでのDelveデバッガーとpprofプロファイラーの使用方法を説明します。

## 🔧 前提条件

- Docker & Docker Compose
- Go 1.24+
- VS Code (推奨) + Go拡張機能
- TMDb API キー

## 🚀 クイックスタート

### 1. 環境設定
```bash
# 環境変数ファイルをコピー
cp .env.example .env

# .envファイルを編集してTMDB_API_KEYを設定
vim .env
```

### 2. デバッグ環境起動
```bash
# デバッグモードでサービス起動
make debug

# または手動で
docker compose --profile debug up -d
```

## 🐛 Delve デバッガー

### Docker内でのリモートデバッグ

#### 起動方法
```bash
# デバッグ環境起動
make debug-detached

# ログ確認
make logs-debug
```

#### VS Codeでの接続
1. VS Codeで「Run and Debug」パネルを開く
2. 「Debug Go (Docker Remote)」を選択
3. F5キーでデバッグ開始

#### ポート情報
- **アプリケーション**: http://localhost:8080
- **Delveデバッガー**: localhost:2345
- **pprof**: http://localhost:6060/debug/pprof/

#### ブレークポイント設定
1. VS Codeで`.go`ファイルを開く
2. 行番号の左をクリックして赤丸のブレークポイントを設置
3. デバッガーがブレークポイントで停止
4. 変数の値、スタックトレースを確認可能

### ローカルデバッグ

#### 前提条件
```bash
# データベースサービスのみ起動
docker compose up -d postgres redis
```

#### VS Codeでの起動
1. 「Debug Go (Local)」を選択
2. F5キーでデバッグ開始

## 📊 pprof プロファイラー

### 基本的な使用方法

#### Web UI
デバッグ環境起動後、以下のURLにアクセス：
```
http://localhost:6060/debug/pprof/
```

#### 主要なエンドポイント
- `/debug/pprof/` - 概要ページ
- `/debug/pprof/goroutine` - ゴルーチン情報
- `/debug/pprof/heap` - ヒープメモリ
- `/debug/pprof/profile` - CPUプロファイル（30秒間）
- `/debug/pprof/trace` - 実行トレース

#### コマンドライン使用
```bash
# CPUプロファイル取得（30秒間）
go tool pprof http://localhost:6060/debug/pprof/profile

# ヒーププロファイル取得
go tool pprof http://localhost:6060/debug/pprof/heap

# ゴルーチン情報
go tool pprof http://localhost:6060/debug/pprof/goroutine
```

#### プロファイル分析例
```bash
# インタラクティブモード
go tool pprof http://localhost:6060/debug/pprof/profile
(pprof) top10    # 上位10関数表示
(pprof) list main.main  # main関数の詳細
(pprof) web      # ブラウザで可視化
```

## 🔍 デバッグシナリオ

### 1. APIエンドポイントのデバッグ

#### 手順
1. `internal/handlers/movie_handler.go`にブレークポイント設置
2. デバッグモードで起動
3. curlでAPIリクエスト送信
```bash
curl http://localhost:8080/health
```
4. ブレークポイントで停止、変数確認

### 2. データベース接続のデバッグ

#### 手順
1. `internal/config/config.go`の`Load()`関数にブレークポイント
2. 環境変数の値を確認
3. データベース接続エラーを調査

### 3. パフォーマンス問題の特定

#### CPU使用率調査
```bash
# 30秒間のCPUプロファイル取得
go tool pprof http://localhost:6060/debug/pprof/profile

# 結果分析
(pprof) top10
(pprof) list 関数名
```

#### メモリリーク調査
```bash
# ヒープ情報取得
go tool pprof http://localhost:6060/debug/pprof/heap

# メモリ使用量分析
(pprof) top10 -cum
(pprof) list 関数名
```

## 🛠️ トラブルシューティング

### よくある問題

#### 1. Delveに接続できない
```bash
# コンテナが起動しているか確認
docker compose ps

# ログ確認
docker compose logs backend-debug

# ポートが開いているか確認
netstat -tlnp | grep 2345
```

#### 2. pprofにアクセスできない
- 環境変数`ENV=development`が設定されているか確認
- ポート6060が他のプロセスで使用されていないか確認

#### 3. ブレークポイントで停止しない
- ソースコードがコンテナ内のものと一致しているか確認
- コンパイル最適化が無効になっているか確認（`-gcflags="all=-N -l"`）

### デバッグ環境のリセット
```bash
# 全てのサービス停止・削除
make clean

# 再ビルド・起動
make build
make debug
```

## ⚙️ 設定カスタマイズ

### VS Code設定

#### launch.json カスタマイズ
`.vscode/launch.json`を編集して、環境変数やデバッグオプションを調整可能。

#### 推奨拡張機能
- Go (golang.go)
- Docker (ms-azuretools.vscode-docker)
- Rest Client (humao.rest-client) - API テスト用

### Docker設定

#### セキュリティ設定
デバッグコンテナでは以下の設定でデバッガーのアタッチを許可：
```yaml
security_opt:
  - apparmor:unconfined
  - seccomp:unconfined
cap_add:
  - SYS_PTRACE
```

## 📚 参考資料

### Delve
- [公式ドキュメント](https://github.com/go-delve/delve)
- [VS Code Go拡張機能](https://code.visualstudio.com/docs/languages/go)

### pprof
- [公式ドキュメント](https://golang.org/pkg/net/http/pprof/)
- [pprofチュートリアル](https://blog.golang.org/pprof)

### Docker
- [Docker Compose公式ドキュメント](https://docs.docker.com/compose/)

## 🎯 ベストプラクティス

1. **ブレークポイントは適切に削除** - デバッグ完了後は不要なブレークポイントを削除
2. **プロファイリングは本番環境で無効化** - `ENV=production`では自動的に無効
3. **ログレベルの活用** - デバッグ時は`LOG_LEVEL=debug`に設定
4. **リソース監視** - デバッグ環境でのメモリ・CPU使用量を定期的に確認

---

💡 **Tip**: `make help`コマンドで利用可能な全てのMakefileコマンドを確認できます。