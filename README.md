# todo — Go製TUI/CLIタスク管理ツール

ターミナルで完結するタスク管理ツール。TUIによるインタラクティブ操作とCLIコマンドの両方に対応し、Google Calendarとの同期機能を備える。

---

## 特徴

- **TUI / CLI の両対応** — キーボードで直感的に操作できるTUIと、スクリプトやエイリアスに組み込めるCLIを統一したデータで管理
- **サブタスク・タグ対応** — タスクの階層化とタグによる分類が可能
- **Google Calendar 同期** — 双方向同期・フィルタ指定・自動同期に対応
- **SQLite による永続化** — `~/.config/todo/todo.db` に一元管理

---

## インストール

```bash
git clone https://github.com/yourname/todo
cd todo
go build -o todo .
sudo mv todo /usr/local/bin/
```

---

## 使い方

### TUIモード

```bash
todo tui
```

| キー              | 操作                                 |
| ----------------- | ------------------------------------ |
| `j` / `k`         | カーソル移動                         |
| `Space` / `Enter` | 完了トグル                           |
| `a`               | タスク追加                           |
| `d`               | タスク削除                           |
| `p`               | 優先度変更                           |
| `Tab`             | フィルタ切替（全件 / 未完了 / 完了） |
| `q`               | 終了                                 |

### CLIモード

**タスク操作**

```bash
todo add "タスクタイトル"
todo add "タスク" --priority high
todo add "タスク" --due 2026-06-30
todo add "タスク" --tag work --tag dev
todo add "サブタスク" --parent <parent-id>

todo list
todo list --filter undone
todo list --filter-tag work
todo list --priority high

todo done <id>
todo delete <id>
```

**タグ操作**

```bash
todo tag list
todo tag add <task-id> work
todo tag remove <task-id> work
```

**同期**

```bash
# 方向指定
todo sync                         # 双方向（デフォルト）
todo sync --push                  # ローカル → GCal
todo sync --pull                  # GCal → ローカル

# フィルタ指定
todo sync --filter-tag work
todo sync --filter-due 2026-06-30
todo sync --filter-priority high

# 確認
todo sync --dry-run               # 差分表示のみ（実行しない）
todo sync --status                # 未同期タスク一覧

# 自動同期設定
todo sync --auto enable --interval 30m
todo sync --auto disable
```

---

## 設定

`~/.config/todo/config.toml`

```toml
[sync]
auto     = true
interval = "30m"
on_start = true   # TUI起動時に自動同期するか
```

---

## データモデル

### tasks

| カラム          | 型          | 説明                                 |
| --------------- | ----------- | ------------------------------------ |
| `id`            | TEXT (UUID) | 主キー                               |
| `title`         | TEXT        | タスク名                             |
| `description`   | TEXT        | 詳細（任意）                         |
| `done`          | INTEGER     | 完了フラグ                           |
| `priority`      | INTEGER     | 1=high / 2=mid / 3=low               |
| `due_date`      | TEXT        | 期日（ISO8601）                      |
| `parent_id`     | TEXT        | 親タスクID（サブタスク用）           |
| `gcal_event_id` | TEXT        | GCal イベントID                      |
| `sync_status`   | TEXT        | `local_only` / `synced` / `conflict` |
| `created_at`    | TEXT        | 作成日時                             |
| `updated_at`    | TEXT        | 更新日時                             |

### tags / task_tags

タグは多対多で管理。`task_tags` テーブルで junction を構成する。

### sync_log

同期履歴と競合記録を保持。コンフリクト解決の手がかりに使う。

---

## アーキテクチャ

```
cmd/                        # CLIエントリポイント（cobra）
  root.go
  add.go / list.go / done.go / delete.go
  tag.go
  sync.go
  tui.go

internal/
  domain/                   # 型定義（Task, Tag, Filter, SyncStatus）
  repository/
    interface.go            # TaskRepository, TagRepository インターフェース
    sqlite/                 # SQLite実装
  service/                  # ビジネスロジック（インターフェースに依存）
    task.go / tag.go / sync.go
  sync/
    interface.go            # Syncer インターフェース
    scheduler.go            # 自動同期（goroutine）
    gcal/                   # Google Calendar実装
  tui/                      # Bubble Tea TUI

config/                     # config.toml 読み書き
main.go
```

**依存の流れ**

```
cmd層
  └─ service（ビジネスロジック）
        ├─ TaskRepository interface ← sqlite実装
        └─ Syncer interface        ← gcal実装
```

`service` 層はインターフェースにのみ依存するため、GCal以外の連携先（Notion、GitHub Issues等）を追加してもservice層は無変更。

---

## 技術スタック

| 用途             | ライブラリ                                                  |
| ---------------- | ----------------------------------------------------------- |
| CLI              | [cobra](https://github.com/spf13/cobra)                     |
| TUI              | [Bubble Tea](https://github.com/charmbracelet/bubbletea)    |
| DB               | [modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite) |
| マイグレーション | [golang-migrate](https://github.com/golang-migrate/migrate) |
| GCal             | Google Calendar API v3                                      |

---

## 開発ロードマップ

| フェーズ | 内容                                        |
| -------- | ------------------------------------------- |
| 1        | `domain/` 型定義・`repository/interface.go` |
| 2        | SQLite接続・マイグレーション実装            |
| 3        | `TaskRepository` SQLite実装                 |
| 4        | `service/task.go` でCRUD疎通                |
| 5        | `cmd/add`, `cmd/list` でCLI動作確認         |
| 6        | Bubble Tea TUI実装                          |
| 7        | GCal連携・Syncer実装                        |

---

## MVP スコープ外

- 複数リスト管理
- クラウドバックアップ（GCal以外）
- 繰り返しタスク
- チーム共有
