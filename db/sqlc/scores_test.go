package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/osobotu/school_mgs/db/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateScore(t *testing.T) {
	createTestScore(t)
}

func TestDeleteScore(t *testing.T) {
	score := createTestScore(t)
	err := testQueries.DeleteScore(context.Background(), score.StudentID)
	require.NoError(t, err)

	score, err = testQueries.GetScoreByStudentId(context.Background(), score.StudentID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, score)
}

func TestGetScoreByStudentId(t *testing.T) {
	score1 := createTestScore(t)
	score2, err := testQueries.GetScoreByStudentId(context.Background(), score1.StudentID)
	require.NoError(t, err)
	require.NotEmpty(t, score2)

	compareScores(t, score1, score2)
}
func TestGetScoreBySubjectId(t *testing.T) {

	const subjectId int32 = 14
	scores, err := testQueries.GetScoresBySubjectId(context.Background(), subjectId)
	require.NoError(t, err)

	for _, score := range scores {
		require.NotEmpty(t, score)
		require.Equal(t, score.SubjectID, subjectId)
	}
}

func TestUpdateScoreByStudentId(t *testing.T) {
	score := createTestScore(t)
	var fa sql.NullInt32
	var fe sql.NullInt32

	fa.Scan(27)
	fe.Scan(45)
	arg := UpdateScoreByStudentIdParams{
		StudentID:           score.StudentID,
		FirstTermAssessment: fa,
		FirstTermExam:       fe,
	}

	score2, err := testQueries.UpdateScoreByStudentId(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, score2)

	require.Equal(t, arg.FirstTermAssessment, score2.FirstTermAssessment)
	require.Equal(t, arg.FirstTermExam, score2.FirstTermExam)
	require.Equal(t, arg.StudentID, score2.StudentID)
}

func createTestScore(t *testing.T) Score {
	arg := CreateScoreParams{
		StudentID: int32(utils.RandomInt(7, 11)),
		SubjectID: int32(utils.RandomInt(10, 15)),
		UpdatedAt: time.Now(),
	}
	score, err := testQueries.CreateScore(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, score)

	require.Equal(t, arg.StudentID, score.StudentID)
	require.Equal(t, arg.SubjectID, score.SubjectID)
	return score
}

func compareScores(t *testing.T, score1, score2 Score) {
	require.Equal(t, score1.StudentID, score2.StudentID)
	require.Equal(t, score1.SubjectID, score2.SubjectID)
}
