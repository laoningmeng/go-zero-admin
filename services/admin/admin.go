package main

import (
	"flag"
	"github.com/fatih/color"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/laoningmeng/go-zero-admin/services/admin/admin"
	"github.com/laoningmeng/go-zero-admin/services/admin/internal/config"
	"github.com/laoningmeng/go-zero-admin/services/admin/internal/interceptor"
	"github.com/laoningmeng/go-zero-admin/services/admin/internal/server"
	"github.com/laoningmeng/go-zero-admin/services/admin/internal/svc"
)

var configFile = flag.String("f", "etc/admin.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	app := zeroApp(c, ctx)
	defer app.Stop()
	_ = consul.RegisterService(c.ListenOn, c.Consul)
	color.Blue("Server Port: %s \n", c.ListenOn)
	color.Blue("Server Name: %s \n", c.Name)
	app.Start()
}

func newRpcServer(c config.Config, user *server.AdminServer) *zrpc.RpcServer {
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		admin.RegisterAdminServer(grpcServer, user)
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	s.AddUnaryInterceptors(interceptor.Intercept)
	return s
}
