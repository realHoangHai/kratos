package server

import (
	"context"
	"github.com/google/wire"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"

	adminv1 "github.com/realHoangHai/kratos/api/gen/go/admin/v1"
	"github.com/realHoangHai/kratos/api/gen/go/conf"
	"github.com/realHoangHai/kratos/app/admin/internal/service"
	"github.com/realHoangHai/kratos/app/client/assets"
	"github.com/realHoangHai/kratos/bootstrap/rpc"
	"github.com/realHoangHai/kratos/bootstrap/swagger"
	authnengine "github.com/realHoangHai/kratos/component/authentication/engine"
	authn "github.com/realHoangHai/kratos/component/authentication/middleware"
	authzengine "github.com/realHoangHai/kratos/component/authorization/engine"
	authz "github.com/realHoangHai/kratos/component/authorization/middleware"
	"github.com/realHoangHai/kratos/handler/middleware/auth"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(
	NewHTTPServer,
)

// newHttpWhiteListMatcher .
func newHttpWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]bool)
	whiteList[adminv1.OperationAuthenticationServiceLogin] = true
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// newHttpMiddleware create middleware.
func newHttpMiddleware(authenticator authnengine.Authenticator, authorizer authzengine.Engine, logger log.Logger) []middleware.Middleware {
	var ms []middleware.Middleware
	ms = append(ms, logging.Server(logger))
	ms = append(ms, selector.Server(
		authn.Server(authenticator),
		auth.Server(),
		authz.Server(authorizer),
	).Match(newHttpWhiteListMatcher()).Build())
	return ms
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	cfg *conf.Bootstrap, logger log.Logger,
	authenticator authnengine.Authenticator, authorizer authzengine.Engine,
	authnSvc *service.AuthenticationService,
	userSvc *service.UserService,
) *http.Server {
	srv := rpc.CreateHttpServer(cfg, newHttpMiddleware(authenticator, authorizer, logger)...)

	adminv1.RegisterAuthenticationServiceHTTPServer(srv, authnSvc)
	adminv1.RegisterUserServiceHTTPServer(srv, userSvc)

	if cfg.GetServer().GetHttp().GetEnableSwagger() {
		swagger.RegisterSwaggerUIServerWithOption(
			srv,
			swagger.WithTitle("Admin Service"),
			swagger.WithMemoryData(assets.OpenApiData, "yaml"),
		)
	}

	return srv
}
