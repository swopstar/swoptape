# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

swoptape is a self-hosted music streaming server. It exposes three API surfaces: its own versioned REST API (`/api`), a frontend RPC channel (`/rpc`), and a Subsonic-compatible API (`/rest`) for compatibility with existing Subsonic clients.

The long-term vision includes a first-party swopstar API (`/api`) for native clients, but this is far off — Subsonic/OpenSubsonic compatibility is the priority for the foreseeable future.

swoptape is part of the **swopstar** suite alongside **swopcart** (same concept, video game library). They share a UI component library (`@swopstar/react-ui`, at `/Users/razz/src/swopstar/ui`) and design language. swopcart came first but stalled. swoptape was started to regain momentum on a more tractable problem — music is well-understood, game cataloguing is harder. The plan is to return to swopcart once swoptape is mature and the tooling is proven.

The defining product goal is **dual-mode library navigation**: users can browse by file system structure (folders, sub-folders) _and_ by tags (artist, album, genre) simultaneously. Most alternatives (Jellyfin, Navidrome, Plex) force a choice — they either flatten everything into a tag hierarchy and lose folder organisation, or expose raw folder browsing with no tag context. swoptape treats both as first-class. This affects library scanning (both paths and tags must be preserved), the data model (tracks need folder ancestry, not just tag metadata), and the API design (endpoints should support both navigation modes).

## Licensing

This project uses dual licensing enforced by [REUSE](https://reuse.software/). `reuse lint` runs as part of `just check`.

Every file must carry SPDX headers. New contributors add their own `SPDX-FileCopyrightText` line when they first touch a file:

```go
// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Rareș Nistor
// SPDX-FileCopyrightText: 2027 Contributor Name
```

**MIT** — packages likely to move into `gokit` or shared libraries:

- `config/`
- `database/`
- `services/` (top-level)
- `services/domain/`
- `services/identity/`
- `www/` (top-level — server and route setup only)
- `www/api/` (router installer)
- `www/rpc/` (router installer)

**AGPL-3.0-only** — everything else: `cmd/`, `frontend/`, `www/api/v0/`, `www/subsonic/`, and handler packages generally.

Non-source files (JSON, YAML, binaries, etc.) are covered by `REUSE.toml` rather than inline headers.

## Commands

Build orchestration uses [`just`](https://github.com/casey/just) with `go.just` (backend) and `web.just` (frontend) as sub-modules.

```sh
just go serve          # run backend with air (hot reload)
just web serve         # run frontend dev server (Vite HMR)
just serve             # both in parallel

just go build          # compile to tmp/bin/swoptape
just go test           # go test ./...
just go test ./services/identity/... -run TestFoo   # target specific tests
just go check          # golangci-lint + go mod verify
just go check-tidy     # verify go.mod/go.sum are tidy
just go check-vuln     # govulncheck (slower, run periodically)
just go tidy           # go mod tidy

just local-db          # start local PostgreSQL via docker compose (dev/swoptape-dev-db/)
```

Config is loaded from `data/config.toml`. The file is gitignored and auto-generated with defaults on first run.

## Architecture

### Request flow

```
gin router (www/)
  ├── /api/v0/   → www/api/v0/
  ├── /rpc/      → www/rpc/        (frontend RPC, stub)
  ├── /rest/     → www/subsonic/   (Subsonic API)
  └── *          → embedded frontend SPA (frontend/dist)
                   returns 404 for /api, /rpc, /rest prefixes;
                   falls back to index.html for all other 404s
```

Each sub-router implements `Install(*gin.RouterGroup)`. They are constructed in `www/server.go` and mounted in `www/routes.go`.

### Testing

Test behaviour, not implementation. Tests should survive internal refactors — if a test breaks because a private method was renamed or a query was restructured (but the observable outcome is unchanged), the test is testing the wrong thing. Prefer fewer, higher-confidence tests over exhaustive coverage of internals.

Use external test packages (`package foo_test`) rather than internal (`package foo`). This enforces testing through the public API and makes it structurally impossible to test implementation details.

### Layering rule

`www/` must never import `database/` directly. All data access goes through `services/`. The dependency chain is strictly: `www` → `services` → `database`.

### Service layer

`services/Services` is the single struct passed into the web server. Each domain area lives in its own sub-package (e.g. `services/identity`). Shared domain errors live in `services/domain` to avoid import cycles — service sub-packages import `domain`, and `services` imports the sub-packages.

`services/identity` uses a **rich domain object** pattern:

- `Service` owns queries and lifecycle (implements `IService`)
- `User` wraps a cached `database.User` and carries behaviour (implements `IUser`)
- `Session` wraps a cached `database.Session` (implements `ISession`)
- `ApplicationToken` wraps a cached `database.ApplicationToken` (implements `IApplicationToken`)
- Read methods operate on the in-memory cache — no DB call, no error
- Mutating methods write to DB then update the cache field directly
- `Reload(ctx)` re-fetches from DB when a fresh snapshot is needed

`IService` covers user management, JWT-based sessions, and application tokens. The JWT signing key is an Ed25519 key auto-generated at `data/jwt.key` on first run.

Handler tests mock `IService`. Service tests use a real DB.

### Error handling

Validation errors are collected with `errors.Join` so multiple failures are returned at once. Service-specific errors (e.g. `ErrUsernameTooShort`) live alongside the service code. Shared cross-service errors live in `services/domain`:

| Error           | Meaning                                                    |
| --------------- | ---------------------------------------------------------- |
| `ErrNotFound`   | Row doesn't exist                                          |
| `ErrStale`      | Row disappeared between load and write (rowsAffected != 1) |
| `ErrUnique`     | Unique constraint violation                                |
| `ErrNotAllowed` | Authenticated but forbidden                                |
| `ErrInternal`   | Unexpected error                                           |

Error message strings are not user-facing — the frontend maps error values to localised strings.

### Database

GORM via `gokit/database`. Models registered in `database/models.go`. `db.AutoMigrate()` runs at startup. Current models: `User`, `UserPermission`, `Session`, `ApplicationToken`.

### Config

TOML, loaded from `data/config.toml` (gitignored). Sections: `server` (address, timeouts), `database`, `auth.jwt` (key path, access/refresh TTLs, issuer), `jobs`. `data/` is covered by `.gitignore` — don't add files from it to the index.

### Subsonic response types

`www/subsonic/response/` contains Go types matching the Subsonic 1.16.1 schema, originally generated from the XSD and now maintained directly. `DateTime` (in `datetime.go`) wraps `time.Time` with RFC3339 XML attribute and JSON marshaling. All types serialise to both XML and JSON via struct tags.

The current target is Subsonic 1.16.1. Adopting [OpenSubsonic](https://opensubsonic.netlify.app/) extensions is a future goal. OpenSubsonic dropped XML, but XML support should be retained alongside JSON for compatibility with the broadest range of mobile clients — the dual XML/JSON serialisation on the response types is intentional and should be preserved.
