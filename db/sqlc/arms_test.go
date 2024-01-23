package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/osobotu/school_mgs/db/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateArm(t *testing.T) {
	createTestArm(t)
}

func TestDeleteArm(t *testing.T) {
	a1 := createTestArm(t)
	err := testQueries.DeleteArm(context.Background(), a1.ID)
	require.NoError(t, err)

	a2, err := testQueries.GetArmByID(context.Background(), a1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, a2)
}

func TestUpdateArm(t *testing.T) {
	a1 := createTestArm(t)

	arg := UpdateArmParams{
		ID:   a1.ID,
		Name: "Updated Arm Name",
	}

	a2, err := testQueries.UpdateArm(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, a2)

	require.Equal(t, arg.Name, a2.Name)
	require.NotZero(t, a2.UpdatedAt)

}

func TestGetArmByID(t *testing.T) {
	a1 := createTestArm(t)

	a2, err := testQueries.GetArmByID(context.Background(), a1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, a2)

	compareArms(t, a1, a2)
}

func createTestArm(t *testing.T) Arm {

	name := utils.RandomString(2)

	arm, err := testQueries.CreateArm(context.Background(), name)

	require.NoError(t, err)
	require.NotEmpty(t, arm)

	require.Equal(t, name, arm.Name)

	require.NotZero(t, arm.ID)
	require.NotZero(t, arm.CreatedAt)
	require.NotZero(t, arm.UpdatedAt)
	return arm
}

func compareArms(t *testing.T, arm1, arm2 Arm) {
	require.Equal(t, arm1.ID, arm2.ID)
	require.Equal(t, arm1.Name, arm2.Name)
}
