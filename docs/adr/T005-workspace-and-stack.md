# ADR-T005: Workspace layout and project dependencies

**Status:** proposed
**Date:** 2026-10-17

## Context

[ADR-0001](https://github.com/swopstar/meta/blob/main/adr/0001-rust.md) sets the suite's Rust stack: tokio, axum, PostgreSQL through sqlx, and tracing. swoptape needs a module layout on top of that, with clear boundaries so that agents can work in parallel and client code can be shared, plus a few music-specific dependencies.

## Decision

A single Cargo workspace:

| Crate               | Responsibility                                                                    |
| ------------------- | --------------------------------------------------------------------------------- |
| `swoptape-core`     | Domain model and library logic: scanning, tag reading, folder and tag navigation  |
| `swoptape-db`       | PostgreSQL schema, migrations and queries                                         |
| `swoptape-subsonic` | Subsonic 1.16.1 response types and handlers                                       |
| `swoptape-server`   | HTTP server, `/api`, mounting `/rest`, the embedded web player, the `swoptape` binary |
| `swoptape-client`   | Post-MVP: shared client core (see [ADR-T009](T009-shared-client-core.md))         |

Dependency direction: `server` → `subsonic` → `core` → `db`. HTTP-facing crates never use `db` directly.

Project dependencies, on top of ADR-0001's stack:

- **Tags:** `lofty`.
- **Transcoding:** an external `ffmpeg` process, only when requested.
- **Migrations:** `sqlx`'s migration runner, with migrations in `swoptape-db`.
- **Errors:** `thiserror` in library crates. The binary maps errors to responses at the edge.
- **Frontend:** the React app from [ADR-R001](https://github.com/swopstar/meta/blob/main/projects/react-ui/adr/R001-stack.md), built with Vite and embedded into the binary.

## Alternatives considered

- **`symphonia` for tags.** It also decodes audio, which swoptape doesn't need, and its tag support is narrower than `lofty`'s.
- **In-process transcoding.** Avoids depending on `ffmpeg`, but adds codec libraries to every build, even where transcoding is never used.
- **A single crate.** Less ceremony, but weaker boundaries for parallel agent work and no clean way to share client code later.

## Consequences

- Each crate can be given to an agent as a clear scope in an issue.
- The `core` → `db` boundary keeps HTTP concerns out of queries and SQL out of handlers.
- Versions are not fixed here. They are chosen when a dependency is added, using the latest release.
