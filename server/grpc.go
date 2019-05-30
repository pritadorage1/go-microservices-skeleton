package server

import (
	"time"

	"go_microservices_skeleton/middleware"
	usecase "go_microservices_skeleton/usecase"

	"github.com/micro/go-grpc"
	micro "github.com/micro/go-micro"
)

//StartGRPCServer : start grpc server
func StartGRPCServer(impl usecase.MemberServiceHandler) (err error) {

	// Create a new service. Optionally include some options here.
	srv := grpc.NewService(
		micro.Name("sc.srv.user"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		// Middleware for error logging
		micro.WrapHandler(middleware.Logger),
	)

	// Init will parse the command line flags.
	srv.Init()

	// Register handler
	usecase.RegisterMemberServiceMicroHandler(srv.Server(), impl)

	return srv.Run()

}
