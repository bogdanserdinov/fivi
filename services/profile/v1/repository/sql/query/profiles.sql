-- name: CreateUser :exec
INSERT INTO profiles(id, name, username, mnemonic)
VALUES (@id, @name, @username, @mnemonic);

-- name: GetUser :one
SELECT id, name, username
FROM profiles
WHERE id = @id;

-- name: GetByUsername :one
SELECT id, name, username, mnemonic
FROM profiles
WHERE username = @username;

-- name: ListUserIDsWithName :many
SELECT id, name, username
FROM profiles
WHERE username ILIKE @username || '%';;

-- name: UpdateUser :exec
UPDATE profiles
SET name = @name, username = @username
WHERE id = @id;
