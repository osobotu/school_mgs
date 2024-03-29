// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: teacher_teaches_class.sql

package db

import (
	"context"
)

const createTeacherTeachesClass = `-- name: CreateTeacherTeachesClass :one
INSERT INTO teacher_teaches_class (
    teacher_id,
    class_id
) VALUES (
    $1, $2
) RETURNING teacher_id, class_id, created_at, updated_at
`

type CreateTeacherTeachesClassParams struct {
	TeacherID int32 `json:"teacher_id"`
	ClassID   int32 `json:"class_id"`
}

func (q *Queries) CreateTeacherTeachesClass(ctx context.Context, arg CreateTeacherTeachesClassParams) (TeacherTeachesClass, error) {
	row := q.db.QueryRowContext(ctx, createTeacherTeachesClass, arg.TeacherID, arg.ClassID)
	var i TeacherTeachesClass
	err := row.Scan(
		&i.TeacherID,
		&i.ClassID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteTeacherTeachesClass = `-- name: DeleteTeacherTeachesClass :exec
DELETE FROM teacher_teaches_class
WHERE teacher_id = $1 AND class_id = $2
`

type DeleteTeacherTeachesClassParams struct {
	TeacherID int32 `json:"teacher_id"`
	ClassID   int32 `json:"class_id"`
}

func (q *Queries) DeleteTeacherTeachesClass(ctx context.Context, arg DeleteTeacherTeachesClassParams) error {
	_, err := q.db.ExecContext(ctx, deleteTeacherTeachesClass, arg.TeacherID, arg.ClassID)
	return err
}

const listTeacherTeachesClassByTeacherID = `-- name: ListTeacherTeachesClassByTeacherID :many
SELECT teacher_id, class_id, created_at, updated_at FROM teacher_teaches_class
WHERE teacher_id = $1
`

func (q *Queries) ListTeacherTeachesClassByTeacherID(ctx context.Context, teacherID int32) ([]TeacherTeachesClass, error) {
	rows, err := q.db.QueryContext(ctx, listTeacherTeachesClassByTeacherID, teacherID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TeacherTeachesClass{}
	for rows.Next() {
		var i TeacherTeachesClass
		if err := rows.Scan(
			&i.TeacherID,
			&i.ClassID,
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
