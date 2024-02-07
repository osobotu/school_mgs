-- name: CreateRole :one 
INSERT INTO roles (
    role
) VALUES (
    $1
) RETURNING *;

-- name: GetRoleByID :one
SELECT * FROM roles
WHERE id = $1 LIMIT 1;