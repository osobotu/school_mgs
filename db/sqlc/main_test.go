package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/osobotu/school_mgs/utils"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {

	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config")
	}

	testDb, err = sql.Open(config.DBDriver, config.DBSource)

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
