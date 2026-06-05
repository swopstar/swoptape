version_year    := `date +%y`
version_month   := `date +%m`
version_release := "0"
version_patch   := "0"
version_pre     := "dev"
version_branch  := `git rev-parse --abbrev-ref HEAD`
version_commit  := `git rev-parse --short HEAD`

daemon_package := "./cmd/swoptape"
cli_package    := './cmd/st'

serve:
    air

generate:
    go generate ./...

build:
    mkdir -p tmp/bin
    go build \
        -ldflags "-X 'github.com/swopstar/gokit/ver.year={{version_year}}' -X 'github.com/swopstar/gokit/ver.month={{version_month}}' -X 'github.com/swopstar/gokit/ver.release={{version_release}}' -X 'github.com/swopstar/gokit/ver.patch={{version_patch}}' -X 'github.com/swopstar/gokit/ver.preRelease={{version_pre}}' -X 'github.com/swopstar/gokit/ver.branch={{version_branch}}' -X 'github.com/swopstar/gokit/ver.commit={{version_commit}}'" \
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

update-gokit:
    go get github.com/swopstar/gokit@latest
    go mod tidy
