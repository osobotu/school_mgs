-- name: CreateSubject :one
INSERT INTO subjects (
    name
) VALUES (
    $1
) RETURNING *;

-- name: GetSubjectById :one
SELECT * FROM subjects
WHERE id = $1 LIMIT 1;

-- name: ListSubjects :many
SELECT * FROM subjects
ORDER by name ASC;

-- name: DeleteSubject :exec
DELETE FROM subjects WHERE id = $1;