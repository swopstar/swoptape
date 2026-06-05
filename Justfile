version_year    := `date +%y`
version_month   := `date +%m`
version_release := "0"
version_patch   := "0"
version_pre     := "dev"
version_branch  := `git rev-parse --abbrev-ref HEAD`
version_commit  := `git rev-parse --short HEAD`

daemon_package := "./cmd/swoptape"
cli_package    := './cmd/st'

[parallel]
serve: serve-backend serve-frontend

setup: setup-frontend

#
# Backend
#

serve-backend:
    air

generate-backend:
    go generate ./...

build-backend:
    mkdir -p tmp/bin
    go build \
        -ldflags "-X 'github.com/swopstar/gokit/ver.year={{version_year}}' -X 'github.com/swopstar/gokit/ver.month={{version_month}}' -X 'github.com/swopstar/gokit/ver.release={{version_release}}' -X 'github.com/swopstar/gokit/ver.patch={{version_patch}}' -X 'github.com/swopstar/gokit/ver.preRelease={{version_pre}}' -X 'github.com/swopstar/gokit/ver.branch={{version_branch}}' -X 'github.com/swopstar/gokit/ver.commit={{version_commit}}'" \
        -o tmp/bin/swoptape \
        {{daemon_package}}

#
# Frontend
#

serve-frontend:
    npm -C frontend run dev

setup-frontend:
    npm -C frontend ci

build-frontend:
    npm -C frontend run build

#
# Database
#

local-db *args='up -d':
    docker compose -f dev/swoptape-dev-db/docker-compose.yml {{args}}

#
# Misc
#

update-gokit:
    go get github.com/swopstar/gokit@latest
