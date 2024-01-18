package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateSession(t *testing.T) {
	createTestSession(t)
}

func TestDeleteSession(t *testing.T) {
	session1 := createTestSession(t)
	err := testQueries.DeleteSession(context.Background(), session1.ID)
	require.NoError(t, err)

	session2, err := testQueries.GetSessionById(context.Background(), session1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, session2)
}

func TestGetSessionById(t *testing.T) {
	session1 := createTestSession(t)
	session2, err := testQueries.GetSessionById(context.Background(), session1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, session2)

	compareSessions(t, session1, session2)
}

func createTestSession(t *testing.T) Session {

	start := time.Date(2023, 1, 9, 0, 0, 0, 0, time.UTC)
	var sd sql.NullTime
	sd.Scan(start)

	var ed sql.NullTime
	ed.Scan(start.AddDate(1, 0, 0))

	arg := CreateSessionParams{
		Session:   "2023/2024",
		StartDate: sd,
		EndDate:   ed,
	}

	session, err := testQueries.CreateSession(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, session)

	require.Equal(t, arg.Session, session.Session)
	require.Equal(t, arg.StartDate, session.StartDate)
	require.Equal(t, arg.EndDate, session.EndDate)

	require.NotZero(t, session.ID)
	require.NotZero(t, session.CreatedAt)
	require.NotZero(t, session.UpdatedAt)

	return session
}

func compareSessions(t *testing.T, session1, session2 Session) {
	require.Equal(t, session1.ID, session2.ID)
	require.Equal(t, session1.Session, session2.Session)
	require.Equal(t, session1.StartDate, session2.StartDate)
	require.Equal(t, session1.EndDate, session2.EndDate)
	require.Equal(t, session1.CreatedAt, session2.CreatedAt)
}
