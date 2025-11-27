package app

import (
	"net"

	"saitama/internal/handler"
	"saitama/internal/repository"
	"saitama/internal/service"
	"saitama/pkg/db"

	"github.com/yudhiana/one-for-all/api/gen/user"
	"google.golang.org/grpc"
)

type Server struct {
	addr string
}

func NewServer(addr string) *Server {
	return &Server{
		addr: addr,
	}
}

func (s *Server) Run() error {
	listen, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	db := db.GetDatabaseConnection()
	repo := repository.NewUserRepository(db)
	userService := service.NewUserService(repo)

	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, handler.NewUserHandler(userService))
	return grpcServer.Serve(listen)
}
