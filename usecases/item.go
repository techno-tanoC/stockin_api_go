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

func (u *ItemUsecaseImpl) FindByRange(ctx context.Context, params *domain.ItemRangeParams) ([]*domain.Item, error) {
	q := queries.New(u.db)

	ps := params.Build()
	models, err := q.FindItemsByRange(ctx, *ps)
	if err != nil {
		return nil, fmt.Errorf("find items by range error: %w", err)
	}

	items := []*domain.Item{}
	for _, model := range models {
		item := domain.ItemFromModel(&model)
		items = append(items, item)
	}

	return items, nil
}

func (u *ItemUsecaseImpl) Create(ctx context.Context, params *domain.ItemParams) (*domain.Item, error) {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("begin tx error: %w", err)
	}
	defer tx.Rollback()
	q := queries.New(tx)

	ps := params.BuildForInsert()

	err = q.InsertItem(ctx, *ps)
	if err != nil {
		return nil, fmt.Errorf("insert item error: %w", err)
	}

	model, err := q.FindItem(ctx, ps.ID)
	if err != nil {
		return nil, fmt.Errorf("find item error: %w", err)
	}
	item := domain.ItemFromModel(&model)

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("commit error: %w", err)
	}

	return item, nil
}

func (u *ItemUsecaseImpl) Update(ctx context.Context, id domain.UUID, params domain.ItemParams) (*domain.Item, error) {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("begin tx error: %w", err)
	}
	defer tx.Rollback()
	q := queries.New(tx)

	err = q.UpdateItem(ctx, *params.BuildForUpdate(id))
	if err != nil {
		return nil, fmt.Errorf("update item error: %w", err)
	}

	model, err := q.FindItem(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("find item error: %w", err)
	}
	item := domain.ItemFromModel(&model)

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("commit error: %w", err)
	}

	return item, nil
}

func (u *ItemUsecaseImpl) Delete(ctx context.Context, id domain.UUID) error {
	q := queries.New(u.db)

	err := q.DeleteItem(ctx, id)
	if err != nil {
		return fmt.Errorf("delete error: %w", err)
	}

	return nil
}
