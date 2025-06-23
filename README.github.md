# 📝 開発フロー・ブランチ戦略

## 🔄 ブランチ戦略：GitHub Flow

シンプルで効率的な**GitHub Flow**を採用し、チーム開発に最適化した運用を行います。

### 基本方針

- **mainブランチ**: 常にデプロイ可能な状態を保つ
- **featureブランチ**: mainから作成し、新機能実装後にmainへマージ
- **fixブランチ**: mainから作成し、バグ修正後にmainへマージ
- **シンプルなブランチ構成**でチーム開発に最適

### ブランチ運用ルール

1. **mainブランチへの直接pushは禁止**
2. **必ずfeature/fixブランチを作成**してPull Requestでマージ
3. **マージ後はfeature/fixブランチを削除**
4. **1 Issue = 1 Branch = 1 Pull Request**を基本とする

## 🚀 GitHubフローの基本

### 1. Issues作成
実装する機能や修正内容をIssueで管理

```markdown
# Issue例
## 概要
映画検索エンドポイントを実装します。

## タスク
- [ ] ハンドラー関数の実装
- [ ] バリデーションの追加
- [ ] テストの作成

## 完了条件
- [ ] エンドポイントが正常に動作する
- [ ] テストが全て通る
```

### 2. ブランチ作成
Issue番号と機能名を含むブランチを作成

```bash
# featureブランチ作成例
git checkout main
git pull origin main
git checkout -b feature/15-search-endpoint

# fixブランチ作成例  
git checkout -b fix/20-response-format
```

