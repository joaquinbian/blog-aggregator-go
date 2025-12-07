-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;


-- name: GetFeeds :many
SELECT * from feeds F INNER JOIN users U on F.user_id = U.id;

-- name: GetFeedBYUrl :one
SELECT * FROM feeds WHERE url = $1;
