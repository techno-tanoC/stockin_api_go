package domain

import (
	"context"
	"database/sql"
	"fmt"
	"stockin/models"

	"github.com/gofrs/uuid"
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

func BuildDB(database string) (*RealDB, func(), error) {
	SetItemInsertHook()

	rawDB, err := sql.Open("postgres", database)
	if err != nil {
		return nil, nil, fmt.Errorf("build db error: %w", err)
	}

	db := &RealDB{DB: rawDB}
	return db, func() {
		rawDB.Close()
	}, nil
}

func SetItemInsertHook() {
	models.AddItemHook(boil.BeforeInsertHook, func(_ context.Context, _ boil.ContextExecutor, item *models.Item) error {
		if item.ID == "" {
			id, err := uuid.NewV7(uuid.NanosecondPrecision)
			if err != nil {
				return err
			}

			item.ID = id.String()
		}

		return nil
	})
}
