package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateDepartmentHasSubject(t *testing.T) {
	createTestDepartmentHasSubject(t)
}

func TestDeleteDepartmentHasSubjects(t *testing.T) {
	dhs := createTestDepartmentHasSubject(t)
	arg := DeleteDepartmentHasSubjectsParams{
		DepartmentID: dhs.DepartmentID,
		SubjectID:    dhs.SubjectID,
	}

	err := testQueries.DeleteDepartmentHasSubjects(context.Background(), arg)
	require.NoError(t, err)
}

func TestListSubjectsByDepartmentID(t *testing.T) {
	department := createTestDepartment(t)
	for i := 0; i < 5; i++ {
		createTestDepartmentHasSubjectWithDepartmentID(t, department.ID)
	}

	subjects, err := testQueries.ListSubjectsByDepartmentID(context.Background(), department.ID)
	require.NoError(t, err)
	require.Len(t, subjects, 5)

	for _, subject := range subjects {
		require.NotEmpty(t, subject)
	}

}

func createTestDepartmentHasSubject(t *testing.T) DepartmentHasSubject {
	department := createTestDepartment(t)
	subject := createTestSubject(t)
	arg := CreateDepartmentHasSubjectParams{
		DepartmentID: department.ID,
		SubjectID:    subject.ID,
	}
	departmentHasSubject, err := testQueries.CreateDepartmentHasSubject(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, departmentHasSubject)

	require.Equal(t, arg.DepartmentID, departmentHasSubject.DepartmentID)
	require.Equal(t, arg.SubjectID, departmentHasSubject.SubjectID)

	require.NotZero(t, departmentHasSubject.CreatedAt)
	require.NotZero(t, departmentHasSubject.UpdatedAt)
	return departmentHasSubject
}

func createTestDepartmentHasSubjectWithDepartmentID(t *testing.T, departmentID int32) DepartmentHasSubject {
	subject := createTestSubject(t)

	arg := CreateDepartmentHasSubjectParams{
		DepartmentID: departmentID,
		SubjectID:    subject.ID,
	}
	departmentHasSubject, err := testQueries.CreateDepartmentHasSubject(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, departmentHasSubject)

	require.Equal(t, arg.DepartmentID, departmentHasSubject.DepartmentID)
	require.Equal(t, arg.SubjectID, departmentHasSubject.SubjectID)

	require.NotZero(t, departmentHasSubject.CreatedAt)
	require.NotZero(t, departmentHasSubject.UpdatedAt)
	return departmentHasSubject
}
