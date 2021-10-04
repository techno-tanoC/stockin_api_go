SCHEMA_FILE = schema.sql
DATABASE_HOST = 0.0.0.0

diff:
	cat $(SCHEMA_FILE) | mysqldef --host=$(DB_HOST) --user=root --password=pass dev --dry-run

apply:
	cat $(SCHEMA_FILE) | mysqldef --host=$(DB_HOST) --user=root --password=pass dev

mysql:
	mysql --host=$(DATABASE_HOST) --user=root --password=pass dev

seed:
	cargo run --bin seed

reset: schema-apply seed
