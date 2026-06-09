# SPDX-License-Identifier: AGPL-3.0-only
# SPDX-FileCopyrightText: 2026 Rareș Nistor

generate:
    npm -C frontend run generate

serve:
    npm -C frontend run dev

setup:
    npm -C frontend ci

build:
    npm -C frontend run build

fmt:
    npm -C frontend run fmt

test:
    npm -C frontend run test

check:
    npm -C frontend run lint
