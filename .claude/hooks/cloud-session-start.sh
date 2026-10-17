#!/usr/bin/env bash
# SPDX-License-Identifier: AGPL-3.0-only
# SPDX-FileCopyrightText: 2026 Rareș Nistor

# SessionStart hook: in Claude Code cloud sessions, start the development
# PostgreSQL from dev/swoptape-dev-db and export DATABASE_URL. Does nothing in
# local sessions, where `just local-db` is used instead.

set -euo pipefail

[[ "${CLAUDE_CODE_REMOTE:-}" == "true" ]] || exit 0

compose_file="$CLAUDE_PROJECT_DIR/dev/swoptape-dev-db/docker-compose.yml"

if ! docker info >/dev/null 2>/dev/null; then
    nohup dockerd >/tmp/dockerd.log 2>/tmp/dockerd.err &
    for _ in $(seq 30); do
        docker info >/dev/null 2>/dev/null && break
        sleep 1
    done
fi

docker compose -f "$compose_file" up --detach --quiet-pull

for _ in $(seq 60); do
    docker compose -f "$compose_file" exec -T db pg_isready -U swoptape -d swoptape >/dev/null && break
    sleep 1
done
docker compose -f "$compose_file" exec -T db pg_isready -U swoptape -d swoptape

echo "export DATABASE_URL=postgres://swoptape:swoptape@localhost:5432/swoptape" >> "$CLAUDE_ENV_FILE"
