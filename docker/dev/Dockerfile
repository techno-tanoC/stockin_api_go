# syntax=docker/dockerfile:1.3-labs
FROM rust:1.56-slim-bullseye
WORKDIR /work

RUN cargo install cargo-watch

COPY Cargo.toml Cargo.lock ./
RUN <<EOF
  mkdir -p src/bin
  echo "fn main() {}" > src/bin/main.rs
  cargo check
  cargo build
EOF
