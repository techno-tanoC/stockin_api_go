SCHEMA_FILE ?= schema.sql
DATABASE_FILE ?= stockin.sqlite3
DATABASE_URL ?= sqlite://$(DATABASE_FILE)

schema-diff:
	sqlite3def -f $(SCHEMA_FILE) --dry-run $(DATABASE_FILE)

schema-apply:
	sqlite3def -f $(SCHEMA_FILE) $(DATABASE_FILE)

sqlite:
	sqlite3 $(DATABASE_FILE)

delete:
	rm -f $(DATABASE_FILE)*

seed:
	cargo run --bin seed

reset: delete schema-apply seed
