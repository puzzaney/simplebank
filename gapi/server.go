package gapi

import (
	"fmt"

	db "github.com/puzzaney/simplebank/db/sqlc"
	"github.com/puzzaney/simplebank/pb"
	"github.com/puzzaney/simplebank/token"
	"github.com/puzzaney/simplebank/util"
	"github.com/puzzaney/simplebank/worker"
)

// Serves gRPC requests for our banking servers
type Server struct {
	pb.UnimplementedSimpleBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

// Creates new gRPC server
func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.SecretKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil

}
