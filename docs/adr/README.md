# Architecture decision records

Suite-wide decisions and standards live in [swopstar/meta](https://github.com/swopstar/meta): ADRs in [`adr/`](https://github.com/swopstar/meta/tree/main/adr) and conventions in [`standards/`](https://github.com/swopstar/meta/tree/main/standards). Decisions specific to swoptape live here, numbered `ADR-T###`.

A swoptape ADR may depart from a suite-wide one. When it does, it names the suite-wide ADR and explains why.

Agents must follow `accepted` records and treat `proposed` ones as the default. A change that contradicts a record must update or supersede that record in the same PR, and say so in the PR description.

New records copy [`TEMPLATE.md`](TEMPLATE.md) and take the next free number.

| ADR                                      | Decision                                           | Status   |
| ---------------------------------------- | -------------------------------------------------- | -------- |
| [T001](T001-rust-rewrite.md)             | Rewrite on a long-lived branch                     | accepted |
| [T002](T002-test-first.md)               | Test-first development                             | accepted |
| [T003](T003-agentic-workflow.md)         | Agentic development workflow                       | accepted |
| [T004](T004-resource-budget.md)          | Resource budget                                    | proposed |
| [T005](T005-workspace-and-stack.md)      | Workspace layout and project dependencies          | proposed |
| [T006](T006-dual-mode-library.md)        | Dual-mode library model                            | proposed |
| [T007](T007-subsonic-auth.md)            | Subsonic authentication via app passwords          | proposed |
| [T008](T008-code-first-openapi.md)       | Code-first OpenAPI contract                        | proposed |
| [T009](T009-shared-client-core.md)       | Native clients share a Rust core                   | proposed |
