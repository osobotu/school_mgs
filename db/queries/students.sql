-- name: CreateStudent :one
INSERT INTO students (
    first_name,
    last_name,
    class_id,
    subjects
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetStudentById :one
SELECT * FROM students
WHERE id = $1 LIMIT 1;

-- name: GetStudentByName :one
SELECT * FROM students
WHERE first_name = $1 LIMIT 1;

-- name: ListStudents :many
SELECT * FROM students
ORDER by first_name
LIMIT $1
OFFSET $2;

-- name: UpdateClass :one
UPDATE students
SET class_id = $2
WHERE id = $1
RETURNING *;

-- name: UpdateSubjectsList :one
UPDATE students
SET subjects = $2
WHERE id = $1
RETURNING *;

-- name: DeleteStudent :exec
DELETE FROM students WHERE id = $1;