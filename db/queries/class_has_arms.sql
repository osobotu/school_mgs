-- name: CreateClassHasArms :one
INSERT INTO class_has_arms (
    class_id,
    arm_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: ListArmsInClass :many
SELECT * FROM class_has_arms
WHERE class_id = $1;

-- name: DeleteClassHasArms :exec
DELETE FROM class_has_arms 
WHERE class_id = $1 AND arm_id = $2;