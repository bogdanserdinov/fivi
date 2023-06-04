-- name: Like :exec
INSERT INTO likes(post_id, user_id)
VALUES(@post_id, @user_id);

-- name: DeleteLike :exec
DELETE FROM likes
WHERE post_id = @post_id AND user_id = @user_id;

-- name: ListLikes :many
SELECT *
FROM likes
WHERE post_id = @post_id;

-- name: CountPostLikes :one
SELECT COUNT(*)
FROM likes
WHERE post_id = @post_id;
