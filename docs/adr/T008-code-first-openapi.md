# ADR-T008: Code-first OpenAPI contract

**Status:** proposed
**Date:** 2026-10-17

## Context

[ADR-0000](https://github.com/swopstar/meta/blob/main/adr/0000-stack.md) notes that the frontend and backend API schemas are kept in sync by hand, and that a better mechanism is future work. The first-party API (`/api`) is consumed by the web player now and by native clients later. Agents working on the server and on clients must not let them drift apart.

## Decision

- The OpenAPI document is generated from the Rust handlers and types with `utoipa`.
- The generated document is committed to the repository. CI regenerates it and fails if the committed copy differs.
- The TypeScript client for the web player is generated from it (orval). The Swift client for iOS is generated from it later.
- The API is versioned as set out in the [versioning standard](https://github.com/swopstar/meta/blob/main/standards/versioning.md): the major version in the path (`/api/v1/`), independent of the application version.
- The Go implementation's `openapi.yaml` on `main` is a reference for resource design, not a contract to match.

## Alternatives considered

- **Spec-first, generating Rust from a hand-written OpenAPI document.** The Rust generators produce whole servers or awkward types rather than handlers that fit `axum`. This is the same problem ADR-0000 hit with Go.
- **Keep schemas in sync by hand.** The status quo, and the main source of drift.

## Consequences

- A server change that alters the API shows up as a diff to the committed document, which is easy to review.
- Client generation is a build step that depends on the server crate.
