package middleware

import "github.com/zeromicro/go-zero/rest"

func RegisterMiddlewares() []rest.Middleware {
	return []rest.Middleware{
		AuthMiddleware,
	}
}
