//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/heytom-labs/heytom-go/internal/biz"
	"github.com/heytom-labs/heytom-go/internal/conf"
	"github.com/heytom-labs/heytom-go/internal/data"
	"github.com/heytom-labs/heytom-go/internal/server"
	"github.com/heytom-labs/heytom-go/internal/service"
)

func wireApp(config.Config, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
