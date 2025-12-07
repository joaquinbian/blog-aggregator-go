-- name: CreatePost :one
INSERT INTO posts (
    id,
    created_at,
    updated_at,
    title,
    url,
    description,
    published_at,
    feed_id
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
ON CONFLICT (url) DO NOTHING
RETURNING *;


-- name: GetPostsForUser :many
SELECT P.*, F.name as feed_name FROM posts P 
INNER JOIN feeds F ON P.feed_id = F.id
WHERE F.user_id = $1
ORDER BY P.updated_at ASC
LIMIT $2;
