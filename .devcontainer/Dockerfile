# syntax=docker/dockerfile:1.4
FROM golang:1.18.3
WORKDIR /workspace

RUN <<EOF
  apt update
  apt install -y wget make default-mysql-client git
EOF

RUN <<EOF
  wget https://github.com/k0kubun/sqldef/releases/download/v0.11.59/mysqldef_linux_amd64.tar.gz
  tar -C /usr/local/bin -xvf mysqldef_linux_amd64.tar.gz
  rm mysqldef_linux_amd64.tar.gz
EOF

RUN <<EOF
  go install github.com/cosmtrek/air@latest

  go install github.com/volatiletech/sqlboiler/v4@latest
  go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest

  go install golang.org/x/tools/gopls@latest
EOF
