-- name: ListLinks :many
SELECT * FROM links
WHERE url != ''
ORDER BY deleted_at, created_at desc;

-- name: CreateLink :one
INSERT INTO links (url) VALUES ($1) RETURNING *;

-- name: FindLinkById :one
SELECT * FROM links WHERE id = $1 LIMIT 1;

-- name: FindLinkByUrl :one
SELECT * FROM links WHERE url = $1 LIMIT 1;

-- name: DeleteLinkById :one
UPDATE links SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1 RETURNING *;

-- name: UnDeleteLinkById :one
UPDATE links SET deleted_at = NULL WHERE id = $1 RETURNING *;
