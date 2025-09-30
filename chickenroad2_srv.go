package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"os"

	"cocogame-max/chickenroad2_srv/internal/config"
	"cocogame-max/chickenroad2_srv/internal/server"
	"cocogame-max/chickenroad2_srv/internal/svc"
	"cocogame-max/chickenroad2_srv/pb_chickenroad2"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", fmt.Sprintf("etc/config_%s.yaml", os.Getenv("ENV")), "the config file")

func main() {
	logx.DisableStat()
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb_chickenroad2.RegisterChickenRoad2SrvServer(grpcServer, server.NewChickenRoad2SrvServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
