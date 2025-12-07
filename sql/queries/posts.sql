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
INNER JOIN feed_follows FF ON P.feed_id = FF.feed_id
INNER JOIN feeds F ON f.id = P.feed_id 
WHERE FF.user_id = $1
ORDER BY P.updated_at ASC
LIMIT $2;