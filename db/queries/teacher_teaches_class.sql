-- name: CreateTeacherTeachesClass :one
INSERT INTO teacher_teaches_class (
    teacher_id,
    class_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: ListTeacherTeachesClassByTeacherID :many
SELECT * FROM teacher_teaches_class
WHERE teacher_id = $1;

-- name: DeleteTeacherTeachesClass :exec
DELETE FROM teacher_teaches_class
WHERE teacher_id = $1 AND class_id = $2;
