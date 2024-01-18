-- name: CreateScore :one
INSERT INTO scores (
    student_id,
    term_scores_id,
    updated_at
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetScoreByStudentId :one
SELECT * FROM scores
WHERE student_id = $1 LIMIT 1;

-- name: DeleteScore :exec
DELETE FROM scores WHERE student_id = $1;