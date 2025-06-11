package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/puzzaney/simplebank/api"
	db "github.com/puzzaney/simplebank/db/sqlc"
	"github.com/puzzaney/simplebank/gapi"
	"github.com/puzzaney/simplebank/pb"
	"github.com/puzzaney/simplebank/util"
	"google.golang.org/grpc"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config file")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Failed to connect to db", err)
	}

	store := db.NewStore(conn)
	runGinServer(config, store)
}

// func runGrpcServer(config util.Config, store db.Store){
// 	server, err := gapi.NewServer(config, store)
// 	if err != nil {
// 		log.Fatal("cannot create gRPC server: %w", err)
// 	}
//
// 	grpcServer := grpc.NewServer()
// 	pb.RegisterSimpleBankServer(grpcServer, server)
//
// }


func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	if err = server.Start(config.HTTPServerAddress); err != nil {
		log.Fatal("Cannot start server")
	}

}
