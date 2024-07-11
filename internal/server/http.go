package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/transport/http"
	userv1 "github.com/realHoangHai/kratos/api/gen/go/user/v1"
	"github.com/realHoangHai/kratos/internal/conf"
	"github.com/realHoangHai/kratos/internal/service"
	"github.com/realHoangHai/kratos/pkg/rpc"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(cfg *conf.Bootstrap, logger log.Logger,
	userSvc *service.UserService,
) *http.Server {
	srv := rpc.CreateHttpServer(cfg, logging.Server(logger))

	userv1.RegisterUserServiceHTTPServer(srv, userSvc)
	return srv
}
