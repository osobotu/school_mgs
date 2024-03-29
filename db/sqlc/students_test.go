package db

import (
	"context"
	"testing"

	"github.com/osobotu/school_mgs/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateStudent(t *testing.T) {
	student := createTestStudent(t)
	testQueries.RunCleaners(t, &student)
}

func TestGetStudentById(t *testing.T) {
	student1 := createTestStudent(t)
	student2, err := testQueries.GetStudentByID(context.Background(), student1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, student2)

	compareStudents(t, student1, student2)

	testQueries.RunCleaners(t, &student1, &student2)

}

func TestListStudents(t *testing.T) {
	for i := 0; i < 10; i++ {
		createTestStudent(t)
	}

	arg := ListStudentsParams{
		Limit:  5,
		Offset: 5,
	}

	students, err := testQueries.ListStudents(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, students, 5)

	for _, student := range students {
		require.NotEmpty(t, student)
		testQueries.RunCleaners(t, &student)
	}
}

func TestUpdateStudent(t *testing.T) {

	newFirstName := utils.RandomString(5)
	newLastName := utils.RandomString(5)
	newClass := createTestClass(t)
	newDepartment := createTestDepartment(t)

	student1 := createTestStudent(t)

	arg := UpdateStudentParams{
		ID:           student1.ID,
		FirstName:    newFirstName,
		LastName:     newLastName,
		MiddleName:   student1.MiddleName,
		ClassID:      newClass.ID,
		DepartmentID: newDepartment.ID,
	}

	student2, err := testQueries.UpdateStudent(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, student2)

	require.Equal(t, arg.ID, student2.ID)
	require.Equal(t, arg.FirstName, student2.FirstName)
	require.Equal(t, arg.LastName, student2.LastName)
	require.Equal(t, arg.MiddleName, student2.MiddleName)
	require.Equal(t, arg.ClassID, student2.ClassID)
	require.Equal(t, arg.DepartmentID, student2.DepartmentID)
}

func createTestStudent(t *testing.T) Student {
	class := createTestClass(t)
	department := createTestDepartment(t)
	user := createTestUser(t)

	arg := CreateStudentParams{
		UserID:       user.ID,
		FirstName:    utils.RandomString(7),
		LastName:     utils.RandomString(7),
		ClassID:      class.ID,
		DepartmentID: department.ID,
	}

	student, err := testQueries.CreateStudent(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, student)

	require.Equal(t, arg.FirstName, student.FirstName)
	require.Equal(t, arg.ClassID, student.ClassID)
	require.Equal(t, arg.LastName, student.LastName)
	require.Equal(t, arg.ClassID, student.ClassID)
	require.Equal(t, arg.DepartmentID, student.DepartmentID)
	return student
}

func compareStudents(t *testing.T, student1, student2 Student) {
	require.Equal(t, student1.FirstName, student2.FirstName)
	require.Equal(t, student1.ClassID, student2.ClassID)
	require.Equal(t, student1.LastName, student2.LastName)
	require.Equal(t, student1.ClassID, student2.ClassID)
	require.Equal(t, student1.DepartmentID, student2.DepartmentID)
}

func (s *Student) Clean() {
	testQueries.DeleteStudent(context.Background(), s.ID)
}
