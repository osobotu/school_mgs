package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:password@localhost:5432/school_mgs?sslmode=disable"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDb, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to the database")
	}

	testQueries = New(testDb)

	os.Exit(m.Run())
}

type Cleaner interface {
	Clean()
}

func (tq *Queries) RunCleaners(t *testing.T, models ...Cleaner) {
	for _, v := range models {
		t.Cleanup(func() {
			v.Clean()
		})
	}
}
