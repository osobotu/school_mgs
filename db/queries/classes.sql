-- name: CreateClass :one
INSERT INTO classes (
    name
) VALUES (
    $1
) RETURNING *;

-- name: GetClassByID :one
SELECT * FROM classes
WHERE id = $1 LIMIT 1;

-- name: GetClassByName :one
SELECT * FROM classes 
WHERE name = $1 LIMIT 1;

-- name: ListClasses :many
SELECT * FROM classes
ORDER by name
LIMIT $1
OFFSET $2;

-- name: DeleteClass :exec
DELETE FROM classes WHERE id = $1;