package main

import (
	"Gym-backend/api"
	db "Gym-backend/db/sqlc"
	_ "Gym-backend/docs"
	"Gym-backend/util"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

// @title Gym-backend API
// @version 1.0
// @description Gym-backend API

// @host 0.0.0.0:8080
// @basePath
func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
