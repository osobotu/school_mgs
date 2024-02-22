package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/osobotu/school_mgs/api"
	db "github.com/osobotu/school_mgs/db/sqlc"
	"github.com/osobotu/school_mgs/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the database")
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server")
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server")
	}
}
