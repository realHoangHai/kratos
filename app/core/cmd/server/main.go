package main

import (
	"github.com/go-kratos/kratos/v2"
	kratoslog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"github.com/realHoangHai/kratos/bootstrap"
	"github.com/realHoangHai/kratos/common/constant"
	"github.com/realHoangHai/kratos/pkg/utils/trans"
)

var version string

// go build -ldflags "-X main.version=x.y.z"

func newApp(ll kratoslog.Logger, rr registry.Registrar, gs *grpc.Server) *kratos.App {
	return bootstrap.NewApp(ll, rr, gs)
}

func main() {
	bootstrap.Bootstrap(wireApp, trans.Ptr(constant.CoreService), trans.Ptr(version))
}
