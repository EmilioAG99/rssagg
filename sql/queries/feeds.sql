-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES($1,$2,$3,$4,$5,$6)
RETURNING *;

-- name: GetFeeds :many

Select * from feeds;

-- name: GetNextFeedsToFetch :many

SELECT * FROM feeds order by last_fetched_at ASC NULLS FIRST 
limit $1;

-- name: MarkFeedAsFetched :one
UPDATE feeds 
set last_fetched_at = NOW(), 
updated_at = NOW()
where id = $1
RETURNING *;