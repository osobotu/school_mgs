-- name: CreateDepartment :one
INSERT INTO departments (
    name,
    description
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetDepartmentByID :one
SELECT * FROM departments 
WHERE id = $1;

-- name: ListAllDepartments :many
SELECT * FROM departments;

-- name: DeleteDepartment :exec
DELETE FROM departments WHERE id = $1;