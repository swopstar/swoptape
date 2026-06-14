# SPDX-License-Identifier: AGPL-3.0-only
# SPDX-FileCopyrightText: 2026 Rareș Nistor

import 'version.just'

daemon_package := "./cmd/swoptape"
cli_package    := './cmd/st'

serve:
    air

generate:
    go generate ./...

build:
    mkdir -p tmp/bin
    go build \
        -ldflags "-X 'github.com/swopstar/gokit/ver.year={{version_year}}' -X 'github.com/swopstar/gokit/ver.month={{version_month}}' -X 'github.com/swopstar/gokit/ver.release={{version_release}}' -X 'github.com/swopstar/gokit/ver.patch={{version_patch}}' -X 'github.com/swopstar/gokit/ver.preRelease={{version_pre}}' -X 'github.com/swopstar/gokit/ver.branch={{version_branch}}' -X 'github.com/swopstar/gokit/ver.commit={{version_commit}}' -X 'github.com/swopstar/gokit/ver.source={{source_url}}'" \
        -o tmp/bin/swoptape \
        {{daemon_package}}

fmt:
    go fmt ./...

test *args='./...':
    go test {{args}}

check:
    golangci-lint run
    go mod verify

check-tidy:
    go mod tidy
    git diff --exit-code go.mod go.sum

check-vuln:
    govulncheck ./...

tidy:
    go mod tidy

notices:
    mkdir -p tmp/notices/go-texts
    go list -m all > tmp/notices/go-versions.txt
    go-licenses report ./... 2>/dev/null | grep -v '^github.com/swopstar/' > tmp/notices/go.csv
    go-licenses save ./... --save_path=tmp/notices/go-texts --force \
        --ignore github.com/swopstar/swoptape \
        --ignore github.com/swopstar/gokit \
        2>/dev/null; true

update-gokit:
    go get github.com/swopstar/gokit@latest
    go mod tidy
