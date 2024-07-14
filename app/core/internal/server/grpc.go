package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"github.com/realHoangHai/kratos/api/gen/go/conf"
	userv1 "github.com/realHoangHai/kratos/api/gen/go/user/v1"
	"github.com/realHoangHai/kratos/app/core/internal/service"
	"github.com/realHoangHai/kratos/bootstrap/rpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(cfg *conf.Bootstrap, logger log.Logger,
	userSvc *service.UserService,
) *grpc.Server {
	srv := rpc.CreateGrpcServer(cfg, logging.Server(logger))

	userv1.RegisterUserServiceServer(srv, userSvc)
	return srv
}
