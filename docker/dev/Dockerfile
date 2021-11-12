# syntax=docker/dockerfile:1.3-labs
FROM rust:1.56-slim-bullseye
WORKDIR /work

RUN cargo install cargo-watch

# mysql client
RUN <<EOF
  apt update
  apt install -y wget make default-mysql-client
EOF

# mysqldef
RUN <<EOF
  wget -O mysqldef.tar.gz https://github.com/k0kubun/sqldef/releases/download/v0.11.13/mysqldef_linux_amd64.tar.gz
  tar -C /usr/local/bin -xzvf mysqldef.tar.gz
  rm mysqldef.tar.gz
EOF

# dockerize
ENV DOCKERIZE_VERSION v0.6.1
RUN <<EOF
  wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz
  tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz
  rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz
EOF

COPY Cargo.toml Cargo.lock ./
RUN <<EOF
  mkdir -p src/bin
  echo "fn main() {}" > src/bin/main.rs
  cargo check
  cargo build
EOF
