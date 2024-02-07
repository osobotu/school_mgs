// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: teachers.sql

package db

import (
	"context"
	"time"
)

const createTeacher = `-- name: CreateTeacher :one
INSERT INTO teachers (
    user_id,
    first_name,
    last_name,
    middle_name,
    subject_id,
    department_id
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING id, user_id, first_name, last_name, middle_name, subject_id, department_id, created_at, updated_at
`

type CreateTeacherParams struct {
	UserID       int32  `json:"user_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	MiddleName   string `json:"middle_name"`
	SubjectID    int32  `json:"subject_id"`
	DepartmentID int32  `json:"department_id"`
}

func (q *Queries) CreateTeacher(ctx context.Context, arg CreateTeacherParams) (Teacher, error) {
	row := q.db.QueryRowContext(ctx, createTeacher,
		arg.UserID,
		arg.FirstName,
		arg.LastName,
		arg.MiddleName,
		arg.SubjectID,
		arg.DepartmentID,
	)
	var i Teacher
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.FirstName,
		&i.LastName,
		&i.MiddleName,
		&i.SubjectID,
		&i.DepartmentID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteTeacher = `-- name: DeleteTeacher :exec
DELETE FROM teachers WHERE id = $1
`

func (q *Queries) DeleteTeacher(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteTeacher, id)
	return err
}

const getTeacherByID = `-- name: GetTeacherByID :one
SELECT id, user_id, first_name, last_name, middle_name, subject_id, department_id, created_at, updated_at FROM teachers
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTeacherByID(ctx context.Context, id int32) (Teacher, error) {
	row := q.db.QueryRowContext(ctx, getTeacherByID, id)
	var i Teacher
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.FirstName,
		&i.LastName,
		&i.MiddleName,
		&i.SubjectID,
		&i.DepartmentID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listTeachers = `-- name: ListTeachers :many
SELECT id, user_id, first_name, last_name, middle_name, subject_id, department_id, created_at, updated_at FROM teachers
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListTeachersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListTeachers(ctx context.Context, arg ListTeachersParams) ([]Teacher, error) {
	rows, err := q.db.QueryContext(ctx, listTeachers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Teacher{}
	for rows.Next() {
		var i Teacher
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.FirstName,
			&i.LastName,
			&i.MiddleName,
			&i.SubjectID,
			&i.DepartmentID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTeacher = `-- name: UpdateTeacher :one
UPDATE teachers
SET first_name = $2, last_name = $3, middle_name = $4, subject_id = $5, department_id = $6, updated_at = $7
WHERE id = $1
RETURNING id, user_id, first_name, last_name, middle_name, subject_id, department_id, created_at, updated_at
`

type UpdateTeacherParams struct {
	ID           int32     `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	MiddleName   string    `json:"middle_name"`
	SubjectID    int32     `json:"subject_id"`
	DepartmentID int32     `json:"department_id"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (q *Queries) UpdateTeacher(ctx context.Context, arg UpdateTeacherParams) (Teacher, error) {
	row := q.db.QueryRowContext(ctx, updateTeacher,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.MiddleName,
		arg.SubjectID,
		arg.DepartmentID,
		arg.UpdatedAt,
	)
	var i Teacher
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.FirstName,
		&i.LastName,
		&i.MiddleName,
		&i.SubjectID,
		&i.DepartmentID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
