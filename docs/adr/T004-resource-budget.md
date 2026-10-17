# ADR-T004: Resource budget

**Status:** proposed
**Date:** 2026-10-17

## Context

[ADR-0000](https://github.com/swopstar/meta/blob/main/adr/0000-stack.md) and [ADR-0001](https://github.com/swopstar/meta/blob/main/adr/0001-rust.md) target an entry-level NAS: a low-power ARM64 CPU and 1 GB of RAM, shared with PostgreSQL and whatever else the user runs. Resource use creeps up unless it is measured.

## Decision

Targets for the swoptape process, excluding PostgreSQL:

| Measure                                  | Target                                   |
| ---------------------------------------- | ---------------------------------------- |
| Idle resident memory, 100k-track library | under 30 MB                              |
| CPU when idle                            | effectively zero, with no polling loops  |
| Release binary, frontend embedded        | under 20 MB                              |
| Streaming without transcoding            | file served directly, no decoding        |
| Supported targets                        | `x86_64` and `aarch64` Linux, static musl |

- Transcoding is opt-in per request and delegated to an external `ffmpeg` process, so it costs nothing when unused.
- Library scans are incremental. They compare modification time and size before reading tags.
- Background work (scans, watching for changes) is bounded and can be switched off.
- Static musl builds replace musl's default allocator, as ADR-0001 requires. Whether that is mimalloc or jemalloc is decided by measuring both against this budget when release builds are set up.

## Alternatives considered

- **No explicit budget.** Simpler, but regressions go unnoticed until users on small hardware report them.

## Consequences

- A benchmark job should eventually track memory and binary size so that regressions show up in CI.
- Dependencies are weighed against binary size and memory, not only convenience.
- PostgreSQL will dominate the idle footprint. The deployment docs should recommend PostgreSQL settings tuned for small machines.
