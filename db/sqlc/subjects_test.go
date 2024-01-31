package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/osobotu/school_mgs/db/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateSubject(t *testing.T) {

	subject := createTestSubject(t)
	testQueries.RunCleaners(t, &subject)
}

func TestGetSubjectById(t *testing.T) {

	subject1 := createTestSubject(t)
	subject2, err := testQueries.GetSubjectByID(context.Background(), subject1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, subject2)

	compareSubjects(t, subject1, subject2)
	testQueries.RunCleaners(t, &subject1, &subject2)
}
func TestGetSubjectByName(t *testing.T) {
	subject1 := createTestSubject(t)
	subject2, err := testQueries.GetSubjectByName(context.Background(), subject1.Name)
	require.NoError(t, err)
	require.NotEmpty(t, subject2)

	compareSubjects(t, subject1, subject2)
	testQueries.RunCleaners(t, &subject1, &subject2)
}

func TestDeleteSubject(t *testing.T) {
	subject1 := createTestSubject(t)
	err := testQueries.DeleteSubject(context.Background(), subject1.ID)
	require.NoError(t, err)

	subject2, err := testQueries.GetSubjectByID(context.Background(), subject1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, subject2)
	testQueries.RunCleaners(t, &subject1, &subject2)

}

func TestListSubjects(t *testing.T) {
	subjects := make([]Subject, 5)
	for i := 0; i < 5; i++ {
		subjects = append(subjects, createTestSubject(t))
	}

	arg := ListSubjectsParams{
		Limit:  5,
		Offset: subjects[0].ID,
	}

	subjects, err := testQueries.ListSubjects(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, subjects, 5)

	for _, subject := range subjects {
		require.NotEmpty(t, subject)
		testQueries.RunCleaners(t, &subject)
	}
}

func createTestSubject(t *testing.T) Subject {
	name := utils.RandomString(5)
	subject, err := testQueries.CreateSubject(context.Background(), name)
	require.NoError(t, err)
	require.NotEmpty(t, subject)

	require.Equal(t, name, subject.Name)

	require.NotZero(t, subject.ID)
	require.NotZero(t, subject.CreatedAt)
	return subject
}

func compareSubjects(t *testing.T, subject1, subject2 Subject) {
	require.Equal(t, subject1.Name, subject2.Name)
	require.Equal(t, subject1.ID, subject2.ID)
}

func (s *Subject) Clean() {
	testQueries.DeleteSubject(context.Background(), s.ID)
}
