SCHEMA_FILE ?= schema.sql
DATABASE_HOST ?= db
DATABASE_PASS ?= pass
export PGPASSWORD = $(DATABASE_PASS)
export PGSSLMODE = disable

start:

seed:

psql:
	psql --host=$(DATABASE_HOST) --user=root dev

create:
	psql --host=$(DATABASE_HOST) --user=root --command "CREATE DATABASE dev"

drop:
	psql --host=$(DATABASE_HOST) --user=root --command "DROP DATABASE dev" | true

apply:
	cat schema.sql | psqldef --host $(DATABASE_HOST) --user root dev

setup: drop create apply

reset: setup seed
