// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/laoningmeng/go-zero-admin/common/logger"
	"github.com/laoningmeng/go-zero-admin/services/admin/internal/config"
	"github.com/laoningmeng/go-zero-admin/services/admin/internal/logic"
	"github.com/laoningmeng/go-zero-admin/services/admin/internal/model"
	"github.com/laoningmeng/go-zero-admin/services/admin/internal/server"
	"github.com/laoningmeng/go-zero-admin/services/admin/internal/svc"
	"github.com/zeromicro/go-zero/zrpc"
)

import (
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

// Injectors from wire.go:

func zeroApp(c config.Config, ctx *svc.ServiceContext) *zrpc.RpcServer {
	nacosConf := model.NewNacosConf(c)
	db := model.NewDB(nacosConf, ctx)
	loggerLogger := logger.NewZapLogger()
	userRepo := model.NewUserModel(db, loggerLogger)
	userLogic := logic.NewUserLogic(userRepo, loggerLogger)
	roleRepo := model.NewRoleModel(db, loggerLogger)
	roleLogic := logic.NewRoleLogic(roleRepo, loggerLogger)
	ruleRepo := model.NewRuleModel(db, loggerLogger)
	ruleLogic := logic.NewRuleLogic(ruleRepo, loggerLogger)
	adminServer := server.NewAdminServer(ctx, userLogic, roleLogic, ruleLogic)
	rpcServer := newRpcServer(c, adminServer)
	return rpcServer
}
