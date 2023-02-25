-- name: InsertItemTag :exec
INSERT INTO item_tags(id, item_id, tag_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
;

-- name: DeleteItemTag :exec
DELETE FROM item_tags
WHERE id = $1
;
