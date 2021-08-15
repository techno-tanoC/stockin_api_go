SCHEMA_FILE ?= schema.sql
DATABASE_FILE ?= stockin.sqlite3

schema-diff:
	sqlite3def -f $(SCHEMA_FILE) --dry-run $(DATABASE_FILE)

schema-apply:
	sqlite3def -f $(SCHEMA_FILE) $(DATABASE_FILE)

sqlite:
	sqlite3 $(DATABASE_FILE)
