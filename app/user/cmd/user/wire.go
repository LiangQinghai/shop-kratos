//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"gitee.com/qinghailiang/shop-kratos/app/user/internal/biz"
	"gitee.com/qinghailiang/shop-kratos/app/user/internal/conf"
	"gitee.com/qinghailiang/shop-kratos/app/user/internal/data"
	"gitee.com/qinghailiang/shop-kratos/app/user/internal/server"
	"gitee.com/qinghailiang/shop-kratos/app/user/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
