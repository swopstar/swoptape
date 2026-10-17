# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

swoptape is a self-hosted music streaming server, designed to run on weak hardware. The MVP is the server with library scanning, an embedded web player and a Subsonic-compatible API (`/rest`). After the MVP come native clients with offline playback: iOS first, then desktop, and later a native Android app.

The defining product goal is **dual-mode library navigation**. Users browse by folder structure _and_ by tags (artist, album, genre), both as first-class views of the same tracks. See [ADR-T006](docs/adr/T006-dual-mode-library.md).

swoptape is part of the **swopstar** suite alongside **swopcart** (a game library). They share a UI component library (`@swopstar/react-ui`, at `../ui`) and design language.

## Current state

The suite is moving to Rust ([ADR-0001](https://github.com/swopstar/meta/blob/main/adr/0001-rust.md)), and swoptape goes first. Its backend is being rewritten from Go to Rust on the long-lived branch `rewrite/rust` ([ADR-T001](docs/adr/T001-rust-rewrite.md)). **PRs target `rewrite/rust`, not `main`.** Milestones are in [`docs/roadmap.md`](docs/roadmap.md). The Go implementation on `main` is reference material. Read it with `git show main:<path>`. The most useful parts are:

- `www/subsonic/response/types.go`: Subsonic 1.16.1 response types
- `www/api/v0/openapi.yaml`: the first-party API's resource design
- `services/identity/`: users, sessions and application tokens

`Dockerfile`, `docker-compose.yml`, `.github/workflows/container.yml` and the frontend's API generation (`frontend/orval.config.ts`) still refer to the Go build. They are reworked in later milestones. Don't rely on them.

## Decisions

Decisions are recorded in two places:

- **Suite-wide** ADRs and standards (stack, licensing, versioning, commit conventions) are in [swopstar/meta](https://github.com/swopstar/meta), under `adr/` and `standards/`. Read them with `gh api repos/swopstar/meta/contents/<path> -H 'Accept: application/vnd.github.raw'`.
- **swoptape-specific** ADRs are in [`docs/adr/`](docs/adr/README.md), numbered `ADR-T###`. Some deliberately depart from suite-wide ones, and say so.

Follow `accepted` records and treat `proposed` ones as the default. If a change contradicts a swoptape ADR, update or supersede it in the same PR and say so in the description. Changes to suite-wide decisions are proposed in swopstar/meta, not here. Don't restate ADR content elsewhere. Link to it.

## Commands

Build orchestration uses [`just`](https://github.com/casey/just), with `rust.just` (backend) and `web.just` (frontend) as modules.

```sh
just rust serve        # run the server
just web serve         # run the frontend dev server (Vite HMR)
just serve             # both in parallel

just test              # all Rust tests
just rust test -p swoptape-server health   # targeted tests (args go to cargo test)
just fmt               # format everything
just check             # fmt check, clippy (warnings denied), cargo deny, reuse lint
just rust build        # release build

just local-db          # start local PostgreSQL via docker compose (dev/swoptape-dev-db/)
```

`just check` and `just test` run the same checks as CI. Both must pass before a PR is opened.

## Working on an issue

The workflow is described in [ADR-T003](docs/adr/T003-agentic-workflow.md). In practice:

1. Read the issue, every ADR it links, and the code in the paths it names. Stay inside that scope. If the issue is ambiguous or seems to conflict with an ADR, stop and ask on the issue rather than guessing.
2. Branch from `rewrite/rust` as `<issue-number>-<slug>`.
3. Write failing tests for the acceptance criteria first ([ADR-T002](docs/adr/T002-test-first.md)), then implement.
4. Run `just fmt`, `just check` and `just test`.
5. Commit following the [commit conventions](https://github.com/swopstar/meta/blob/main/standards/commit-conventions.md), then open a PR against `rewrite/rust` using the PR template. Close exactly one issue.
6. When reviewing someone else's PR: check that the tests would fail without the change, that the change stays within the issue's scope, and that it follows the ADRs.

## Code conventions

- **Layering:** `swoptape-server` → `swoptape-subsonic` → `swoptape-core` → `swoptape-db`. HTTP-facing crates never use `swoptape-db` directly.
- **Lints:** `unsafe_code` is forbidden and clippy's `pedantic` group is on, workspace-wide. Fix lints rather than allowing them. If an `#[allow]` is genuinely warranted, keep it narrow and add a short reason.
- **Errors:** library crates return `Result` with specific `thiserror` enums. No `unwrap`, `expect` or `panic!` outside tests and genuinely unreachable states. Cross-cutting domain errors (not found, stale, unique violation, not allowed, internal) live in `swoptape-core`. Error messages are not user-facing. The frontend maps error values to localised strings.
- **Async:** never block a tokio worker thread. Tag parsing, filesystem walks, password hashing and other CPU-bound or blocking calls go through `tokio::task::spawn_blocking` or a dedicated pool.
- **Logging:** use `tracing` with structured fields (`tracing::info!(status = 404, "request failed")`), not values interpolated into the message.
- **Tests:** prefer integration tests in each crate's `tests/` directory, which can only see the public API. Test behaviour, not implementation.
- **Dependencies:** add them to `[workspace.dependencies]` in the root `Cargo.toml` and reference them with `.workspace = true`. Look up the latest version when adding one. Weigh binary size and memory against convenience ([ADR-T004](docs/adr/T004-resource-budget.md)).
- **Subsonic types** serialise to both XML and JSON. Keep both. Many mobile clients still use XML.
- **Spelling:** New Zealand English in prose and docs, US English in identifiers.

## Licensing

REUSE is enforced by `reuse lint` in `just check`. Every source file starts with SPDX headers:

```rust
// SPDX-License-Identifier: AGPL-3.0-only
// SPDX-FileCopyrightText: 2026 Rareș Nistor
```

Contributors add their own `SPDX-FileCopyrightText` line when they first touch a file. Non-source files are covered by `REUSE.toml`. Licence policy is in the suite's [licensing standard](https://github.com/swopstar/meta/blob/main/standards/licensing.md): the server is AGPL-3.0-only, and individual infrastructure files may be marked MIT for later promotion to a shared library.

## Config and data

`data/` is gitignored and holds runtime config and generated keys. Don't add files from it to the index.
