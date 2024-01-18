package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateTermScore(t *testing.T) {
	createTestTermScore(t)
}

func TestDeleteTermScore(t *testing.T) {
	termScore1 := createTestTermScore(t)
	err := testQueries.DeleteTermScore(context.Background(), termScore1.ID)
	require.NoError(t, err)

	termScore2, err := testQueries.GetTermScoreById(context.Background(), termScore1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, termScore2)
}

func TestGetTermScoreByID(t *testing.T) {
	termScore1 := createTestTermScore(t)
	termScore2, err := testQueries.GetTermScoreById(context.Background(), termScore1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, termScore2)

	compareTermScores(t, termScore1, termScore2)
}

func TestListTermScoresForSubjectAndClass(t *testing.T) {

	for i := 0; i < 5; i++ {
		createTestTermScore(t)
	}

	arg := ListTermScoresForSubjectAndClassParams{
		Limit:     5,
		Offset:    5,
		SubjectID: 1,
		ClassID:   1,
	}

	termScores, err := testQueries.ListTermScoresForSubjectAndClass(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, termScores, 5)

	for _, termScore := range termScores {
		require.NotEmpty(t, termScore)
	}

}

func TestUpdateScoreById(t *testing.T) {
	termScore1 := createTestTermScore(t)

	var ass sql.NullFloat64
	ass.Scan(termScore1.Assessment.Float64 + 5)
	arg := UpdateTermScoreByIdParams{
		ID:         termScore1.ID,
		Assessment: ass,
		Exam:       termScore1.Exam,
		UpdatedAt:  time.Now().UTC(),
	}

	termScore2, err := testQueries.UpdateTermScoreById(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, termScore2)

	require.Equal(t, arg.ID, termScore2.ID)
	require.Equal(t, arg.Assessment, termScore2.Assessment)
	require.Equal(t, arg.Exam, termScore2.Exam)
	require.Equal(t, arg.UpdatedAt, termScore2.UpdatedAt)
}

func createTestTermScore(t *testing.T) TermScore {
	var ass sql.NullFloat64
	ass.Scan(10.0)
	var exam sql.NullFloat64
	exam.Scan(45.0)

	const demoId = 1
	arg := CreateTermScoreParams{
		Assessment: ass,
		Exam:       exam,
		SubjectID:  demoId,
		TermID:     demoId,
		SessionID:  demoId,
		ClassID:    demoId,
	}

	termScore, err := testQueries.CreateTermScore(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, termScore)

	require.Equal(t, arg.Assessment, termScore.Assessment)
	require.Equal(t, arg.Exam, termScore.Exam)
	require.Equal(t, arg.SubjectID, termScore.SubjectID)
	require.Equal(t, arg.TermID, termScore.TermID)
	require.Equal(t, arg.SessionID, termScore.SessionID)
	require.Equal(t, arg.ClassID, termScore.ClassID)

	require.NotZero(t, termScore.CreatedAt)
	require.NotZero(t, termScore.UpdatedAt)
	return termScore
}

func compareTermScores(t *testing.T, termScore1, termScore2 TermScore) {
	require.Equal(t, termScore1.ID, termScore2.ID)
	require.Equal(t, termScore1.Assessment, termScore2.Assessment)
	require.Equal(t, termScore1.Exam, termScore2.Exam)
	require.Equal(t, termScore1.SessionID, termScore2.SessionID)
	require.Equal(t, termScore1.SubjectID, termScore2.SubjectID)
	require.Equal(t, termScore1.ClassID, termScore2.ClassID)
	require.Equal(t, termScore1.TermID, termScore2.TermID)
	require.Equal(t, termScore1.CreatedAt, termScore2.CreatedAt)
}
