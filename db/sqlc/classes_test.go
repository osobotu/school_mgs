package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/osobotu/school_mgs/db/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateClass(t *testing.T) {
	class := createTestClass(t)
	testQueries.RunCleaners(t, &class)
}

func TestGetClassById(t *testing.T) {
	class1 := createTestClass(t)
	class2, err := testQueries.GetClassById(context.Background(), class1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, class2)

	compareClass(t, class1, class2)
	testQueries.RunCleaners(t, &class1, &class2)

}

func TestGetClassByName(t *testing.T) {
	class1 := createTestClass(t)
	class2, err := testQueries.GetClassByName(context.Background(), class1.Name)
	require.NoError(t, err)
	require.NotEmpty(t, class2)

	compareClass(t, class1, class2)
	testQueries.RunCleaners(t, &class1, &class2)
}

func TestDeleteClass(t *testing.T) {
	class1 := createTestClass(t)
	err := testQueries.DeleteClass(context.Background(), class1.ID)
	require.NoError(t, err)

	class1, err = testQueries.GetClassById(context.Background(), class1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, class1)

}

func TestListClass(t *testing.T) {
	for i := 0; i < 10; i++ {
		createTestClass(t)
	}

	arg := ListClassesParams{
		Limit:  5,
		Offset: 5,
	}

	classes, err := testQueries.ListClasses(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, classes, 5)

	for _, class := range classes {
		require.NotEmpty(t, class)
		testQueries.RunCleaners(t, &class)
	}

}

func TestUpdateFormMaster(t *testing.T) {
	class1 := createTestClass(t)
	var formMasterId sql.NullInt32
	teacher := createTestTeacher(t)
	formMasterId.Scan(teacher.ID)
	arg := UpdateFormMasterParams{
		ID:           class1.ID,
		Name:         class1.Name,
		FormMasterID: formMasterId,
	}
	class2, err := testQueries.UpdateFormMaster(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, class2)

	require.Equal(t, arg.FormMasterID, class2.FormMasterID)
	require.Equal(t, arg.Name, class2.Name)

	testQueries.RunCleaners(t, &teacher, &class1, &class2)
}

func createTestClass(t *testing.T) Class {
	name := utils.RandomString(5)
	class, err := testQueries.CreateClass(context.Background(), name)
	require.NoError(t, err)
	require.NotEmpty(t, class)

	require.Equal(t, name, class.Name)
	require.NotZero(t, class.ID)
	return class
}

func compareClass(t *testing.T, class1, class2 Class) {
	require.Equal(t, class1.Name, class2.Name)
	require.Equal(t, class1.ID, class2.ID)
}

func (c *Class) Clean() {
	testQueries.DeleteClass(context.Background(), c.ID)
}
