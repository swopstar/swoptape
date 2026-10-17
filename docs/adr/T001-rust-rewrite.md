# ADR-T001: Rewrite on a long-lived branch

**Status:** accepted
**Date:** 2026-10-17

## Context

[ADR-0001](https://github.com/swopstar/meta/blob/main/adr/0001-rust.md) moves the suite's backends to Rust, and swoptape goes first. The Go backend had a working identity service (users, sessions, application tokens), but the library, Subsonic and API handlers were stubs. Very little working code is lost by starting again, but the Go code is still useful as a reference.

## Decision

- The rewrite happens on the long-lived branch `rewrite/rust`. Feature branches branch off it and merge back into it.
- `main` keeps the Go implementation until the MVP is reached. The rewrite branch then replaces `main`.
- The Go code on `main` is reference material, not a porting target. The identity design, the OpenAPI resource layout (`www/api/v0/openapi.yaml`) and the Subsonic 1.16.1 response types (`www/subsonic/response/`) are worth consulting, using `git show main:<path>`.
- The React frontend (`frontend/`) and `@swopstar/react-ui` are kept.

## Alternatives considered

- **Rewrite on `main`.** Simpler branching, but `main` would have no working build until the MVP.
- **Port the Go code file by file.** Keeps behaviour identical, but carries Go idioms into Rust, and there is little working behaviour worth preserving.
- **A new repository.** Clean history, but splits issues and history across two repositories for the same product.

## Consequences

- Nothing ships from the rewrite until the MVP. Container builds on `main` keep publishing the Go version in the meantime.
- `main` should be treated as frozen. Fixes there are not ported to the branch.
- PRs target `rewrite/rust`, not `main`.
- Links to this repository's ADRs from swopstar/meta point at `rewrite/rust` until it replaces `main`.
