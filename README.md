<div align="center"><img src="logo.svg" alt="" width="64" height="64"></div>
<h1 align="center">swoptape</h1>
<p align="center">Rip it. Tag it. Stream it.</p>

## Structure

```
crates/      # Rust workspace: core, db, subsonic, server
frontend/    # Vite + React web player (uses @swopstar/react-ui)
docs/        # overview, architecture, roadmap, ADRs and designs (suite-wide ones are in swopstar/meta)
dev/         # local development environment
```

The backend is being rewritten in Rust on the `rewrite/rust` branch. See [ADR-T001](docs/adr/T001-rust-rewrite.md).

## Development

### Prerequisites

- Rust (the toolchain version is pinned in `rust-toolchain.toml` and installed by `rustup` automatically)
- [`just`](https://github.com/casey/just), [`reuse`](https://reuse.software/) and [`cargo-deny`](https://github.com/EmbarkStudios/cargo-deny)
- Node.js, for the frontend
- Docker, for the local PostgreSQL instance

### Common tasks

```bash
just local-db   # start PostgreSQL
just serve      # run the server and the frontend dev server
just test       # run the tests
just check      # lint, licence and dependency checks
```

## UI library

The frontend consumes [`@swopstar/react-ui`](../ui), the shared component library. See its README for setup and theming. The swoptape brand seed colour is `#DE00EE`.
