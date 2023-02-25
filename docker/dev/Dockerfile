# syntax=docker/dockerfile:1.4
FROM golang:1.20.1
WORKDIR /workspace

RUN <<EOF
  groupadd --gid 1000 vscode
  useradd --shell /bin/bash --gid 1000 --uid 1000 -m vscode
EOF

RUN <<EOF
  apt update
  apt install -y wget make git postgresql-client
EOF

RUN <<EOF
  wget https://github.com/k0kubun/sqldef/releases/download/v0.11.59/psqldef_linux_amd64.tar.gz
  tar -C /usr/local/bin -xvf psqldef_linux_amd64.tar.gz
  rm psqldef_linux_amd64.tar.gz
EOF

RUN <<EOF
  go install github.com/cosmtrek/air@latest
  go install github.com/kisielk/errcheck@latest
  go install golang.org/x/tools/gopls@latest
EOF

RUN <<EOF
  chown -R vscode:vscode /go
EOF
