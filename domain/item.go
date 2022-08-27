package domain

import (
	"context"
	"fmt"
	"stockin/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func ItemIndex(ctx context.Context, db DB, before string, limit int) ([]*models.Item, error) {
	items, err := models.Items(
		models.ItemWhere.ID.LT(before),
		qm.Limit(limit),
		qm.OrderBy(models.ItemColumns.ID+" DESC"),
	).All(ctx, db)
	if err != nil {
		return nil, fmt.Errorf("index error: %w", err)
	}

	// > Array and slice values encode as JSON arrays, except that []byte encodes as a base64-encoded string, and a nil slice encodes as the null JSON value.
	// > https://pkg.go.dev/encoding/json#Marshal
	if len(items) == 0 {
		items = []*models.Item{}
	}

	return items, nil
}

func ItemCreate(ctx context.Context, db DB, params *ItemParams) (*models.Item, error) {
	item := &models.Item{
		Title:     params.Title,
		URL:       params.URL,
		Thumbnail: params.Thumbnail,
	}

	err := item.Insert(ctx, db, boil.Infer())
	if err != nil {
		return nil, fmt.Errorf("insert error: %w", err)
	}

	return item, nil
}

func ItemUpdate(ctx context.Context, db DB, id string, params *ItemParams) (*models.Item, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("begin error: %w", err)
	}
	defer func() { _ = tx.Commit() }()

	item := &models.Item{
		ID:        id,
		Title:     params.Title,
		URL:       params.URL,
		Thumbnail: params.Thumbnail,
	}

	_, err = item.Update(ctx, tx, boil.Infer())
	if err != nil {
		_ = tx.Rollback()
		return nil, fmt.Errorf("update error: %w", err)
	}

	err = item.Reload(ctx, tx)
	if err != nil {
		_ = tx.Rollback()
		return nil, fmt.Errorf("reload error: %w", err)
	}

	return item, nil
}

func ItemDelete(ctx context.Context, db DB, id string) error {
	item := &models.Item{
		ID: id,
	}
	_, err := item.Delete(ctx, db)
	if err != nil {
		return fmt.Errorf("delete error: %w", err)
	}

	return nil
}

func ItemExport(ctx context.Context, db DB) ([]*models.Item, error) {
	items, err := models.Items().All(ctx, db)
	if err != nil {
		return nil, fmt.Errorf("item all error: %w", err)
	}

	// > Array and slice values encode as JSON arrays, except that []byte encodes as a base64-encoded string, and a nil slice encodes as the null JSON value.
	// > https://pkg.go.dev/encoding/json#Marshal
	if len(items) == 0 {
		items = []*models.Item{}
	}

	return items, nil
}

func ItemImport(ctx context.Context, db DB, items []*models.Item) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin error: %w", err)
	}
	defer func() { _ = tx.Commit() }()

	for _, item := range items {
		err = item.Insert(ctx, db, boil.Infer())
		if err != nil {
			_ = tx.Rollback()
			return fmt.Errorf("insert error: %w", err)
		}
	}

	return nil
}
