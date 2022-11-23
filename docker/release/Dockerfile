# syntax=docker/dockerfile:1.4
FROM golang:1.19.3 AS builder
WORKDIR /workspace

COPY go.mod go.sum ./
RUN <<EOF
  go mod download
EOF

COPY . .
RUN <<EOF
  CGO_ENABLED=0 go build ./cmd/api/
EOF

FROM debian:bullseye-slim
WORKDIR /workspace

RUN <<EOF
  apt update
  apt install -y wget make postgresql-client locales-all
EOF

RUN <<EOF
  wget https://github.com/k0kubun/sqldef/releases/download/v0.11.59/psqldef_linux_amd64.tar.gz
  tar -C /usr/local/bin -xvf psqldef_linux_amd64.tar.gz
  rm psqldef_linux_amd64.tar.gz
EOF

ENV DOCKERIZE_VERSION v0.6.1
RUN <<EOF
  wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz
  tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz
  rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz
EOF

COPY schema.sql Makefile ./
COPY --from=builder /workspace/api /workspace/app

CMD ["/workspace/app"]
