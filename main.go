package Gym_backend

import (
	"database/sql"
	"github.com/JaswanthKarangula/Gym-backend/api"
	db "github.com/JaswanthKarangula/Gym-backend/db/sqlc"
	_ "github.com/JaswanthKarangula/Gym-backend/docs"
	"github.com/JaswanthKarangula/Gym-backend/util"
	_ "github.com/lib/pq"
	"log"
)

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
