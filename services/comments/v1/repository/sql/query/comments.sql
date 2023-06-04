-- name: Create :exec
INSERT INTO comments(id, text, post_id, creator_id)
VALUES(@id, @text, @post_id, @creator_id);

-- name: ListPostComments :many
SELECT *
FROM comments
WHERE post_id = @post_id;
