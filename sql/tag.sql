-- name: FindTagById :one
SELECT *
FROM tags
WHERE id = $1
;

-- name: IndexTags :many
SELECT *
FROM tags
ORDER BY id
;

-- name: InsertTag :exec
INSERT INTO tags(id, name, created_at, updated_at)
VALUES ($1, $2, $3, $4)
;

-- name: UpdateTag :exec
UPDATE tags
SET name = $2, updated_at = $3
WHERE id = $1
;

-- name: DeleteTag :exec
DELETE FROM tags
WHERE id = $1
;
