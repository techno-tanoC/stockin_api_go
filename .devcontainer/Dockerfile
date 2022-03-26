# syntax=docker/dockerfile:1.3-labs
FROM rust:1.56-slim-bullseye

RUN cargo install cargo-watch

RUN <<EOF
  apt update
  apt install -y wget make default-mysql-client
EOF

RUN <<EOF
  wget https://github.com/k0kubun/sqldef/releases/latest/download/mysqldef_linux_amd64.tar.gz
  tar -C /usr/local/bin -xvf mysqldef_linux_amd64.tar.gz
  rm mysqldef_linux_amd64.tar.gz
EOF