### 3. 実装・コミット
[コミットメッセージ規約](#-コミットメッセージ)に従って実装

```bash
# 実装
git add .
git commit -m "feat: ✨ #15 映画検索エンドポイントを実装"

# プッシュ
git push origin feature/15-search-endpoint
```

### 4. プルリクエスト作成
GitHubでPull Requestを作成し、レビューを依頼

### 5. コードレビュー
チームメンバーによるコードレビューを実施

### 6. マージ
レビュー完了後にmainブランチへマージ

```bash
# マージ後のクリーンアップ
git checkout main
git pull origin main
git branch -d feature/15-search-endpoint
```

## 🏷️ ブランチ命名規則

### フォーマット
```
<type>/<issue-number>-<short-description>
```

### Type分類

| Type | 説明 | 例 |
|---|---|---|
| `feature` | 新機能実装 | `feature/15-search-endpoint` |
| `fix` | バグ修正 | `fix/20-json-response-format` |
| `docs` | ドキュメント更新 | `docs/23-api-documentation` |
| `refactor` | リファクタリング | `refactor/12-handler-structure` |
| `test` | テスト追加・修正 | `test/16-unit-tests` |
| `chore` | 設定変更・メンテナンス | `chore/11-project-setup` |

### 命名例

```bash
# 良い例
feature/15-search-endpoint           # 検索エンドポイント実装
feature/17-user-authentication       # ユーザー認証機能
fix/20-json-response-format          # JSONレスポンス形式修正
docs/23-api-documentation            # APIドキュメント更新
refactor/14-tmdb-client              # TMDbクライアントリファクタリング

# 避けるべき例
search                               # 何のブランチか不明
fix-bug                              # 具体性に欠ける
feature-new                          # 内容が分からない
```

## 📋 コミットメッセージ

### 参考資料
- [Gitのコミットメッセージの書き方（2023年ver.）](https://zenn.dev/itosho/articles/git-commit-message-2023)
- [Conventional Commits](https://www.conventionalcommits.org/)
- [gitmoji](https://gitmoji.dev/)

### 基本フォーマット

```
<Type>: <Emoji> #<Issue Number> <Title>
```

### Type（必須）

| Type | 説明 | 例 |
|---|---|---|
| `feat` | ユーザー向けの機能の追加や変更 | 新しいエンドポイント追加 |
| `fix` | ユーザー向けの不具合の修正 | API レスポンスエラー修正 |
| `docs` | ドキュメントの更新 | README更新、API仕様書追加 |
| `style` | フォーマットなどのスタイルに関する修正 | コードフォーマット、lint修正 |
| `refactor` | リファクタリングを目的とした修正 | コード構造改善、性能最適化 |
| `test` | テストコードの追加や修正 | 単体テスト、統合テスト追加 |
| `chore` | タスクファイルなどプロダクションに影響のない修正 | 依存関係更新、設定ファイル変更 |

### Emoji（任意）

よく使用するgitmojiの例：

| Emoji | Code | 説明 |
|---|---|---|
| ✨ | `:sparkles:` | 新機能追加 |
| 🐛 | `:bug:` | バグ修正 |
| 📝 | `:memo:` | ドキュメント更新 |
| 🎨 | `:art:` | コード構造・フォーマット改善 |
| ⚡ | `:zap:` | パフォーマンス改善 |
| 🔧 | `:wrench:` | 設定ファイル更新 |
| ✅ | `:white_check_mark:` | テスト追加・更新 |
| 🚀 | `:rocket:` | デプロイ関連 |

### Issue Number（強く推奨）

- そのコミットに紐づく**Issue番号を必ず記載**
- GitHub上でリンクされ、トラッキングしやすくなる
- `#123` 形式で記載

### Title（必須・日本語推奨）

- 変更内容を**現在形で記載**
- **20〜30文字以内**が適切
- 具体的で分かりやすい内容

## 💡 コミットの粒度・品質

### 基本原則

1. **1 Commit = 1つの意味あるまとまり**であるべき
2. レビュアーがPull Requestを見たときに**"ストーリー"が分かる**ことを意識
3. **1 Issue、1 Pull Request、1 Commit が理想**（複雑な場合は複数コミットでも可）

### 良いコミット例

```bash
# 新機能追加
feat: ✨ #15 映画検索エンドポイントを実装
feat: ✨ #16 映画詳細取得APIを追加
feat: ✨ #17 レビュー投稿機能を実装

# バグ修正
fix: 🐛 #20 JSONレスポンス形式のエラーを修正
fix: 🐛 #21 セッションタイムアウトの問題を解決

# ドキュメント更新
docs: 📝 #23 API仕様書を更新
docs: 📝 #23 READMEセットアップガイドを追加

# リファクタリング
refactor: ⚡ #14 TMDbクライアントのパフォーマンス改善
refactor: 🎨 #12 ハンドラー関数の構造を改善

# テスト追加
test: ✅ #15 検索エンドポイントの単体テストを追加
test: ✅ #16 映画詳細APIの統合テストを実装

# 設定・メンテナンス
chore: 🔧 #11 プロジェクト初期設定を追加
chore: 📦 #11 依存関係を更新
```

### 避けるべきコミット例

```bash
# ❌ 避けるべき例
update code                          # 何を更新したか不明
fix bug                             # どのバグを修正したか不明
add feature                         # どの機能を追加したか不明
とりあえず保存                        # 作業途中のコミット
WIP                                 # 作業中を示すが内容不明
ログイン機能                          # Issue番号がない
feat: 映画API, レビューAPI, 検索APIを実装  # 複数機能を1コミットに含める
```

## 🔍 コミットメッセージの詳細指針

### Whyについて

- **Whyはコミットメッセージではなく、IssueやPull Requestで説明**
- コミットメッセージには**Issue番号を必ず紐付ける**
- Subjectは**What に寄った書き方でOK**

### 例：詳細な説明が必要な場合

```bash
# コミットメッセージ（What）
feat: ✨ #14 TMDb APIクライアントのキャッシュ機能を追加

# Issue #14での説明（Why）
## 背景
現在、TMDb APIへの同一リクエストが重複して発生し、
レスポンス時間が遅く、レート制限に引っかかる可能性がある。

## 解決方法
インメモリキャッシュを実装し、一定時間内の同一リクエストは
キャッシュされた結果を返すようにする。

## 期待される効果
- レスポンス時間の短縮
- TMDb APIへのリクエスト数削減
- レート制限回避
```

## 📊 プルリクエストのベストプラクティス

### プルリクエストテンプレート

```markdown
## 概要
<!-- このPRで何を実装・修正したかを簡潔に説明 -->

## 関連Issue
Closes #XXX

## 変更内容
- [ ] 機能A を実装
- [ ] 機能B を修正
- [ ] テストを追加

## テスト方法
<!-- 動作確認の手順を記載 -->
1. サーバーを起動
2. `curl "http://localhost:8080/api/v1/search?query=test"` を実行
3. 正常なレスポンスが返ることを確認

## チェックリスト
- [ ] テストが追加され、全て通る
- [ ] ドキュメントが更新されている（必要な場合）
- [ ] コードレビューの準備ができている

## スクリーンショット
<!-- UI変更がある場合は画像を添付 -->
```

### レビュー観点

1. **機能要件**: 仕様通りに実装されているか
2. **コード品質**: 可読性、保守性、効率性
3. **テスト**: 適切なテストが書かれているか
4. **セキュリティ**: 脆弱性がないか
5. **パフォーマンス**: 性能に問題がないか

## 🎯 チーム開発のルール

### 1. コミュニケーション

- **Issue**で要件・仕様を明確にする
- **Pull Request**で実装内容とレビューポイントを説明
- **コメント**で疑問点や改善提案を積極的に共有

### 2. コードレビュー

- **建設的なフィードバック**を心がける
- **なぜそうするべきか**の理由も併せて説明
- **学習機会**として捉え、知識を共有する

### 3. 品質管理

- **テストファースト**を意識
- **継続的インテグレーション**でコード品質を維持
- **ドキュメント**を常に最新に保つ

## 🚨 緊急時のホットフィックス

### 手順

1. **mainブランチから直接hotfixブランチを作成**
2. **最小限の修正で問題を解決**
3. **即座にPull Requestを作成**
4. **レビュー後、緊急マージ**

```bash
# ホットフィックス例
git checkout main
git pull origin main
git checkout -b hotfix/critical-security-fix
# 修正作業
git commit -m "fix: 🚨 #XXX 重要なセキュリティ問題を修正"
git push origin hotfix/critical-security-fix
# Pull Request作成・緊急レビュー
