// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: students.sql

package db

import (
	"context"

	"github.com/lib/pq"
)

const createStudent = `-- name: CreateStudent :one
INSERT INTO students (
    first_name,
    last_name,
    class_id,
    subjects
) VALUES (
    $1, $2, $3, $4
) RETURNING id, first_name, last_name, middle_name, class_id, subjects, created_at, updated_at
`

type CreateStudentParams struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	ClassID   []int32 `json:"class_id"`
	Subjects  []int32 `json:"subjects"`
}

func (q *Queries) CreateStudent(ctx context.Context, arg CreateStudentParams) (Student, error) {
	row := q.db.QueryRowContext(ctx, createStudent,
		arg.FirstName,
		arg.LastName,
		pq.Array(arg.ClassID),
		pq.Array(arg.Subjects),
	)
	var i Student
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.MiddleName,
		pq.Array(&i.ClassID),
		pq.Array(&i.Subjects),
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteStudent = `-- name: DeleteStudent :exec
DELETE FROM students WHERE id = $1
`

func (q *Queries) DeleteStudent(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteStudent, id)
	return err
}

const getStudentById = `-- name: GetStudentById :one
SELECT id, first_name, last_name, middle_name, class_id, subjects, created_at, updated_at FROM students
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetStudentById(ctx context.Context, id int32) (Student, error) {
	row := q.db.QueryRowContext(ctx, getStudentById, id)
	var i Student
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.MiddleName,
		pq.Array(&i.ClassID),
		pq.Array(&i.Subjects),
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getStudentByName = `-- name: GetStudentByName :one
SELECT id, first_name, last_name, middle_name, class_id, subjects, created_at, updated_at FROM students
WHERE first_name = $1 LIMIT 1
`

func (q *Queries) GetStudentByName(ctx context.Context, firstName string) (Student, error) {
	row := q.db.QueryRowContext(ctx, getStudentByName, firstName)
	var i Student
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.MiddleName,
		pq.Array(&i.ClassID),
		pq.Array(&i.Subjects),
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listStudents = `-- name: ListStudents :many
SELECT id, first_name, last_name, middle_name, class_id, subjects, created_at, updated_at FROM students
ORDER by first_name
LIMIT $1
OFFSET $2
`

type ListStudentsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListStudents(ctx context.Context, arg ListStudentsParams) ([]Student, error) {
	rows, err := q.db.QueryContext(ctx, listStudents, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Student{}
	for rows.Next() {
		var i Student
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.MiddleName,
			pq.Array(&i.ClassID),
			pq.Array(&i.Subjects),
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

const updateClass = `-- name: UpdateClass :one
UPDATE students
SET class_id = $2
WHERE id = $1
RETURNING id, first_name, last_name, middle_name, class_id, subjects, created_at, updated_at
`

type UpdateClassParams struct {
	ID      int32   `json:"id"`
	ClassID []int32 `json:"class_id"`
}

func (q *Queries) UpdateClass(ctx context.Context, arg UpdateClassParams) (Student, error) {
	row := q.db.QueryRowContext(ctx, updateClass, arg.ID, pq.Array(arg.ClassID))
	var i Student
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.MiddleName,
		pq.Array(&i.ClassID),
		pq.Array(&i.Subjects),
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateSubjectsList = `-- name: UpdateSubjectsList :one
UPDATE students
SET subjects = $2
WHERE id = $1
RETURNING id, first_name, last_name, middle_name, class_id, subjects, created_at, updated_at
`

type UpdateSubjectsListParams struct {
	ID       int32   `json:"id"`
	Subjects []int32 `json:"subjects"`
}

func (q *Queries) UpdateSubjectsList(ctx context.Context, arg UpdateSubjectsListParams) (Student, error) {
	row := q.db.QueryRowContext(ctx, updateSubjectsList, arg.ID, pq.Array(arg.Subjects))
	var i Student
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.MiddleName,
		pq.Array(&i.ClassID),
		pq.Array(&i.Subjects),
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
