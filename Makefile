SCHEMA_FILE ?= schema.sql
DATABASE_HOST ?= db
DATABASE_PASS ?= pass

seed:
	go run ./cmd/seed

mysql:
	mysql --host=$(DATABASE_HOST) --user=root --password=$(DATABASE_PASS) dev

create:
	echo "CREATE DATABASE IF NOT EXISTS dev" | mysql --host=$(DATABASE_HOST) --user=root --password=$(DATABASE_PASS)

drop:
	echo "DROP DATABASE IF EXISTS dev" | mysql --host=$(DATABASE_HOST) --user=root --password=$(DATABASE_PASS)

apply:
	cat $(SCHEMA_FILE) | mysqldef --host=$(DATABASE_HOST) --user=root --password=$(DATABASE_PASS) dev

reset: drop create apply seed
