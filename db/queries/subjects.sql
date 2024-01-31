-- name: CreateSubject :one
INSERT INTO subjects (
    name
) VALUES (
    $1
) RETURNING *;

-- name: GetSubjectByID :one
SELECT * FROM subjects
WHERE id = $1 LIMIT 1;

-- name: GetSubjectByName :one
SELECT * FROM subjects
WHERE name = $1 LIMIT 1;

-- name: ListSubjects :many
SELECT * FROM subjects
ORDER by name ASC
LIMIT $1
OFFSET $2;

-- name: DeleteSubject :exec
DELETE FROM subjects WHERE id = $1;