package main

import (
	"flag"
	"os"

	_ "go.uber.org/automaxprocs"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	kratoslog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"github.com/realHoangHai/kratos/bootstrap/logger"
	"github.com/realHoangHai/kratos/bootstrap/registry"
	"github.com/realHoangHai/kratos/bootstrap/tracer"
	"github.com/realHoangHai/kratos/conf"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "configs", "config path, eg: -conf config.yaml")
}

func newApp(logger kratoslog.Logger, gs *grpc.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
		),
	)
}

func main() {
	flag.Parse()

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer func() {
		_ = c.Close()
	}()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	serviceInfo := conf.NewServiceInfo(Name, Version, id)

	// init logger
	ll := logger.NewLoggerProvider(bc.Logger, serviceInfo)

	// init registrar
	reg := registry.NewRegistry(bc.Registry)

	// init tracer
	if err := tracer.NewTracerProvider(bc.Trace, serviceInfo); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(ll, reg, &bc)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err = app.Run(); err != nil {
		panic(err)
	}
}
