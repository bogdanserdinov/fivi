-- name: CreateFollow :exec
INSERT INTO "followers" (
    "id",
    "follower_id",
    "followee_id"
)
VALUES (@id, @follower_id, @followee_id);

-- name: ListFollowers :many
SELECT id, follower_id FROM "followers" WHERE "followee_id" = @followee_id;

-- name: ListFollowings :many
SELECT id, followee_id FROM "followers" WHERE "follower_id" = @follower_id;

-- name: GetFollower :one
SELECT id FROM "followers" WHERE "follower_id" = @follower_id AND "followee_id" = @followee_id;

-- name: CountFollowers :one
SELECT COUNT(*)::BIGINT FROM "followers" WHERE "followee_id" = @followee_id;

-- name: CountFollowings :one
SELECT COUNT(*)::BIGINT FROM "followers" WHERE "follower_id" = @follower_id;

-- name: DeleteFollow :exec
DELETE FROM "followers" WHERE "id" = @id;

-- name: IsFollowUser :one
SELECT EXISTS(
               SELECT * FROM "followers"
               WHERE "follower_id" = @follower_id
                 AND "followee_id" = @followee_id
           );

