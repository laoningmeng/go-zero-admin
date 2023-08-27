package main

import (
	"flag"
	"fmt"
	"github.com/laoningmeng/go-zero-admin/app/admin/internal/router"
	"github.com/zeromicro/go-zero/core/stat"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"strconv"

	"github.com/laoningmeng/go-zero-admin/app/admin/internal/config"
	"github.com/laoningmeng/go-zero-admin/app/admin/internal/svc"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/admin-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	stat.DisableLog()

	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(nil, nil))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	router.RegisterRouters(server, ctx)
	_ = consul.RegisterService(c.Host+":"+strconv.Itoa(c.Port), c.Consul)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
