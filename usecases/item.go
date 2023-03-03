package usecases

import (
	"context"
	"database/sql"
	"fmt"
	"stockin-api/domain"
	"stockin-api/queries"
)

type ItemUsecaseImpl struct {
	db *sql.DB
}

func NewItemUsecase(db *sql.DB) *ItemUsecaseImpl {
	return &ItemUsecaseImpl{db}
}

func (u *ItemUsecaseImpl) Find(ctx context.Context, id domain.UUID) (*domain.Item, error) {
	q := queries.New(u.db)

	model, err := q.FindItem(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("find item error: %w", err)
	}
	item := domain.ItemFromModel(&model)

	return item, nil
}
