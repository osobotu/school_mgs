// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"database/sql"
	"time"
)

type Arm struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Class struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ClassHasArm struct {
	ClassID   int32     `json:"class_id"`
	ArmID     int32     `json:"arm_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Department struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type DepartmentHasSubject struct {
	SubjectID    int32     `json:"subject_id"`
	DepartmentID int32     `json:"department_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type FormMaster struct {
	ID        int32         `json:"id"`
	TeacherID sql.NullInt32 `json:"teacher_id"`
	ClassID   sql.NullInt32 `json:"class_id"`
	ArmID     sql.NullInt32 `json:"arm_id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

type Score struct {
	StudentID   int32     `json:"student_id"`
	TermScoreID int32     `json:"term_score_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Session struct {
	ID        int32        `json:"id"`
	Session   string       `json:"session"`
	StartDate sql.NullTime `json:"start_date"`
	EndDate   sql.NullTime `json:"end_date"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type Student struct {
	ID           int32          `json:"id"`
	FirstName    string         `json:"first_name"`
	LastName     string         `json:"last_name"`
	MiddleName   sql.NullString `json:"middle_name"`
	ClassID      sql.NullInt32  `json:"class_id"`
	DepartmentID sql.NullInt32  `json:"department_id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

type StudentOffersSubject struct {
	StudentID int32     `json:"student_id"`
	SubjectID int32     `json:"subject_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Subject struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Teacher struct {
	ID           int32          `json:"id"`
	FirstName    string         `json:"first_name"`
	LastName     string         `json:"last_name"`
	MiddleName   sql.NullString `json:"middle_name"`
	SubjectID    sql.NullInt32  `json:"subject_id"`
	DepartmentID sql.NullInt32  `json:"department_id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

type TeacherTeachesClass struct {
	TeacherID int32     `json:"teacher_id"`
	ClassID   int32     `json:"class_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Term struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Number    int32     `json:"number"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TermScore struct {
	ID         int32     `json:"id"`
	Assessment float64   `json:"assessment"`
	Exam       float64   `json:"exam"`
	SubjectID  int32     `json:"subject_id"`
	TermID     int32     `json:"term_id"`
	SessionID  int32     `json:"session_id"`
	ClassID    int32     `json:"class_id"`
	ArmID      int32     `json:"arm_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
