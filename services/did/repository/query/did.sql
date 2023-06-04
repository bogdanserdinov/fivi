-- name: CreateDIDMapping :one
INSERT INTO did_mappings (
    did,
    long_form_uri,
    document
)
VALUES (
    @did,
    @long_form_uri,
    @document
) RETURNING *;

-- name: GetDIDMapping :one
SELECT *
FROM did_mappings
WHERE did = $1;

-- name: DeleteDIDMapping :exec
DELETE FROM did_mappings
WHERE did = @did;
