-- name: CreateFeedFollow :one 
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
    )
    RETURNING *
)

SELECT
     IFF.*,
    F.name AS feed_name,
    U.name AS user_name
FROM inserted_feed_follow IFF
INNER JOIN users U ON U.id = IFF.user_id
INNER JOIN feeds F ON F.id = IFF.feed_id;


-- name: GetFeedFollowsForUser :many
SELECT FF.*, U.name as user_name, F.name as feed_name FROM feed_follows FF 
INNER JOIN users U ON U.id = FF.user_id
INNER JOIN feeds F ON F.id = FF.feed_id
WHERE U.name = $1;