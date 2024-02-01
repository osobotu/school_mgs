// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: form_masters.sql

package db

import (
	"context"
	"database/sql"
)

const createFormMaster = `-- name: CreateFormMaster :one
INSERT INTO form_masters (
    teacher_id,
    class_id,
    arm_id
) VALUES (
    $1, $2, $3
) RETURNING id, teacher_id, class_id, arm_id, created_at, updated_at
`

type CreateFormMasterParams struct {
	TeacherID sql.NullInt32 `json:"teacher_id"`
	ClassID   sql.NullInt32 `json:"class_id"`
	ArmID     sql.NullInt32 `json:"arm_id"`
}

func (q *Queries) CreateFormMaster(ctx context.Context, arg CreateFormMasterParams) (FormMaster, error) {
	row := q.db.QueryRowContext(ctx, createFormMaster, arg.TeacherID, arg.ClassID, arg.ArmID)
	var i FormMaster
	err := row.Scan(
		&i.ID,
		&i.TeacherID,
		&i.ClassID,
		&i.ArmID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteFormMasterByClassID = `-- name: DeleteFormMasterByClassID :exec
DELETE FROM form_masters WHERE class_id = $1 AND arm_id = $2
`

type DeleteFormMasterByClassIDParams struct {
	ClassID sql.NullInt32 `json:"class_id"`
	ArmID   sql.NullInt32 `json:"arm_id"`
}

func (q *Queries) DeleteFormMasterByClassID(ctx context.Context, arg DeleteFormMasterByClassIDParams) error {
	_, err := q.db.ExecContext(ctx, deleteFormMasterByClassID, arg.ClassID, arg.ArmID)
	return err
}

const getFormMasterByID = `-- name: GetFormMasterByID :one
SELECT id, teacher_id, class_id, arm_id, created_at, updated_at FROM form_masters
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetFormMasterByID(ctx context.Context, id int32) (FormMaster, error) {
	row := q.db.QueryRowContext(ctx, getFormMasterByID, id)
	var i FormMaster
	err := row.Scan(
		&i.ID,
		&i.TeacherID,
		&i.ClassID,
		&i.ArmID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateFormMaster = `-- name: UpdateFormMaster :one
UPDATE form_masters
SET teacher_id = $3
WHERE class_id = $1 AND arm_id = $2 
RETURNING id, teacher_id, class_id, arm_id, created_at, updated_at
`

type UpdateFormMasterParams struct {
	ClassID   sql.NullInt32 `json:"class_id"`
	ArmID     sql.NullInt32 `json:"arm_id"`
	TeacherID sql.NullInt32 `json:"teacher_id"`
}

func (q *Queries) UpdateFormMaster(ctx context.Context, arg UpdateFormMasterParams) (FormMaster, error) {
	row := q.db.QueryRowContext(ctx, updateFormMaster, arg.ClassID, arg.ArmID, arg.TeacherID)
	var i FormMaster
	err := row.Scan(
		&i.ID,
		&i.TeacherID,
		&i.ClassID,
		&i.ArmID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}