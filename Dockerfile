FROM node:24 AS build-frontend
WORKDIR /src

COPY ./frontend/package.json ./frontend/package-lock.json ./
RUN npm ci

COPY ./frontend/ ./
RUN npm run build

#

FROM golang:1.26 AS build-backend
WORKDIR /src

COPY ./go.mod ./go.sum ./
RUN go get ./...

COPY ./ ./
COPY --from=build-frontend /src/dist ./frontend/dist

RUN go build -o /bin/swoptape ./cmd/swoptape
#RUN go build -o /bin/st ./cmd/st

#

FROM alpine:3

COPY --from=build-backend /bin/swoptape /bin/swoptape
#COPY --from=build-backend /bin/st /bin/st

VOLUME [ "/data" ]

ENTRYPOINT [ "/bin/swoptape" ]
