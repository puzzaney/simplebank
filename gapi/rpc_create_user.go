package gapi

import (
	"context"

	"github.com/lib/pq"
	db "github.com/puzzaney/simplebank/db/sqlc"
	"github.com/puzzaney/simplebank/pb"
	"github.com/puzzaney/simplebank/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	arg := db.CreateUserParams{
		Username:       req.GetUsername(),
		Email:          req.GetEmail(),
		HashedPassword: hashedPassword,
		FullName:       req.GetFullName(),
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "username already exists: %s", err)
			}
			return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
		}
	}

	res := &pb.CreateUserResponse{
		User: convertUser(user),
	}
	return res, nil
}
