package internal

import (
	"context"
	"database/sql"
	"stockin-api/domain"
	"stockin-api/queries"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type Item struct {
	ID        domain.UUID `fake:"{uuidv7}"`
	Title     string      `fake:"{sentence}"`
	URL       string      `fake:"{url}"`
	Thumbnail string      `fake:"{imageurl}"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewItem() *Item {
	var item Item
	gofakeit.Struct(&item)
	return &item
}

func CreateItem(ctx context.Context, db *sql.DB) (*Item, error) {
	item := NewItem()
	params := queries.InsertItemParams{
		ID:        item.ID,
		Title:     item.Title,
		Url:       item.URL,
		Thumbnail: item.Thumbnail,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
	err := queries.New(db).InsertItem(ctx, params)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func CreateItems(ctx context.Context, db *sql.DB, n int) ([]*Item, error) {
	items := []*Item{}
	for i := 0; i < n; i++ {
		item, err := CreateItem(ctx, db)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
