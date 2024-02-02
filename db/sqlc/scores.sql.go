// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: scores.sql

package db

import (
	"context"
)

const createScore = `-- name: CreateScore :one
INSERT INTO scores (
    student_id,
    term_score_id
) VALUES (
    $1, $2
) RETURNING student_id, term_score_id, created_at, updated_at
`

type CreateScoreParams struct {
	StudentID   int32 `json:"student_id"`
	TermScoreID int32 `json:"term_score_id"`
}

func (q *Queries) CreateScore(ctx context.Context, arg CreateScoreParams) (Score, error) {
	row := q.db.QueryRowContext(ctx, createScore, arg.StudentID, arg.TermScoreID)
	var i Score
	err := row.Scan(
		&i.StudentID,
		&i.TermScoreID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteScore = `-- name: DeleteScore :exec
DELETE FROM scores WHERE student_id = $1
`

func (q *Queries) DeleteScore(ctx context.Context, studentID int32) error {
	_, err := q.db.ExecContext(ctx, deleteScore, studentID)
	return err
}

const getScoreByStudentID = `-- name: GetScoreByStudentID :one
SELECT student_id, term_score_id, created_at, updated_at FROM scores
WHERE student_id = $1 LIMIT 1
`

func (q *Queries) GetScoreByStudentID(ctx context.Context, studentID int32) (Score, error) {
	row := q.db.QueryRowContext(ctx, getScoreByStudentID, studentID)
	var i Score
	err := row.Scan(
		&i.StudentID,
		&i.TermScoreID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
