package svc

import (
	"github.com/google/wire"
	"github.com/laoningmeng/go-zero-admin/services/admin/internal/config"
)

var ProviderSet = wire.NewSet(NewServiceContext)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
