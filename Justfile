# SPDX-License-Identifier: AGPL-3.0-only
# SPDX-FileCopyrightText: 2026 Rareș Nistor

mod go
mod web

[parallel]
serve: (go::serve) (web::serve)

setup: setup-precommit (web::setup)
setup-ci: (web::setup)

build: (web::build) (go::build)

[parallel]
fmt: (go::fmt) (web::fmt)

check: (go::check) (web::check) (go::check-tidy)
    reuse lint

[parallel]
test: (go::test) (web::test)

tidy: (go::tidy)

#
# Database
#

local-db *args='up -d':
    docker compose -f dev/swoptape-dev-db/docker-compose.yml {{args}}

#
# Misc
#

setup-precommit:
    pre-commit install
