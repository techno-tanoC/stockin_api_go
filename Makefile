SCHEMA_FILE ?= schema.sql
DATABASE_HOST ?= db
DATABASE_PASS ?= pass
export PGPASSWORD = $(DATABASE_PASS)
export PGSSLMODE = disable

seed:
	go run ./cmd/seed

psql:
	psql --host=$(DATABASE_HOST) --user=root dev

psql-test:
	psql --host=$(DATABASE_HOST) --user=root test

create:
	psql --host=$(DATABASE_HOST) --user=root --command "CREATE DATABASE dev"

create-test:
	psql --host=$(DATABASE_HOST) --user=root --command "CREATE DATABASE test"

drop:
	psql --host=$(DATABASE_HOST) --user=root --command "DROP DATABASE dev" | true

drop-test:
	psql --host=$(DATABASE_HOST) --user=root --command "DROP DATABASE test" | true

apply:
	cat schema.sql | psqldef --host $(DATABASE_HOST) --user root dev

apply-test:
	cat schema.sql | psqldef --host $(DATABASE_HOST) --user root test

setup: drop drop-test create create-test apply apply-test

reset: setup seed

migrate:
	psql --host=$(DATABASE_HOST) --user=root --command "CREATE DATABASE prod" | true
	cat schema.sql | psqldef --host $(DATABASE_HOST) --user root prod --dry-run
	cat schema.sql | psqldef --host $(DATABASE_HOST) --user root prod
