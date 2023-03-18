package internal

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofrs/uuid"
)

func WithTestDatabase(ctx context.Context, base string, path string) (*sql.DB, func()) {
	uuid := uuid.Must(uuid.NewV4()).String()
	dbname := fmt.Sprintf("test_%s", strings.ReplaceAll(uuid, "-", "_"))

	err := createTestDatabase(ctx, base, dbname)
	if err != nil {
		log.Fatalf("create test database error %v", err)
	}

	database := fmt.Sprintf("%s dbname=%s", base, dbname)
	db, err := sql.Open("postgres", database)
	if err != nil {
		_ = dropTestDatabase(ctx, base, dbname)
		log.Fatalf("open database error %v", err)
	}

	err = createSchema(ctx, db, path)
	if err != nil {
		_ = db.Close()
		_ = dropTestDatabase(ctx, base, dbname)
		log.Fatalf("create schema error %v", err)
	}

	return db, func() {
		_ = db.Close()
		err = dropTestDatabase(ctx, base, dbname)
		if err != nil {
			log.Fatalf("drop database error %v", err)
		}
	}
}

func createTestDatabase(ctx context.Context, base, dbname string) error {
	db, err := sql.Open("postgres", base)
	if err != nil {
		return err
	}

	sql := fmt.Sprintf("CREATE DATABASE %s", dbname)
	_, err = db.ExecContext(ctx, sql)
	return err
}

func createSchema(ctx context.Context, db *sql.DB, path string) error {
	bs, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	sql := string(bs)
	_, err = db.ExecContext(ctx, sql)
	return err
}

func dropTestDatabase(ctx context.Context, base, dbname string) error {
	db, err := sql.Open("postgres", base)
	if err != nil {
		return err
	}

	sql := fmt.Sprintf("DROP DATABASE %s", dbname)
	_, err = db.ExecContext(ctx, sql)
	return err
}
