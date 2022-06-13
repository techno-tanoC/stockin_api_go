package domain

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type TxBeginner interface {
	BeginTx(context.Context, *sql.TxOptions) (boil.ContextTransactor, error)
}

type DB interface {
	boil.ContextExecutor
	TxBeginner
}

type RealDB struct {
	*sql.DB
}

func (db *RealDB) BeginTx(ctx context.Context, opts *sql.TxOptions) (boil.ContextTransactor, error) {
	return db.DB.BeginTx(ctx, opts)
}

type MockTx struct {
	*sql.Tx
}

func (tx *MockTx) Commit() error {
	return nil
}

func (tx *MockTx) Rollback() error {
	return nil
}

type MockDB struct {
	*MockTx
}

func (db *MockDB) BeginTx(ctx context.Context, opts *sql.TxOptions) (boil.ContextTransactor, error) {
	return db.MockTx, nil
}
