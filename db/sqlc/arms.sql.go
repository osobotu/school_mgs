// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: arms.sql

package db

import (
	"context"
)

const createArm = `-- name: CreateArm :one
INSERT INTO arms (
    name
) VALUES (
    $1
) RETURNING id, name, created_at, updated_at
`

func (q *Queries) CreateArm(ctx context.Context, name string) (Arm, error) {
	row := q.db.QueryRowContext(ctx, createArm, name)
	var i Arm
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteArm = `-- name: DeleteArm :exec
DELETE FROM arms WHERE id = $1
`

func (q *Queries) DeleteArm(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteArm, id)
	return err
}

const getArmByID = `-- name: GetArmByID :one
SELECT id, name, created_at, updated_at FROM arms
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetArmByID(ctx context.Context, id int32) (Arm, error) {
	row := q.db.QueryRowContext(ctx, getArmByID, id)
	var i Arm
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateArm = `-- name: UpdateArm :one
UPDATE arms
SET name = $2
WHERE id = $1
RETURNING id, name, created_at, updated_at
`

type UpdateArmParams struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdateArm(ctx context.Context, arg UpdateArmParams) (Arm, error) {
	row := q.db.QueryRowContext(ctx, updateArm, arg.ID, arg.Name)
	var i Arm
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
