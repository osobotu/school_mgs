-- name: CreateDepartmentHasSubject :one
INSERT INTO department_has_subjects (
    department_id,
    subject_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: ListSubjectsByDepartmentID :many
SELECT * FROM department_has_subjects
WHERE department_id = $1;

-- name: DeleteDepartmentHasSubjects :exec
DELETE FROM department_has_subjects
WHERE department_id = $1 AND subject_id = $2;