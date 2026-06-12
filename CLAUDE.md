# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project state

This is a **terminal todo manager (TUI + CLI)** in its earliest stage. The repo is currently a
fresh `cobra-cli` scaffold: only `main.go`, `cmd/root.go`, and `cmd/add.go` exist, and they still
contain generated boilerplate (placeholder `Short`/`Long` strings, an unused `--toggle` flag, no
real `Run` logic).

`README.md` is the **design spec / roadmap**, not a description of working code. It documents the
intended architecture (`internal/domain`, `repository/sqlite`, `service`, `sync/gcal`, `tui`),
SQLite schema, CLI surface, and Google Calendar sync — almost none of which is built yet. When the
README and the code disagree, the README is the target and the code is the current reality. Treat
the README's "開発ロードマップ" (phases 1–7) as the implementation order.

Note the module path is `example.com/todo_tui` (placeholder) while the README refers to the binary
as `todo`.

## Commands

```bash
go build -o todo .        # build binary
go run . <subcommand>     # run without building, e.g. `go run . add`
go test ./...             # run all tests (none exist yet)
go test ./cmd/ -run TestName   # run a single test once tests are added
go vet ./...              # static checks
go mod tidy               # sync go.mod/go.sum after adding deps
```

Go 1.25.3.

## Architecture (target, per README)

Dependencies flow strictly inward via interfaces — this is the key constraint to preserve as code
is added:

```
cmd (cobra)  →  service (business logic)  →  domain (types)
                  ├─ repository.TaskRepository / TagRepository  ← internal/repository/sqlite
                  └─ sync.Syncer                                ← internal/sync/gcal
```

The `service` layer depends **only on interfaces**, never on concrete SQLite or Google Calendar
implementations. New integrations (Notion, GitHub Issues, alternate storage) must be added behind
the existing interfaces without modifying `service`. Persistence target is SQLite at
`~/.config/todo/todo.db`; config lives at `~/.config/todo/config.toml`.

## Adding cobra subcommands

New commands follow the pattern in `cmd/add.go`: declare a `*cobra.Command` and register it in the
file's `init()` with `rootCmd.AddCommand(...)`. Replace the generated boilerplate (`Short`, `Long`,
the stub `Run`) rather than copying it forward into new files.
