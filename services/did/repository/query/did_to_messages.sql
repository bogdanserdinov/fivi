-- name: UpsertDIDMessage :exec
INSERT INTO did_to_message (
    did,
    message
)
VALUES (
    @did,
    @message
) ON CONFLICT (did) DO
    UPDATE
    SET message = @message;

-- name: GetDIDMessage :one
SELECT *
FROM did_to_message
WHERE did = $1;
