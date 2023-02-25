-- name: FindItem :one
SELECT *
FROM items
WHERE id = $1
LIMIT 1
;

-- name: IndexItems :many
SELECT *
FROM items
ORDER BY id
;

-- name: InsertItem :exec
INSERT INTO items(id, title, url, thumbnail, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
;

-- name: UpdateItem :exec
UPDATE items
SET title = $2, url = $3, thumbnail = $4, updated_at = $5
WHERE id = $1
;

-- name: DeleteItem :exec
DELETE FROM items
WHERE id = $1
;
