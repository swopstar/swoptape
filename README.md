<div align="center"><img src="logo.svg" alt="" width="64" height="64"></div>
<h1 align="center">swoptape</h1>
<p align="center">Rip it. Tag it. Stream it.</p>

## Structure

```
cmd/swoptape/   # main package — binary entry point
internal/       # Go packages (not exported)
frontend/       # Vite + React app (uses @swopstar/react-ui)
data/           # database migrations and seed data
local/          # local development environment
```

## Development

### Prerequisites

- Go 1.26+
- Node.js (for the frontend)
- See `local/swoptape-dev/` for the local environment setup

### Backend

```bash
go run ./cmd/swoptape
```

### Frontend

```bash
cd frontend
npm install
npm run dev
```

## UI library

The frontend consumes [`@swopstar/react-ui`](../ui), the shared component library. See its README for setup and theming. The swoptape brand seed colour is `#DE00EE`.
