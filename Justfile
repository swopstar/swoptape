# SPDX-License-Identifier: AGPL-3.0-only
# SPDX-FileCopyrightText: 2026 Rareș Nistor

mod go
mod web

import 'version.just'

image := "swoptape"

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

notices: (go::notices) (web::notices)
    node scripts/combine-notices.mjs

#
# Container
#

container tag="dev":
    docker build \
        --build-arg version_year={{version_year}} \
        --build-arg version_month={{version_month}} \
        --build-arg version_release={{version_release}} \
        --build-arg version_patch={{version_patch}} \
        --build-arg version_pre={{version_pre}} \
        --build-arg version_branch={{version_branch}} \
        --build-arg version_commit={{version_commit}} \
        --build-arg source_url={{source_url}} \
        -t {{image}}:{{tag}} \
        .

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
