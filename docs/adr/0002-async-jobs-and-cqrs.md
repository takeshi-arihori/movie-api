# ADR-0002: 非同期JobとCQRSは必要なユースケースで導入する

- 状態: Accepted（現時点の方針）
- 日付: 2026-09-23

## 背景

現状のバックエンドはTMDbへの同期的な参照が中心で、Job基盤、ローカルDB、Read Modelを持たない。将来、TMDb Changesを使ったデータ同期や参照の最適化を検討している。chiはルーターであり、永続Jobの処理機構を提供しない。

## 決定

現在の同期APIのためにQueue、Worker、CQRSを追加しない。用途が具体化した段階で次を検討する。

1. TMDb Changesを定期取得し、変更された作品をローカルに同期するユースケースが必要になったら、永続化、Job、再試行、冪等性を一緒に設計する。
2. 実際の参照要件で、更新モデルからの取得が複雑・遅い場合に限り、そのユースケースのRead ModelとQuery経路を分離する。

TMDbのFavoriteやRatingなど単発のPOSTも、非同期化する要件がない限り同期的に扱う。JobとCQRSの導入はそれぞれ別の設計判断とPRで確定する。

## 判断理由

- 現在のTMDb参照APIにはJobの完了状態、失敗通知、再試行を管理する要件がない。
- 永続データのない段階でRead Modelを先に設計しても、解決すべき参照課題を測定できない。
- 抽象を先に追加せず、必要なユースケースでApplication側にポートを定義し、Infrastructure側でQueueやTMDbの実装を担う方針にする。

## 将来の検討事項

Changes同期を実装する際は、取得期間とページング、429等のレート制限、再試行とバックオフ、重複配信時の冪等性、失敗時の観測・再実行、DB更新とJob投入の整合性を設計する。QueueはSQSを候補とし、DBとQueueをまたぐ更新には必要に応じてOutbox等を検討する。SQSの採用や特定の構成はこのADRでは確定しない。

CQRSは全APIに適用せず、ローカルDBのSearchやDiscover等で必要性を確認する。読み書きの責務分離と、別のRead Modelの導入は区別して判断する。

## 影響

現時点のAPIにJob状態や `202 Accepted` を追加しない。後続の実装順と導入条件は [Backend Roadmap](../architecture/backend-roadmap.md) に記載する。
