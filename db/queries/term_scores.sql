-- name: CreateTermScore :one
INSERT INTO term_scores (
    assessment,
    exam,
    subject_id,
    term_id,
    session_id,
    class_id,
    arm_id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetTermScoreByID :one
SELECT * FROM term_scores 
WHERE id = $1 LIMIT 1;

-- name: UpdateTermScoreByID :one
UPDATE term_scores
SET assessment = $2, exam = $3, updated_at = $4
WHERE id = $1
RETURNING *;

-- name: DeleteTermScore :exec
DELETE FROM term_scores WHERE id = $1;

-- name: ListTermScoresForSubjectAndClass :many
SELECT * FROM term_scores
WHERE subject_id = $3 AND class_id = $4
ORDER BY id
LIMIT $1
OFFSET $2;


