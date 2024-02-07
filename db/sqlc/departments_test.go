package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/osobotu/school_mgs/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateDepartment(t *testing.T) {
	createTestDepartment(t)
}

func TestDeleteDepartment(t *testing.T) {
	department := createTestDepartment(t)

	err := testQueries.DeleteDepartment(context.Background(), department.ID)
	require.NoError(t, err)

	department2, err := testQueries.GetDepartmentByID(context.Background(), department.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, department2)
}

func TestGetDepartmentByID(t *testing.T) {
	department := createTestDepartment(t)

	department2, err := testQueries.GetDepartmentByID(context.Background(), department.ID)
	require.NoError(t, err)
	require.NotEmpty(t, department2)

	compareDepartments(t, department, department2)

}

func TestListDepartments(t *testing.T) {
	departments, err := testQueries.ListAllDepartments(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, departments)
}

func compareDepartments(t *testing.T, dp1, dp2 Department) {
	require.Equal(t, dp1.ID, dp2.ID)
	require.Equal(t, dp1.Name, dp2.Name)
	require.Equal(t, dp1.Description, dp2.Description)
}

func createTestDepartment(t *testing.T) Department {
	arg := CreateDepartmentParams{
		Name:        utils.RandomString(5),
		Description: utils.RandomString(20),
	}
	department, err := testQueries.CreateDepartment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, department)

	require.Equal(t, arg.Name, department.Name)
	require.Equal(t, arg.Description, department.Description)

	require.NotZero(t, department.ID)
	require.NotZero(t, department.CreatedAt)
	require.NotZero(t, department.UpdatedAt)

	return department
}
