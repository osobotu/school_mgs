package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateClassHasArms(t *testing.T) {
	createTestClassHasArms(t)
}

func TestDeleteClassHasArms(t *testing.T) {
	c := createTestClassHasArms(t)

	arg := DeleteClassHasArmsParams{
		ClassID: c.ClassID,
		ArmID:   c.ArmID,
	}

	err := testQueries.DeleteClassHasArms(context.Background(), arg)
	require.NoError(t, err)
}

func TestListArmsInClass(t *testing.T) {
	class := createTestClass(t)
	for i := 0; i < 5; i++ {
		createTestClassHasArmsWithClassID(t, class.ID)
	}

	classes, err := testQueries.ListArmsInClass(context.Background(), class.ID)
	require.NoError(t, err)
	require.Len(t, classes, 5)
}

func createTestClassHasArms(t *testing.T) ClassHasArm {
	class := createTestClass(t)
	arm := createTestArm(t)

	arg := CreateClassHasArmsParams{
		ClassID: class.ID,
		ArmID:   arm.ID,
	}

	classHasArms, err := testQueries.CreateClassHasArms(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, classHasArms)

	require.Equal(t, arg.ClassID, classHasArms.ClassID)
	require.Equal(t, arg.ArmID, classHasArms.ArmID)

	require.NotZero(t, classHasArms.CreatedAt)
	require.NotZero(t, classHasArms.UpdatedAt)
	return classHasArms
}
func createTestClassHasArmsWithClassID(t *testing.T, classID int32) ClassHasArm {
	arm := createTestArm(t)

	arg := CreateClassHasArmsParams{
		ClassID: classID,
		ArmID:   arm.ID,
	}

	classHasArms, err := testQueries.CreateClassHasArms(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, classHasArms)

	require.Equal(t, arg.ClassID, classHasArms.ClassID)
	require.Equal(t, arg.ArmID, classHasArms.ArmID)

	require.NotZero(t, classHasArms.CreatedAt)
	require.NotZero(t, classHasArms.UpdatedAt)
	return classHasArms
}
