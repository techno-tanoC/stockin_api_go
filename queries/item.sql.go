// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: item.sql

package queries

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
)

const deleteItem = `-- name: DeleteItem :exec
DELETE FROM items
WHERE id = $1
`

func (q *Queries) DeleteItem(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteItem, id)
	return err
}

const findItem = `-- name: FindItem :one
SELECT id, title, url, thumbnail, created_at, updated_at
FROM items
WHERE id = $1
LIMIT 1
`

func (q *Queries) FindItem(ctx context.Context, id uuid.UUID) (Item, error) {
	row := q.db.QueryRowContext(ctx, findItem, id)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Url,
		&i.Thumbnail,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findItemsByRange = `-- name: FindItemsByRange :many
SELECT id, title, url, thumbnail, created_at, updated_at
FROM items
WHERE id < $1
ORDER BY id DESC
LIMIT $2
`

type FindItemsByRangeParams struct {
	ID    uuid.UUID
	Limit int32
}

func (q *Queries) FindItemsByRange(ctx context.Context, arg FindItemsByRangeParams) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, findItemsByRange, arg.ID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Url,
			&i.Thumbnail,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const indexItems = `-- name: IndexItems :many
SELECT id, title, url, thumbnail, created_at, updated_at
FROM items
ORDER BY id
`

func (q *Queries) IndexItems(ctx context.Context) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, indexItems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Url,
			&i.Thumbnail,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertItem = `-- name: InsertItem :exec
INSERT INTO items(id, title, url, thumbnail, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
`

type InsertItemParams struct {
	ID        uuid.UUID
	Title     string
	Url       string
	Thumbnail string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) InsertItem(ctx context.Context, arg InsertItemParams) error {
	_, err := q.db.ExecContext(ctx, insertItem,
		arg.ID,
		arg.Title,
		arg.Url,
		arg.Thumbnail,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const updateItem = `-- name: UpdateItem :exec
UPDATE items
SET title = $2, url = $3, thumbnail = $4, updated_at = $5
WHERE id = $1
`

type UpdateItemParams struct {
	ID        uuid.UUID
	Title     string
	Url       string
	Thumbnail string
	UpdatedAt time.Time
}

func (q *Queries) UpdateItem(ctx context.Context, arg UpdateItemParams) error {
	_, err := q.db.ExecContext(ctx, updateItem,
		arg.ID,
		arg.Title,
		arg.Url,
		arg.Thumbnail,
		arg.UpdatedAt,
	)
	return err
}
