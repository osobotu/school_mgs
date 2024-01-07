package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/osobotu/school_mgs/db/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateTeacher(t *testing.T) {
	createTestTeacher(t)
}

func TestGetTeacherById(t *testing.T) {
	teacher1 := createTestTeacher(t)
	teacher2, err := testQueries.GetTeacherById(context.Background(), teacher1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, teacher2)

	compareTeachers(t, teacher1, teacher2)
}

func TestGetTeacherByName(t *testing.T) {
	teacher1 := createTestTeacher(t)
	teacher2, err := testQueries.GetTeacherByName(context.Background(), teacher1.LastName)

	require.NoError(t, err)
	require.NotEmpty(t, teacher2)

	compareTeachers(t, teacher1, teacher2)

}

func TestUpdateTeacher(t *testing.T) {
	teacher1 := createTestTeacher(t)
	arg := UpdateTeacherParams{
		ID:         teacher1.ID,
		FirstName:  teacher1.FirstName,
		LastName:   "Updated Last Name",
		MiddleName: teacher1.MiddleName,
		SubjectID:  teacher1.SubjectID,
		Classes:    teacher1.Classes,
	}
	teacher2, err := testQueries.UpdateTeacher(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, teacher2)

	require.Equal(t, arg.FirstName, teacher2.FirstName)
	require.Equal(t, arg.LastName, teacher2.LastName)
	require.Equal(t, arg.SubjectID, teacher2.SubjectID)
	require.Equal(t, arg.Classes, teacher2.Classes)

}

func TestDeleteTeacher(t *testing.T) {
	teacher1 := createTestTeacher(t)
	err := testQueries.DeleteTeacher(context.Background(), teacher1.ID)
	require.NoError(t, err)

	teacher2, err := testQueries.GetTeacherById(context.Background(), teacher1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, teacher2)
}

func TestListTeachers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createTestTeacher(t)
	}

	arg := ListTeachersParams{
		Limit:  5,
		Offset: 5,
	}

	teachers, err := testQueries.ListTeachers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, teachers, 5)

	for _, teacher := range teachers {
		require.NotEmpty(t, teacher)
	}
}

func createTestTeacher(t *testing.T) Teacher {
	classes := make([]int32, 3)

	arg := CreateTeacherParams{
		FirstName: utils.RandomString(5),
		LastName:  utils.RandomString(5),
		SubjectID: 1,
		Classes:   classes,
	}

	teacher, err := testQueries.CreateTeacher(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, teacher)

	require.Equal(t, arg.FirstName, teacher.FirstName)
	require.Equal(t, arg.LastName, teacher.LastName)
	require.Equal(t, arg.SubjectID, teacher.SubjectID)
	require.Equal(t, arg.Classes, teacher.Classes)

	require.NotZero(t, teacher.ID)
	require.NotZero(t, teacher.CreatedAt)

	return teacher

}

func compareTeachers(t *testing.T, teacher1, teacher2 Teacher) {
	require.Equal(t, teacher1.FirstName, teacher2.FirstName)
	require.Equal(t, teacher1.LastName, teacher2.LastName)
	require.Equal(t, teacher1.SubjectID, teacher2.SubjectID)
	require.Equal(t, teacher1.Classes, teacher2.Classes)
}
