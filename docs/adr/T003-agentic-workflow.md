# ADR-T003: Agentic development workflow

**Status:** accepted
**Date:** 2026-10-17

## Context

The rewrite is also an exercise in agentic development. Agents implement changes and other agents review them, while the maintainer stays the final reviewer. This only works if each unit of work is small, specified up front, and checked mechanically before a human looks at it.

## Decision

- **Issues are specifications.** Each issue states the goal, links the relevant ADRs, lists acceptance criteria and names the crates or paths in scope. Issues use the task template.
- **One issue, one PR.** An implementing agent works on its own branch (`<issue-number>-<slug>`, off `rewrite/rust`) and opens a PR that closes exactly one issue.
- **Agent review.** A separate reviewing agent reviews every PR and posts its findings as PR comments. The implementing agent addresses them.
- **Separate identity.** Agents push branches, open PRs and post reviews as the `rz-aurora-bot` account, which has write access but no admin rights. They never use the maintainer's credentials, with one accepted exception: Claude Code cloud sessions act as the maintainer's GitHub account, because the cloud's GitHub proxy uses the account connected to the maintainer's Claude account. The maintainer merges those PRs through the PR-only bypass, since GitHub doesn't let anyone approve their own PR.
- **Human merge.** The maintainer reviews and merges every PR. Branch protection on `rewrite/rust` requires passing CI and the maintainer's approval.
- **CI is the gate.** Formatting, clippy with warnings denied, tests, `cargo deny` and `reuse lint` must all pass. `just check` and `just test` run the same checks locally.
- **Decisions are written down.** Anything that changes architecture or contradicts an ADR updates the ADR in the same PR.

## Alternatives considered

- **Agents work directly with the maintainer in one session, without issues or agent review.** Less ceremony, but it doesn't scale to several agents working in parallel, and every review falls on the maintainer.

## Consequences

- Issues take more effort to write, but PRs become faster to review.
- PRs should stay small. As a guide, a PR that spans several concerns or runs to many hundreds of lines outside tests should be split.
- `CLAUDE.md` holds the operational instructions agents follow. ADRs hold the reasoning.
- If the workflow proves itself here, it is a candidate for a suite-wide ADR in swopstar/meta.
