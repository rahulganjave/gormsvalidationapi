package main

import (
	"context"
	"log"

	"github.com/gormsvalidationapi/api"
	db "github.com/gormsvalidationapi/db/sqlc"
	"github.com/gormsvalidationapi/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		//log.Fatal().Err(err).Msg("cannot load config")
		log.Fatal("cannot load config:", err)
	}

	// if config.Environment == "development" {
	// 	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	// }

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(connPool) //DB
	runGinServer(config, store)
	//go runGatewayServer(config, store)
	//runGrpcServer(config, store)
}
func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store) //Gin Server
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
