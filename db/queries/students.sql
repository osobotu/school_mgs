-- name: CreateStudent :one
INSERT INTO students (
    first_name,
    last_name,
    middle_name,
    class_id,
    subjects
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetStudentById :one
SELECT * FROM students
WHERE id = $1 LIMIT 1;

-- name: ListStudents :many
SELECT * FROM students
ORDER by first_name
LIMIT $1
OFFSET $2;

-- name: UpdateStudent :one
UPDATE students
SET first_name = $2, last_name = $3, middle_name = $4, class_id = $5, subjects = $6
WHERE id = $1
RETURNING *;

-- name: DeleteStudent :exec
DELETE FROM students WHERE id = $1;