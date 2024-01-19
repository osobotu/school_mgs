-- name: CreateSubject :one
INSERT INTO subjects (
    name,
    classes
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetSubjectById :one
SELECT * FROM subjects
WHERE id = $1 LIMIT 1;

-- name: GetSubjectByName :one
SELECT * FROM subjects 
WHERE name = $1 LIMIT 1;

-- name: ListSubjects :many
SELECT * FROM subjects
ORDER by id
LIMIT $1
OFFSET $2;

-- name: DeleteSubject :exec
DELETE FROM subjects WHERE id = $1;