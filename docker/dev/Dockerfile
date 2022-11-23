# syntax=docker/dockerfile:1.4
FROM golang:1.19.3
WORKDIR /workspace

RUN <<EOF
  apt update
  apt install -y wget make git postgresql-client locales-all
EOF

RUN <<EOF
  wget https://github.com/k0kubun/sqldef/releases/download/v0.11.59/psqldef_linux_amd64.tar.gz
  tar -C /usr/local/bin -xvf psqldef_linux_amd64.tar.gz
  rm psqldef_linux_amd64.tar.gz
EOF

RUN <<EOF
  go install github.com/cosmtrek/air@latest

  go install github.com/volatiletech/sqlboiler/v4@v4.13.0
  go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@v4.13.0

  go install github.com/kisielk/errcheck@latest
  go install golang.org/x/tools/gopls@latest
EOF
