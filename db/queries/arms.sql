-- name: CreateArm :one
INSERT INTO arms (
    name
) VALUES (
    $1
) RETURNING *;

-- name: GetArmByID :one
SELECT * FROM arms
WHERE id = $1 LIMIT 1;

-- name: UpdateArm :one
UPDATE arms
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteArm :exec
DELETE FROM arms WHERE id = $1;