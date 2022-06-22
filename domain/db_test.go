package domain_test

import (
	"context"
	"database/sql"
	"stockin/domain"

	_ "github.com/lib/pq"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	TestDatabase string `env:"TEST_DATABASE,required"`
}

func buildMockDB(ctx context.Context) (*domain.MockDB, func(), error) {
	domain.SetItemInsertHook()

	conf := new(Config)
	err := envconfig.Process(ctx, conf)
	if err != nil {
		return nil, nil, err
	}

	rawDB, err := sql.Open("postgres", conf.TestDatabase)
	if err != nil {
		return nil, nil, err
	}

	tx, err := rawDB.BeginTx(ctx, nil)
	if err != nil {
		return nil, func() {
			defer rawDB.Close()
		}, err
	}

	db := &domain.MockDB{&domain.MockTx{tx}}
	return db, func() {
		defer rawDB.Close()
		defer tx.Rollback()
	}, nil
}
