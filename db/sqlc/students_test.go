package db

import (
	"context"
	"testing"

	"github.com/osobotu/school_mgs/db/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateStudent(t *testing.T) {
	student := createTestStudent(t)
	testQueries.RunCleaners(t, &student)
}

func TestGetStudentById(t *testing.T) {
	student1 := createTestStudent(t)
	student2, err := testQueries.GetStudentById(context.Background(), student1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, student2)

	compareStudents(t, student1, student2)

	testQueries.RunCleaners(t, &student1, &student2)

}
func TestGetStudentByName(t *testing.T) {
	student1 := createTestStudent(t)
	student2, err := testQueries.GetStudentByName(context.Background(), student1.FirstName)
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

func TestUpdateClass(t *testing.T) {
	classID := make([]int32, 1)
	student1 := createTestStudent(t)
	arg := UpdateClassParams{
		ID:      student1.ID,
		ClassID: classID,
	}
	student2, err := testQueries.UpdateClass(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, student2)

	require.Equal(t, arg.ClassID, student2.ClassID)
	testQueries.RunCleaners(t, &student1, &student2)
}

func TestUpdateSubjectsList(t *testing.T) {
	student1 := createTestStudent(t)
	newSubjects := utils.RandomList(7)
	arg := UpdateSubjectsListParams{
		ID:       student1.ID,
		Subjects: newSubjects,
	}
	student2, err := testQueries.UpdateSubjectsList(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, student2)

	require.Equal(t, arg.Subjects, student2.Subjects)
	testQueries.RunCleaners(t, &student1, &student2)
}

func createTestStudent(t *testing.T) Student {
	classID := make([]int32, 1)

	arg := CreateStudentParams{
		FirstName: utils.RandomString(7),
		LastName:  utils.RandomString(7),
		ClassID:   classID,
		Subjects:  utils.RandomList(5),
	}

	student, err := testQueries.CreateStudent(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, student)

	require.Equal(t, arg.FirstName, student.FirstName)
	require.Equal(t, arg.ClassID, student.ClassID)
	require.Equal(t, arg.LastName, student.LastName)
	require.Equal(t, arg.Subjects, student.Subjects)
	return student
}

func compareStudents(t *testing.T, student1, student2 Student) {
	require.Equal(t, student1.FirstName, student2.FirstName)
	require.Equal(t, student1.ClassID, student2.ClassID)
	require.Equal(t, student1.LastName, student2.LastName)
	require.Equal(t, student1.Subjects, student2.Subjects)
}

func (s *Student) Clean() {
	testQueries.DeleteStudent(context.Background(), s.ID)
}
