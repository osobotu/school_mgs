-- name: CreateScore :one
INSERT INTO scores (
    student_id,
    term_scores_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetScoreByStudentID :one
SELECT * FROM scores
WHERE student_id = $1 LIMIT 1;

-- name: DeleteScore :exec
DELETE FROM scores WHERE student_id = $1;