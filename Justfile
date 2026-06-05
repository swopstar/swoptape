version := "2606.1.0"

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
    go build -o tmp/bin/swoptape {{daemon_package}}

build-backend-cross-all: (build-backend-cross "linux" "amd64") (build-backend-cross "linux" "arm64") (build-backend-cross "linux" "riscv64")

build-backend-cross platform arch: prebuild-backend
    mkdir -p dist/bin
    env GOOS={{platform}} GOARCH={{arch}} go build -o dist/bin/swoptape-{{platform}}-{{arch}} {{daemon_package}}

#
# Frontend
#

serve-frontend:
    npm -C frontend run dev

setup-frontend:
    npm -C frontend ci

build-frontend:
    npm -C frontend run build
