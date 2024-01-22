-- name: CreateStudentOffersSubject :one
INSERT INTO student_offers_subject (
    student_id,
    subject_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: ListSubjectsOfferedByStudentID :many
SELECT * FROM student_offers_subject
WHERE student_id = $1;

-- name: DeleteStudentOffersSubject :exec
DELETE FROM student_offers_subject
WHERE student_id = $1 AND subject_id = $2;