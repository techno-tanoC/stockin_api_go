SCHEMA_FILE = schema.sql
DATABASE_HOST ?= 0.0.0.0

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
	mysql --host=$(DATABASE_HOST) --user=root --password=pass dev

mysql-test:
	mysql --host=$(DATABASE_HOST) --user=root --password=pass test

create:
	echo "CREATE DATABASE IF NOT EXISTS dev" | mysql --host=$(DATABASE_HOST) --user=root --password=pass

create-test:
	echo "CREATE DATABASE IF NOT EXISTS test" | mysql --host=$(DATABASE_HOST) --user=root --password=pass

diff:
	cat $(SCHEMA_FILE) | mysqldef --host=$(DATABASE_HOST) --user=root --password=pass dev --dry-run

apply:
	cat $(SCHEMA_FILE) | mysqldef --host=$(DATABASE_HOST) --user=root --password=pass dev

apply-test:
	cat $(SCHEMA_FILE) | mysqldef --host=$(DATABASE_HOST) --user=root --password=pass test

drop:
	echo "DROP DATABASE IF EXISTS dev" | mysql --host=$(DATABASE_HOST) --user=root --password=pass
	echo "DROP DATABASE IF EXISTS test" | mysql --host=$(DATABASE_HOST) --user=root --password=pass

reset: drop create create-test apply apply-test seed

setup-test: create-test apply-test
