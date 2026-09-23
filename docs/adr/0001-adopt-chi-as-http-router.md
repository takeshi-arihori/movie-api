# ADR-0001: HTTPルーターにchiを採用する

- 状態: Accepted（方針決定、実装は別PR）
- 日付: 2026-09-23

## 背景

現在のバックエンドは `net/http + gorilla/mux` で動作している。GoのHTTPルーターをGin、Echo、chiから選び直す。DDDへの移行はユースケースごとに進め、依存方向を Presentation → Application → Domain とする。Domain/ApplicationはWebフレームワークやTMDbクライアントに依存させない。

## 決定

`chi` をHTTPルーターとして採用する。既存の `gorilla/mux` は後続の独立したPRで置き換える。このADRの追加だけでは実行コードや依存ライブラリは変更しない。

## 判断理由

- chiのルーティングとミドルウェアは `net/http` の型に沿うため、HTTP処理をPresentation層に閉じ込めやすい。
- 既存のHTTPハンドラーや `httptest` を活かしながら段階的に移行できる。
- このプロジェクトのAPI規模では、組み込みBinding等よりも依存を小さく保つ利点を優先する。

DDDの依存方向はGin/Echoでも守れる。chi自体がDDDを保証するわけではない。

## 比較した選択肢

| 候補 | 利点 | 今回の判断 |
| --- | --- | --- |
| Gin | Binding・Validationなどの機能が充実 | 固有のContextをPresentation内に限定する設計が必要 |
| Echo | Bindingや共通のHTTP機能が充実 | 同様に固有のContextをPresentation内に限定する設計が必要 |
| chi | `net/http` 互換の軽量なルーティング | 今回の移行方針に合うため採用 |

## 影響・移行時の確認事項

- Binding、Validation、JSON応答の共通処理は必要な機能だけ選ぶ。入力検証とDomainの不変条件は分けて扱う。
- `gorilla/mux` の数値IDルート制約、405/404、OPTIONS/CORS、既存URLとレスポンスを移行テストで確認する。
- コード移行が完了したら `CLAUDE.md` 等に記載された現行スタックを `chi` に更新する。
- 機能や性能上の実測課題が出た場合は、別ADRで見直す。
