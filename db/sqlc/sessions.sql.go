// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: sessions.sql

package db

import (
	"context"
	"database/sql"
)

const createSession = `-- name: CreateSession :one
INSERT INTO sessions (
    session,
    start_date,
    end_date
) VALUES (
    $1, $2, $3
) RETURNING id, session, start_date, end_date, created_at, updated_at
`

type CreateSessionParams struct {
	Session   string       `json:"session"`
	StartDate sql.NullTime `json:"start_date"`
	EndDate   sql.NullTime `json:"end_date"`
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, createSession, arg.Session, arg.StartDate, arg.EndDate)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Session,
		&i.StartDate,
		&i.EndDate,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteSession = `-- name: DeleteSession :exec
DELETE FROM sessions WHERE id = $1
`

func (q *Queries) DeleteSession(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteSession, id)
	return err
}

const getSessionById = `-- name: GetSessionById :one
SELECT id, session, start_date, end_date, created_at, updated_at FROM sessions
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetSessionById(ctx context.Context, id int32) (Session, error) {
	row := q.db.QueryRowContext(ctx, getSessionById, id)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Session,
		&i.StartDate,
		&i.EndDate,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
