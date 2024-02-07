package db

import (
	"context"
	"testing"

	"github.com/osobotu/school_mgs/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateRole(t *testing.T) {
	createTestRole(t)
}

func TestGetRoleByID(t *testing.T) {
	role1 := createTestRole(t)
	role2, err := testQueries.GetRoleByID(context.Background(), role1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, role2)
	compareRoles(t, role1, role2)

}

func createTestRole(t *testing.T) Role {
	name := utils.RandomString(5)
	role, err := testQueries.CreateRole(context.Background(), name)
	require.NoError(t, err)
	require.NotEmpty(t, role)
	return role
}

func compareRoles(t *testing.T, role1, role2 Role) {
	require.Equal(t, role1.ID, role2.ID)
	require.Equal(t, role1.Role, role2.Role)
}
