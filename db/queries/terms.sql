-- name: CreateTerm :one
INSERT INTO terms (
    name,
    number
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetTermByID :one
SELECT * FROM terms
WHERE id = $1 LIMIT 1;

-- name: DeleteTerm :exec
DELETE FROM terms WHERE id = $1;