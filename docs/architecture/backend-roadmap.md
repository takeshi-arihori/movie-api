# Backend Architecture Roadmap

この文書は実装順序と着手条件を示す。採用理由は [ADR-0001](../adr/0001-adopt-chi-as-http-router.md) と [ADR-0002](../adr/0002-async-jobs-and-cqrs.md) を参照する。各段階はユースケース単位のIssue/PRで進める。

## 現状（2026-09-23）

- `gorilla/mux`、`net/http`、TMDbクライアントで同期的なAPIを提供。
- `internal/handlers`、`internal/services`、`internal/models` が存在する。DDDの層分離は未完了。
- Job基盤、ローカルDB、Read Modelは未導入。

## Phase 1: chiとユースケース単位のDDD移行

- [ ] `gorilla/mux` を `chi` に置き換え、既存ルート・HTTP振る舞いの回帰テストを追加する。
- [ ] まず映画詳細取得など1ユースケースを選び、Presentation → Application → Domainの依存方向に整理する。
- [ ] TMDbとの通信・DTO変換をInfrastructureに置き、DomainをHTTPやTMDbの型から切り離す。
- [ ] 次のユースケースへの移行は個別Issue/PRで繰り返す。既存APIを一括で作り直さない。

## Phase 2: Changes同期のユースケース

着手条件: TMDbの変更をローカルに保存して利用する具体的な要件ができたとき。

- [ ] 同期対象（まずMovie、必要ならTV/Person）、更新頻度、保持項目を決める。
- [ ] 最小限のローカル永続化を用意し、Changes取得から対象詳細の更新までを実装する。
- [ ] 定期起動、Worker/Queue、失敗時の再試行・冪等性・レート制限を設計する。SQSは候補であり未確定。
- [ ] 更新とJob投入の整合性、同期漏れの検出・再実行を検証する。

永続化が同期の前提なので、このPhase内でDBとJobを一緒に扱う。採用するDBやQueueはその時点の要件で決める。

## Phase 3: ローカルデータの参照

着手条件: Phase 2のデータを参照する機能が必要になったとき。

- [ ] データ鮮度、欠損時の挙動、TMDbとの責務分担を定義する。
- [ ] 現在の参照経路を測定し、必要な画面・検索に限定して参照用データを整える。

## Phase 4: 必要な箇所だけCQRS

着手条件: 参照要件・計測結果から、通常の取得処理では複雑さや性能問題が解消できないとき。

- [ ] Search/Discoverなど対象ユースケースを選ぶ。
- [ ] Query側のモデルと更新側の整合性・許容遅延を明文化する。
- [ ] 対象ごとにQuery経路を分離し、効果を検証する。
