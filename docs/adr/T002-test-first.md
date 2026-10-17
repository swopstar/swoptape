# ADR-T002: Test-first development

**Status:** accepted
**Date:** 2026-10-17

## Context

Most changes will be written by agents and reviewed by both agents and a human. Reviewers need a precise, executable statement of what a change is supposed to do. Tests that encode behaviour give that, and they keep holding as the code underneath is refactored.

## Decision

- Changes are developed test-first: write a failing test that captures the behaviour, then make it pass.
- Tests exercise public interfaces. Prefer integration tests in a crate's `tests/` directory, which can only see the public API, over `#[cfg(test)]` unit tests of private functions.
- Test behaviour, not implementation. A refactor that keeps observable behaviour must not break tests.
- Prefer table-driven tests where they reduce repetition.
- Database-backed code is tested against real PostgreSQL: `just local-db` locally, and a service container in CI. HTTP handlers are tested through the router.
- The Subsonic API is covered by conformance tests that issue the request shapes real clients send against a fixture library and check the responses in both XML and JSON.

## Alternatives considered

- **Tests after the initial build-out.** Faster to start, but agents would have no executable definition of done while the core is being built, which is when it matters most.
- **Mocked database in tests.** Faster, but hides query and migration bugs, which are a large share of the risk in a library server.

## Consequences

- Every PR that changes behaviour includes the tests that specify it. Reviewers check that the tests would have failed without the change.
- An issue's acceptance criteria should be written so that they translate directly into tests.
- A small fixture music library (freely licensed or generated audio with varied tags and folder layouts) is needed early on.
