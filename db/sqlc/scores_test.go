package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateScore(t *testing.T) {
	score := createTestScore(t)
	testQueries.RunCleaners(t, &score)
}

func TestDeleteScore(t *testing.T) {
	score := createTestScore(t)
	err := testQueries.DeleteScore(context.Background(), score.StudentID)
	require.NoError(t, err)

	score, err = testQueries.GetScoreByStudentID(context.Background(), score.StudentID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, score)
}

func TestGetScoreByStudentId(t *testing.T) {
	score1 := createTestScore(t)
	score2, err := testQueries.GetScoreByStudentID(context.Background(), score1.StudentID)
	require.NoError(t, err)
	require.NotEmpty(t, score2)

	compareScores(t, score1, score2)

	testQueries.RunCleaners(t, &score1, &score2)

}

func createTestScore(t *testing.T) Score {
	student := createTestStudent(t)
	termScore := createTestTermScore(t)

	arg := CreateScoreParams{
		StudentID:   student.ID,
		TermScoreID: termScore.ID,
	}

	score, err := testQueries.CreateScore(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, score)

	require.Equal(t, arg.StudentID, score.StudentID)
	require.Equal(t, arg.TermScoreID, score.TermScoreID)

	testQueries.RunCleaners(t, &student, &termScore)
	return score
}

func compareScores(t *testing.T, score1, score2 Score) {
	require.Equal(t, score1.StudentID, score2.StudentID)
	require.Equal(t, score1.TermScoreID, score2.TermScoreID)
}

func (s *Score) Clean() {
	testQueries.DeleteScore(context.Background(), s.StudentID)
}
