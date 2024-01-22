-- name: CreateFormMaster :one
INSERT INTO form_masters (
    teacher_id,
    class_id,
    arm_id
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetFormMasterByID :one
SELECT * FROM form_masters
WHERE id = $1 LIMIT 1;

-- name: UpdateFormMaster :one
UPDATE form_masters
SET teacher_id = $3
WHERE class_id = $1 AND arm_id = $2 
RETURNING *;

-- name: DeleteFormMasterByClassID :exec
DELETE FROM form_masters WHERE class_id = $1 AND arm_id = $2;