package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/realHoangHai/kratos/bootstrap"
	"github.com/realHoangHai/kratos/common/constant"
	"github.com/realHoangHai/kratos/pkg/utils/trans"
)

var version string

// go build -ldflags "-X main.version=x.y.z"

func newApp(ll log.Logger, rr registry.Registrar, hs *http.Server) *kratos.App {
	return bootstrap.NewApp(ll, rr, hs)
}

func main() {
	bootstrap.Bootstrap(wireApp, trans.Ptr(constant.FrontendService), trans.Ptr(version))
}
