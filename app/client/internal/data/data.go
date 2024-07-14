package data

import (
	"context"
	"github.com/google/wire"

	"github.com/redis/go-redis/v9"

	"github.com/go-kratos/kratos/v2/log"
	kratosregistry "github.com/go-kratos/kratos/v2/registry"

	"github.com/realHoangHai/kratos/api/gen/go/conf"
	userv1 "github.com/realHoangHai/kratos/api/gen/go/user/v1"
	rdb "github.com/realHoangHai/kratos/bootstrap/cache/redis"
	"github.com/realHoangHai/kratos/bootstrap/registry"
	"github.com/realHoangHai/kratos/bootstrap/rpc"
	"github.com/realHoangHai/kratos/common/constant"
	authn "github.com/realHoangHai/kratos/component/authentication/engine"
	"github.com/realHoangHai/kratos/component/authentication/engine/jwt"
	authz "github.com/realHoangHai/kratos/component/authorization/engine"
	"github.com/realHoangHai/kratos/component/authorization/engine/noop"
	"github.com/realHoangHai/kratos/handler/cache"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,

	NewRedisClient,
	NewDiscovery,

	NewAuthenticator,
	NewAuthorizer,

	NewUserTokenRepo,

	NewUserServiceClient,
)

// Data .
type Data struct {
	log *log.Helper
	rdb *redis.Client

	authenticator authn.Authenticator
	authorizer    authz.Engine
}

// NewData .
func NewData(rdb *redis.Client,
	authenticator authn.Authenticator,
	authorizer authz.Engine,
	logger log.Logger,
) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "client/data"))

	d := &Data{
		rdb:           rdb,
		log:           l,
		authenticator: authenticator,
		authorizer:    authorizer,
	}

	return d, func() {
		l.Info("message", "Closing the data resources")
		if err := d.rdb.Close(); err != nil {
			l.Error(err)
		}
	}, nil
}

// NewRedisClient create a redis client
func NewRedisClient(cfg *conf.Bootstrap, _ log.Logger) *redis.Client {
	return rdb.NewClient(cfg.Data)
}

// NewDiscovery create a service discovery client
func NewDiscovery(cfg *conf.Bootstrap) kratosregistry.Discovery {
	return registry.NewDiscovery(cfg.Registry)
}

// NewAuthenticator create authenticator
func NewAuthenticator(cfg *conf.Bootstrap) authn.Authenticator {
	authenticator, _ := jwt.NewAuthenticator(
		jwt.WithKey([]byte(cfg.Server.Http.Middleware.Auth.Key)),
		jwt.WithSigningMethod(cfg.Server.Http.Middleware.Auth.Method),
	)
	return authenticator
}

// NewAuthorizer create authorizer
func NewAuthorizer() authz.Engine {
	return noop.State{}
}

func NewUserTokenRepo(data *Data, authenticator authn.Authenticator, logger log.Logger) *cache.UserToken {
	const (
		userAccessTokenKeyPrefix  = "uat_"
		userRefreshTokenKeyPrefix = "urt_"
	)
	return cache.NewUserToken(data.rdb, authenticator, logger, userAccessTokenKeyPrefix, userRefreshTokenKeyPrefix)
}

func NewUserServiceClient(r kratosregistry.Discovery, c *conf.Bootstrap) userv1.UserServiceClient {
	return userv1.NewUserServiceClient(rpc.CreateGrpcClient(context.Background(), r, constant.CoreService, c))
}
