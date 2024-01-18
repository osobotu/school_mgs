package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/osobotu/school_mgs/db/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateSubject(t *testing.T) {

	classes := make([]int32, 1)

	arg := CreateSubjectParams{
		Name:    utils.RandomString(5),
		Classes: classes,
	}

	subject, err := testQueries.CreateSubject(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, subject)

	require.Equal(t, arg.Name, subject.Name)
	require.Equal(t, arg.Classes, subject.Classes)

	require.NotZero(t, subject.ID)
	require.NotZero(t, subject.CreatedAt)

}

func TestGetSubjectById(t *testing.T) {
	subject1 := createTestSubject(t)
	subject2, err := testQueries.GetSubjectById(context.Background(), subject1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, subject2)

	compareSubjects(t, subject1, subject2)
}
func TestGetSubjectByName(t *testing.T) {
	subject1 := createTestSubject(t)
	subject2, err := testQueries.GetSubjectByName(context.Background(), subject1.Name)
	require.NoError(t, err)
	require.NotEmpty(t, subject2)

	compareSubjects(t, subject1, subject2)
}

func TestDeleteSubject(t *testing.T) {
	subject1 := createTestSubject(t)
	err := testQueries.DeleteSubject(context.Background(), subject1.ID)
	require.NoError(t, err)

	subject1, err = testQueries.GetSubjectById(context.Background(), subject1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, subject1)

}

func TestListSubjects(t *testing.T) {
	for i := 0; i < 10; i++ {
		createTestSubject(t)
	}

	arg := ListSubjectsParams{
		Limit:  5,
		Offset: 5,
	}

	subjects, err := testQueries.ListSubjects(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, subjects, 5)

	for _, subject := range subjects {
		require.NotEmpty(t, subject)
	}
}

func createTestSubject(t *testing.T) Subject {
	arg := CreateSubjectParams{
		Name:    utils.RandomString(5),
		Classes: utils.RandomList(3),
	}
	subject, err := testQueries.CreateSubject(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, subject)

	require.Equal(t, arg.Name, subject.Name)
	require.Equal(t, arg.Classes, subject.Classes)

	require.NotZero(t, subject.ID)
	require.NotZero(t, subject.CreatedAt)
	return subject
}

func compareSubjects(t *testing.T, subject1, subject2 Subject) {
	require.Equal(t, subject1.Name, subject2.Name)
	require.Equal(t, subject1.Classes, subject2.Classes)
	require.Equal(t, subject1.ID, subject2.ID)
}
