SCHEMA_FILE ?= schema.sql
DATABASE_HOST ?= db
DATABASE_PASS ?= pass

seed:
	go run ./cmd/seed

mysql:
	mysql --host=$(DATABASE_HOST) --user=root --password=$(DATABASE_PASS) dev

mysql-test:
	mysql --host=$(DATABASE_HOST) --user=root --password=$(DATABASE_PASS) test

create:
	echo "CREATE DATABASE IF NOT EXISTS dev" | mysql --host=$(DATABASE_HOST) --user=root --password=$(DATABASE_PASS)

create-test:
	echo "CREATE DATABASE IF NOT EXISTS test" | mysql --host=$(DATABASE_HOST) --user=root --password=$(DATABASE_PASS)

drop:
	echo "DROP DATABASE IF EXISTS dev" | mysql --host=$(DATABASE_HOST) --user=root --password=$(DATABASE_PASS)

drop-test:
	echo "DROP DATABASE IF EXISTS test" | mysql --host=$(DATABASE_HOST) --user=root --password=$(DATABASE_PASS)

apply:
	cat $(SCHEMA_FILE) | mysqldef --host=$(DATABASE_HOST) --user=root --password=$(DATABASE_PASS) dev

apply-test:
	cat $(SCHEMA_FILE) | mysqldef --host=$(DATABASE_HOST) --user=root --password=$(DATABASE_PASS) test

reset: drop drop-test create create-test apply apply-test seed
