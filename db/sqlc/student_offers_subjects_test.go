package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateStudentOffersSubject(t *testing.T) {
	createTestStudentOffersSubject(t)
}

func TestDeleteStudentOffersSubject(t *testing.T) {
	sos := createTestStudentOffersSubject(t)
	arg := DeleteStudentOffersSubjectParams{
		StudentID: sos.StudentID,
		SubjectID: sos.SubjectID,
	}
	err := testQueries.DeleteStudentOffersSubject(context.Background(), arg)
	require.NoError(t, err)
}

func TestListSubjectsOfferedByStudent(t *testing.T) {
	student := createTestStudent(t)
	for i := 0; i < 5; i++ {
		createTestStudentOffersSubjectWithStudentID(t, student.ID)
	}

	subjects, err := testQueries.ListSubjectsOfferedByStudentID(context.Background(), student.ID)
	require.NoError(t, err)
	require.Len(t, subjects, 5)

	for _, subject := range subjects {
		require.NotEmpty(t, subject)
	}

}

func createTestStudentOffersSubject(t *testing.T) StudentOffersSubject {
	student := createTestStudent(t)
	subject := createTestSubject(t)

	arg := CreateStudentOffersSubjectParams{
		StudentID: student.ID,
		SubjectID: subject.ID,
	}
	studentOffersSubject, err := testQueries.CreateStudentOffersSubject(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, studentOffersSubject)

	require.NotZero(t, studentOffersSubject.CreatedAt)
	require.NotZero(t, studentOffersSubject.UpdatedAt)
	return studentOffersSubject
}

func createTestStudentOffersSubjectWithStudentID(t *testing.T, studentID int32) StudentOffersSubject {

	subject := createTestSubject(t)

	arg := CreateStudentOffersSubjectParams{
		StudentID: studentID,
		SubjectID: subject.ID,
	}
	studentOffersSubject, err := testQueries.CreateStudentOffersSubject(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, studentOffersSubject)

	require.NotZero(t, studentOffersSubject.CreatedAt)
	require.NotZero(t, studentOffersSubject.UpdatedAt)
	return studentOffersSubject
}
