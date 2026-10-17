# SPDX-License-Identifier: AGPL-3.0-only
# SPDX-FileCopyrightText: 2026 Rareș Nistor

serve:
    cargo run --bin swoptape

build:
    cargo build --release --locked --bin swoptape

test *args:
    cargo test --workspace --locked {{args}}

fmt:
    cargo fmt --all

check:
    cargo fmt --all --check
    CARGO_BUILD_WARNINGS=deny cargo clippy --workspace --all-targets --locked
    cargo deny check
