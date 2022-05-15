# syntax=docker/dockerfile:1.4
FROM rust:1.59-slim-bullseye

RUN cargo install cargo-watch

RUN <<EOF
  apt update
  apt install -y wget make default-mysql-client git
EOF

RUN <<EOF
  wget https://github.com/k0kubun/sqldef/releases/download/v0.11.45/mysqldef_linux_amd64.tar.gz
  tar -C /usr/local/bin -xvf mysqldef_linux_amd64.tar.gz
  rm mysqldef_linux_amd64.tar.gz
EOF
