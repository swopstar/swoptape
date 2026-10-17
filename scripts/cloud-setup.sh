#!/usr/bin/env bash
# SPDX-License-Identifier: AGPL-3.0-only
# SPDX-FileCopyrightText: 2026 Rareș Nistor

# Setup script for swoptape's Claude Code cloud environment. Paste this file's
# contents into the environment's "Setup script" field at claude.ai/code. It
# runs as root on Ubuntu 24.04 before Claude starts, and its result is cached.
# Keep the pinned versions in step with rust-toolchain.toml and
# dev/swoptape-dev-db/docker-compose.yml.

set -euo pipefail

RUST_TOOLCHAIN=1.99
JUST_VERSION=1.59.0
CARGO_DENY_VERSION=0.20.2
REUSE_VERSION=6.2.0
POSTGRES_IMAGE=postgres:18-alpine

install_release_binary() {
    local url=$1 binary=$2 workdir
    workdir=$(mktemp -d)
    curl --fail --silent --show-error --location "$url" | tar -xz -C "$workdir"
    install -m 0755 "$(find "$workdir" -type f -name "$binary" | head -n 1)" "/usr/local/bin/$binary"
    rm -rf "$workdir"
}

start_docker() {
    docker info >/dev/null 2>/dev/null && return
    nohup dockerd >/tmp/dockerd.log 2>/tmp/dockerd.err &
    for _ in $(seq 30); do
        docker info >/dev/null 2>/dev/null && return
        sleep 1
    done
    echo "dockerd did not start; see /tmp/dockerd.err" >&2
    return 1
}

rustup toolchain install "$RUST_TOOLCHAIN" --profile minimal --component clippy,rustfmt

install_release_binary \
    "https://github.com/casey/just/releases/download/${JUST_VERSION}/just-${JUST_VERSION}-x86_64-unknown-linux-musl.tar.gz" \
    just
install_release_binary \
    "https://github.com/EmbarkStudios/cargo-deny/releases/download/${CARGO_DENY_VERSION}/cargo-deny-${CARGO_DENY_VERSION}-x86_64-unknown-linux-musl.tar.gz" \
    cargo-deny

python3 -m pip install --quiet --break-system-packages "reuse==${REUSE_VERSION}"

start_docker
docker pull --quiet "$POSTGRES_IMAGE"
