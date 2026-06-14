FROM node:24 AS build-frontend

RUN mkdir /src
COPY ./frontend/package.json ./frontend/package-lock.json /src/

WORKDIR /src
RUN npm ci

COPY ./frontend/ /src/
COPY ./www/api/v0/openapi.yaml /www/api/v0/openapi.yaml
RUN npm run generate
RUN npm run build

#

FROM golang:1.26-alpine AS build-backend
WORKDIR /src

COPY ./go.mod ./go.sum /src/
RUN go get ./...

COPY ./ ./
COPY --from=build-frontend /src/dist /src/frontend/dist
RUN rm /src/frontend/gen.go

ARG version_year
ARG version_month
ARG version_release
ARG version_patch
ARG version_pre
ARG version_branch
ARG version_commit
ARG source_url

RUN go generate ./...
RUN go build \
    -ldflags "-X 'github.com/swopstar/gokit/ver.year=${version_year}' -X 'github.com/swopstar/gokit/ver.month=${version_month}' -X 'github.com/swopstar/gokit/ver.release=${version_release}' -X 'github.com/swopstar/gokit/ver.patch=${version_patch}' -X 'github.com/swopstar/gokit/ver.preRelease=${version_pre}' -X 'github.com/swopstar/gokit/ver.branch=${version_branch}' -X 'github.com/swopstar/gokit/ver.commit=${version_commit}' -X 'github.com/swopstar/gokit/ver.source=${source_url}'" \
    -o /bin/swoptape \
    ./cmd/swoptape

#

FROM alpine:3

RUN apk update && apk add --no-cache \
    ca-certificates tzdata

COPY --from=build-backend /bin/swoptape /bin/swoptape
#COPY --from=build-backend /bin/st /bin/st

ENV SWOPTAPE_DATA=/data
VOLUME [ "/data" ]

ENTRYPOINT [ "/bin/swoptape" ]
