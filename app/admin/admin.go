package main

import (
	"flag"
	"fmt"
	"github.com/laoningmeng/go-zero-admin/app/admin/internal/router"

	"github.com/laoningmeng/go-zero-admin/app/admin/internal/config"
	"github.com/laoningmeng/go-zero-admin/app/admin/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/admin-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	router.RegisterRouters(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
