package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/osobotu/school_mgs/api"
	db "github.com/osobotu/school_mgs/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres:password@localhost:5432/school_mgs?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to the database")
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server")
	}
}
