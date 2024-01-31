-- name: CreateTeacher :one
INSERT INTO teachers (
    first_name,
    last_name,
    middle_name,
    subject_id,
    department_id
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetTeacherByID :one
SELECT * FROM teachers
WHERE id = $1 LIMIT 1;

-- name: ListTeachers :many
SELECT * FROM teachers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateTeacher :one
UPDATE teachers
SET first_name = $2, last_name = $3, middle_name = $4, subject_id = $5, department_id = $6, updated_at = $7
WHERE id = $1
RETURNING *;

-- name: DeleteTeacher :exec
DELETE FROM teachers WHERE id = $1;