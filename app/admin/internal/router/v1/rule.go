package v1

import (
	"github.com/laoningmeng/go-zero-admin/app/admin/internal/handler/rule"
	"github.com/laoningmeng/go-zero-admin/app/admin/internal/middleware"
	"github.com/laoningmeng/go-zero-admin/app/admin/internal/svc"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

func RegisterRuleRouter(server *rest.Server, serverCtx *svc.ServiceContext) {
	accessRuleWithToken(server, serverCtx)
}

func accessRuleWithToken(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			middleware.RegisterMiddlewares(),
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/add",
					Handler: rule.AddHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/list",
					Handler: rule.ListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/del",
					Handler: rule.DelHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: rule.UpdateHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/v1/rule"),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
	)
}
