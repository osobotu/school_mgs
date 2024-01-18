package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/osobotu/school_mgs/db/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateScore(t *testing.T) {
	score := createTestScore(t)
	cleanUpScore(t, score)
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
	cleanUpScore(t, score1)
	cleanUpScore(t, score2)

}

func createTestScore(t *testing.T) Score {
	arg := CreateScoreParams{
		StudentID:    int32(utils.RandomInt(1, 2)),
		TermScoresID: int32(utils.RandomInt(1, 2)),
		UpdatedAt:    time.Now(),
	}
	score, err := testQueries.CreateScore(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, score)

	require.Equal(t, arg.StudentID, score.StudentID)
	require.Equal(t, arg.TermScoresID, score.TermScoresID)
	return score
}

func compareScores(t *testing.T, score1, score2 Score) {
	require.Equal(t, score1.StudentID, score2.StudentID)
	require.Equal(t, score1.TermScoresID, score2.TermScoresID)
}

func cleanUpScore(t *testing.T, score Score) {
	t.Cleanup(func() {
		fmt.Println("CleanUp called")
		testQueries.DeleteScore(context.Background(), score.StudentID)
	})
}
