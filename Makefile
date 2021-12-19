SCHEMA_FILE = schema.sql
DATABASE_HOST ?= 0.0.0.0
DATABASE_PASS ?= pass

start:
	cargo watch -x run

test:
	cargo watch -s "cargo test -- --test-threads=1"

lint:
	cargo fmt --all -- --check
	cargo clippy --all-targets --all-features -- -D warnings

fmt:
	cargo fmt --all
	cargo clippy --all-targets --all-features --fix --allow-dirty

seed:
	cargo run --bin seed

dump:
	cargo run --bin dump

mysql:
	mysql --host=$(DATABASE_HOST) --user=root --password=$(DATABASE_PASS) dev

mysql-test:
	mysql --host=$(DATABASE_HOST) --user=root --password=$(DATABASE_PASS) test

create:
	echo "CREATE DATABASE IF NOT EXISTS dev" | mysql --host=$(DATABASE_HOST) --user=root --password=$(DATABASE_PASS)

create-test:
	echo "CREATE DATABASE IF NOT EXISTS test" | mysql --host=$(DATABASE_HOST) --user=root --password=$(DATABASE_PASS)

diff:
	cat $(SCHEMA_FILE) | mysqldef --host=$(DATABASE_HOST) --user=root --password=$(DATABASE_PASS) dev --dry-run

apply:
	cat $(SCHEMA_FILE) | mysqldef --host=$(DATABASE_HOST) --user=root --password=$(DATABASE_PASS) dev

apply-test:
	cat $(SCHEMA_FILE) | mysqldef --host=$(DATABASE_HOST) --user=root --password=$(DATABASE_PASS) test

drop:
	echo "DROP DATABASE IF EXISTS dev" | mysql --host=$(DATABASE_HOST) --user=root --password=$(DATABASE_PASS)
	echo "DROP DATABASE IF EXISTS test" | mysql --host=$(DATABASE_HOST) --user=root --password=$(DATABASE_PASS)

reset: drop create create-test apply apply-test seed

setup-test: create-test apply-test

migrate:
	echo "CREATE DATABASE IF NOT EXISTS prod" | mysql --host=$(DATABASE_HOST) --user=root --password=$(DATABASE_PASS)
	cat $(SCHEMA_FILE) | mysqldef --host=$(DATABASE_HOST) --user=root --password=$(DATABASE_PASS) prod --dry-run
	cat $(SCHEMA_FILE) | mysqldef --host=$(DATABASE_HOST) --user=root --password=$(DATABASE_PASS) prod
