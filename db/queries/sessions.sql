-- name: CreateSession :one
INSERT INTO sessions (
    session,
    start_date,
    end_date
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetSessionById :one
SELECT * FROM sessions
WHERE id = $1 LIMIT 1;

-- name: DeleteSession :exec
DELETE FROM sessions WHERE id = $1;

