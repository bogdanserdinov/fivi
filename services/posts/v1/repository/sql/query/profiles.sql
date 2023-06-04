-- name: CreatePosts :exec
INSERT INTO posts(id, payload, creator_id)
VALUES (@id, @payload, @creator_id);

-- name: Get :one
SELECT id, payload, creator_id, created_at
FROM posts
WHERE id = @id;

-- name: ListUserPosts :many
SELECT id, payload, creator_id, created_at
FROM posts
WHERE creator_id = @creator_id;

-- name: UpdatePost :exec
UPDATE posts
SET payload = @payload
WHERE id = @id;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = @id;

-- name: ListUsersPosts :many
SELECT id, payload, creator_id, created_at
FROM posts
WHERE creator_id = ANY(@creator_id::uuid[]);
