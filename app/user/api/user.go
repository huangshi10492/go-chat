package main

import (
	"flag"
	"fmt"

	"go-chat/app/user/api/internal/config"
	"go-chat/app/user/api/internal/handler"
	"go-chat/app/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting app at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
