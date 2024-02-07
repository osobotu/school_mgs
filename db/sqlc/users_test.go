package db

import (
	"context"
	"testing"

	"github.com/osobotu/school_mgs/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	createTestUser(t)
}

func TestGetUserByID(t *testing.T) {
	user1 := createTestUser(t)
	user2, err := testQueries.GetUserByID(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	compareUsers(t, user1, user2)

}

func createTestUser(t *testing.T) User {
	role := createTestRole(t)

	hashedPassword, err := utils.HashPassword(utils.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Email:        utils.RandomEmail(),
		PasswordHash: hashedPassword,
		RoleID:       role.ID,
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	return user
}

func compareUsers(t *testing.T, user1, user2 User) {
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.PasswordHash, user2.PasswordHash)
}
