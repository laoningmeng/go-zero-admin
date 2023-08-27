package v1

import (
	"github.com/laoningmeng/go-zero-admin/app/admin/internal/handler/user"
	"github.com/zeromicro/go-zero/rest"
	"net/http"

	"github.com/laoningmeng/go-zero-admin/app/admin/internal/handler"
	"github.com/laoningmeng/go-zero-admin/app/admin/internal/middleware"
	"github.com/laoningmeng/go-zero-admin/app/admin/internal/svc"
)

func RegisterUserRouter(server *rest.Server, serverCtx *svc.ServiceContext) {
	accessUserWithToken(server, serverCtx)
	whiteUserList(server, serverCtx)
}

func accessUserWithToken(server *rest.Server, serverCtx *svc.ServiceContext) {
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
					Method:  http.MethodGet,
					Path:    "/info",
					Handler: user.UserInfoHandler(serverCtx),
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
		rest.WithPrefix("/v1/user"),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
	)
}

func whiteUserList(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: handler.LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)
}
