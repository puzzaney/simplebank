package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/puzzaney/simplebank/api"
	db "github.com/puzzaney/simplebank/db/sqlc"
	"github.com/puzzaney/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load config file", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Failed to connect to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}
}
