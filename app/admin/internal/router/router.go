package router

import (
	v1 "github.com/laoningmeng/go-zero-admin/app/admin/internal/router/v1"
	"github.com/laoningmeng/go-zero-admin/app/admin/internal/svc"
	"github.com/zeromicro/go-zero/rest"
)

func RegisterRouters(server *rest.Server, serverCtx *svc.ServiceContext) {
	v1.RegisterUserRouter(server, serverCtx)
	v1.RegisterRoleRouter(server, serverCtx)
	v1.RegisterRuleRouter(server, serverCtx)
}
