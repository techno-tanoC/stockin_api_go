package domain

import (
	"context"
	"fmt"
	"stockin/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func ItemIndex(ctx context.Context, db DB, from int64, limit int) ([]*models.Item, error) {
	items, err := models.Items(
		models.ItemWhere.ID.LT(from),
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

func ItemCreate(ctx context.Context, db DB, title, url, thumbnail string) (*models.Item, error) {
	item := &models.Item{
		Title:     title,
		URL:       url,
		Thumbnail: thumbnail,
	}

	err := item.Insert(ctx, db, boil.Infer())
	if err != nil {
		return nil, fmt.Errorf("insert error: %w", err)
	}

	return item, nil
}

func ItemUpdate(ctx context.Context, db DB, id int64, title, url, thumbnail string) (*models.Item, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("begin error: %w", err)
	}
	defer tx.Commit()

	item := &models.Item{
		ID:        id,
		Title:     title,
		URL:       url,
		Thumbnail: thumbnail,
	}

	_, err = item.Update(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("update error: %w", err)
	}

	err = item.Reload(ctx, tx)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("reload error: %w", err)
	}

	return item, nil
}

func ItemDelete(ctx context.Context, db DB, id int64) error {
	item := &models.Item{
		ID: id,
	}
	_, err := item.Delete(ctx, db)
	if err != nil {
		return fmt.Errorf("delete error: %w", err)
	}

	return nil
}
