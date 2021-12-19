# syntax=docker/dockerfile:1.3-labs
FROM rust:1.56-slim-bullseye as build
WORKDIR /work

COPY Cargo.toml Cargo.lock ./
RUN <<EOF
  mkdir -p src/bin
  echo "fn main() {}" > src/bin/main.rs
  cargo build --release
EOF

ENV SQLX_OFFLINE true

COPY . .
RUN <<EOF
  touch src/bin/main.rs
  cargo build --release
EOF


FROM debian:bullseye-slim
WORKDIR /app

RUN apt-get update && apt-get install -y wget make default-mysql-client

ENV DOCKERIZE_VERSION v0.6.1
RUN <<EOF
  wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz
  tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz
  rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz
EOF

RUN <<EOF
  wget https://github.com/k0kubun/sqldef/releases/latest/download/mysqldef_linux_amd64.tar.gz
  tar -C /usr/local/bin -xvf mysqldef_linux_amd64.tar.gz
  rm mysqldef_linux_amd64.tar.gz
EOF

COPY Makefile schema.sql ./
COPY --from=build /work/target/release/main /usr/local/bin/main
CMD ["/usr/local/bin/main"]
