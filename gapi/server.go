package gapi

import (
	"fmt"

	db "github.com/srisri332/simplebank/db/sqlc"
	"github.com/srisri332/simplebank/pb"
	"github.com/srisri332/simplebank/token"
	"github.com/srisri332/simplebank/util"
)

// serves gRPC request for banking service
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      *db.Store
	tokenMaker token.Maker
}

// // mustEmbedUnimplementedSimpleBankServer implements pb.SimpleBankServer.
// func (*Server) mustEmbedUnimplementedSimpleBankServer() {
// 	panic("unimplemented")
// }

// // CreateUser implements pb.SimpleBankServer.
// func (*Server) CreateUser(context.Context, *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
// 	panic("unimplemented")
// }

// // LoginUser implements pb.SimpleBankServer.
// func (*Server) LoginUser(context.Context, *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
// 	panic("unimplemented")
// }

// // mustEmbedUnimplementedSimpleBankServer implements pb.SimpleBankServer.
// func (*Server) mustEmbedUnimplementedSimpleBankServer() {
// 	panic("unimplemented")
// }

func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker")
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
