package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

const DemoTerm = "Demo Term"

func TestCreateTerm(t *testing.T) {
	term := createTestTerm(t)
	testQueries.RunCleaners(t, &term)
}

func TestDeleteTerm(t *testing.T) {
	term1 := createTestTerm(t)

	err := testQueries.DeleteTerm(context.Background(), term1.ID)
	require.NoError(t, err)

	term2, err := testQueries.GetTermByID(context.Background(), term1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, term2)

	testQueries.RunCleaners(t, &term1, &term2)

}

func TestGetTermByID(t *testing.T) {
	term1 := createTestTerm(t)
	term2, err := testQueries.GetTermByID(context.Background(), term1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, term2)
	compareTerms(t, term1, term2)

	testQueries.RunCleaners(t, &term1, &term2)
}

func createTestTerm(t *testing.T) Term {
	arg := CreateTermParams{
		Name:   DemoTerm,
		Number: 0,
	}

	term, err := testQueries.CreateTerm(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, term)

	require.Equal(t, arg.Name, term.Name)
	require.Equal(t, arg.Number, term.Number)

	require.NotZero(t, term.ID)
	require.NotZero(t, term.CreatedAt)
	require.NotZero(t, term.UpdatedAt)

	return term
}

func compareTerms(t *testing.T, term1, term2 Term) {
	require.Equal(t, term1.ID, term2.ID)
	require.Equal(t, term1.Name, term2.Name)
	require.Equal(t, term1.Number, term2.Number)
}

func (t *Term) Clean() {
	testQueries.DeleteTerm(context.Background(), t.ID)
}
