-- name: CreateUser :one 
INSERT INTO users (
    email,
    password_hash,
    role_id
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users 
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;