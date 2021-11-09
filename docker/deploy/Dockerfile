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

COPY --from=build /work/target/release/main /usr/local/bin/main
CMD ["/usr/local/bin/main"]
