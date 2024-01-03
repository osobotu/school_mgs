// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"database/sql"
	"time"
)

type Class struct {
	ID           int32         `json:"id"`
	Name         string        `json:"name"`
	FormMasterID sql.NullInt32 `json:"form_master_id"`
	CreatedAt    time.Time     `json:"created_at"`
}

type Score struct {
	StudentID            int32         `json:"student_id"`
	SubjectID            int32         `json:"subject_id"`
	FirstTermAssessment  sql.NullInt32 `json:"first_term_assessment"`
	FirstTermExam        sql.NullInt32 `json:"first_term_exam"`
	SecondTermAssessment sql.NullInt32 `json:"second_term_assessment"`
	SecondTermExam       sql.NullInt32 `json:"second_term_exam"`
	ThirdTermAssessment  sql.NullInt32 `json:"third_term_assessment"`
	ThirdTermExam        sql.NullInt32 `json:"third_term_exam"`
	CreatedAt            time.Time     `json:"created_at"`
	UpdatedAt            time.Time     `json:"updated_at"`
}

type Student struct {
	ID         int32          `json:"id"`
	FirstName  string         `json:"first_name"`
	LastName   string         `json:"last_name"`
	MiddleName sql.NullString `json:"middle_name"`
	ClassID    int32          `json:"class_id"`
	Subjects   []int32        `json:"subjects"`
	CreatedAt  time.Time      `json:"created_at"`
}

type Subject struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	// These are the classes that can take this subject
	Classes   []int32   `json:"classes"`
	CreatedAt time.Time `json:"created_at"`
}

type Teacher struct {
	ID         int32          `json:"id"`
	FirstName  string         `json:"first_name"`
	LastName   string         `json:"last_name"`
	MiddleName sql.NullString `json:"middle_name"`
	SubjectID  int32          `json:"subject_id"`
	Classes    []int32        `json:"classes"`
	CreatedAt  time.Time      `json:"created_at"`
}
