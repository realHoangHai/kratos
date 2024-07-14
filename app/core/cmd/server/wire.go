//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"

	"github.com/realHoangHai/kratos/api/gen/go/conf"
	"github.com/realHoangHai/kratos/app/core/internal/data"
	"github.com/realHoangHai/kratos/app/core/internal/server"
	"github.com/realHoangHai/kratos/app/core/internal/service"
)

// wireApp init kratos application.
func wireApp(log.Logger, registry.Registrar, *conf.Bootstrap) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, service.ProviderSet, newApp))
}
