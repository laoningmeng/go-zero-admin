package v1

import (
	"github.com/laoningmeng/go-zero-admin/app/admin/internal/handler/user"
	"github.com/laoningmeng/go-zero-admin/app/admin/internal/middleware"
	"github.com/laoningmeng/go-zero-admin/app/admin/internal/svc"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

func RegisterRoleRouter(server *rest.Server, serverCtx *svc.ServiceContext) {
	accessRoleWithToken(server, serverCtx)
}

func accessRoleWithToken(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			middleware.RegisterMiddlewares(),
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/add",
					Handler: user.AddHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/list",
					Handler: user.ListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/del",
					Handler: user.DelHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: user.UpdateHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/v1/role"),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
	)
}
