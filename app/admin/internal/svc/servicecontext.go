package svc

import (
	"github.com/laoningmeng/go-zero-admin/app/admin/internal/config"
	"github.com/laoningmeng/go-zero-admin/services/admin/admin"
	"github.com/laoningmeng/go-zero-admin/services/admin/adminclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc admin.AdminClient
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config:  c,
		UserRpc: adminclient.NewAdmin(zrpc.MustNewClient(c.AdminConf)),
	}
}
