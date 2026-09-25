-- name: AddFeed :one
INSERT INTO feeds
  (id,created_at,updated_at,name,url,user_id)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
)
RETURNING *;

-- name: GetFeed :one
SELECT
  *
FROM
  feeds
WHERE
  name = $1;


-- name: GetFeedByUrl :one
SELECT
  *
FROM
  feeds
WHERE
  url = $1;


-- name: GetFeeds :many
SELECT
  *
FROM
  feeds;

-- name: GetFeedsWithUser :many
SELECT
  feeds.name,
  feeds.url,
  users.name AS user_name
FROM
  feeds
  INNER JOIN users ON feeds.user_id = users.id;

-- name: ResetFeeds :exec
DELETE FROM feeds;


-- name: MarkFeedFetched :exec
UPDATE
  feeds
SET
  last_fetched_at = $1,
  updated_at = $2
WHERE
  feeds.id = $3;

-- name: GetNextFeedToFetch :one
SELECT
  *
FROM
  feeds
ORDER BY
  last_fetched_at DESC
LIMIT 1;

