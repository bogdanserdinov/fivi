-- name: CreateUser :exec
INSERT INTO profiles(id, name, email, username, mnemonic)
VALUES (@id, @name, @email, @username, @mnemonic);

-- name: GetUser :one
SELECT id, name, email, username
FROM profiles
WHERE id = @id;

-- name: GetByUsername :one
SELECT id, name, email, username, mnemonic
FROM profiles
WHERE username = @username;

-- name: ListUserIDsWithName :many
SELECT id, name, email, username
FROM profiles
WHERE username ILIKE @username || '%';;

-- name: UpdateUser :exec
UPDATE profiles
SET name = @name, username = @username, email = @email
WHERE id = @id;
