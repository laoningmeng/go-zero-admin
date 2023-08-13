//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/laoningmeng/go-zero-admin/services/admin/internal/config"
	"github.com/laoningmeng/go-zero-admin/services/admin/internal/logic"
	"github.com/laoningmeng/go-zero-admin/services/admin/internal/model"
	"github.com/laoningmeng/go-zero-admin/services/admin/internal/server"
	"github.com/laoningmeng/go-zero-admin/services/admin/internal/svc"
	"github.com/zeromicro/go-zero/zrpc"
)

func zeroApp(c config.Config, ctx *svc.ServiceContext) *zrpc.RpcServer {
	panic(wire.Build(logic.ProviderSet, server.ProviderSet, model.ProviderSet, newRpcServer))
}
