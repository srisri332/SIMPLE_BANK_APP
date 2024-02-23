package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/srisri332/simplebank/api"
	db "github.com/srisri332/simplebank/db/sqlc"
	"github.com/srisri332/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to DB: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}

}
