-- name: CreateScore :one
INSERT INTO scores (
    student_id,
    subject_id,
    first_term_assessment,
    first_term_exam,
    second_term_assessment,
    second_term_exam,
    third_term_assessment,
    third_term_exam,
    updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING *;

-- name: GetScoreByStudentId :one
SELECT * FROM scores
WHERE student_id = $1 LIMIT 1;

-- name: GetScoresBySubjectId :many
SELECT * FROM scores
WHERE subject_id = $1;

-- name: UpdateScoreByStudentId :one
UPDATE scores 
SET first_term_assessment = $2,
    first_term_exam = $3,
    second_term_assessment = $4,
    second_term_exam = $5,
    third_term_assessment = $6,
    third_term_exam = $7,
    updated_at = $8
WHERE student_id = $1
RETURNING *;

-- name: DeleteScore :exec
DELETE FROM scores WHERE student_id = $1;