package bootstrap

import (
	"fmt"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	kratosregistry "github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport"

	"github.com/realHoangHai/kratos/api/gen/go/conf"
	"github.com/realHoangHai/kratos/bootstrap/config"
	"github.com/realHoangHai/kratos/bootstrap/logger"
	"github.com/realHoangHai/kratos/bootstrap/registry"
	"github.com/realHoangHai/kratos/bootstrap/tracer"
)

var (
	Service = config.NewServiceInfo(
		"",
		"1.0.0",
		"",
	)
)

// NewApp create application
func NewApp(ll log.Logger, rr kratosregistry.Registrar, srv ...transport.Server) *kratos.App {
	return kratos.New(
		kratos.ID(Service.GetInstanceId()),
		kratos.Name(Service.Name),
		kratos.Version(Service.Version),
		kratos.Metadata(Service.Metadata),
		kratos.Logger(ll),
		kratos.Server(
			srv...,
		),
		kratos.Registrar(rr),
	)
}

// DoBootstrap execute boot
func DoBootstrap(serviceInfo *config.ServiceInfo) (*conf.Bootstrap, log.Logger, kratosregistry.Registrar) {
	// inject command flags
	Flags := config.NewCommandFlags()
	Flags.Init()

	var err error

	// load configs
	if err = config.LoadBootstrapConfig(Flags.Conf); err != nil {
		panic(fmt.Sprintf("load config failed: %v", err))
	}

	// init logger
	ll := logger.NewLoggerProvider(config.GetBootstrapConfig().Logger, serviceInfo)

	// init registrar
	reg := registry.NewRegistry(config.GetBootstrapConfig().Registry)

	// init tracer
	if err = tracer.NewTracerProvider(config.GetBootstrapConfig().Trace, serviceInfo); err != nil {
		panic(fmt.Sprintf("init tracer failed: %v", err))
	}

	return config.GetBootstrapConfig(), ll, reg
}

type InitApp func(logger log.Logger, registrar kratosregistry.Registrar, bootstrap *conf.Bootstrap) (*kratos.App, func(), error)

// Bootstrap application boot startup
func Bootstrap(initApp InitApp, serviceName, version *string) {
	if serviceName != nil && len(*serviceName) != 0 {
		Service.Name = *serviceName
	}
	if version != nil && len(*version) != 0 {
		Service.Version = *version
	}

	// bootstrap
	cfg, ll, reg := DoBootstrap(Service)

	// init app
	app, cleanup, err := initApp(ll, reg, cfg)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// run the app.
	if err = app.Run(); err != nil {
		panic(err)
	}
}
