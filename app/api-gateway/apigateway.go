package main

import (
	"flag"
	"os"
	"path/filepath"

	"go-zero-boilerplate/app/api-gateway/internal/config"
	"go-zero-boilerplate/app/api-gateway/internal/svc"
	"go-zero-boilerplate/pkg/zero-contrib/errx"
	gwx "go-zero-boilerplate/pkg/zero-contrib/gatewayx"
	"go-zero-boilerplate/pkg/zero-contrib/middleware"

	"github.com/zeromicro/go-zero/gateway"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/joho/godotenv"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
)

var configFile = flag.String("f", "etc/api-gateway.yaml", "the config file")

func init() {
	dir, err := os.Getwd()
	if err != nil {
		logx.Error(err)
	}
	envPath := filepath.Join(dir, ".env")
	logx.Infof(".env file=%s", envPath)
	err = godotenv.Load(envPath)
	if err != nil {
		logx.Info("loading .env file fail")
	}
}

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())
	ctx := svc.NewServiceContext(c)

	// s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
	// 	pb.RegisterIdentityServer(grpcServer, identity.NewIdentityServer(ctx))

	// 	if c.Mode == service.DevMode || c.Mode == service.TestMode {
	// 		reflection.Register(grpcServer)
	// 	}
	// })
	
	// s.AddUnaryInterceptors(interceptorx.MetadataInterceptor)
	
	gw := gateway.MustNewServer(c.Gateway)
	mapper := gwx.MustNewRouter(c.Gateway.Upstreams)
	gw.Use(middleware.Auth(ctx.JwtManager, mapper))
	httpx.SetErrorHandler(errx.NewErrorHandler())

	group := service.NewServiceGroup()
	// group.Add(s)
	group.Add(gw)

	defer group.Stop()

	// logx.Infof("Starting rpc server at %s...\n", c.ListenOn)
	logx.Infof("Starting gateway at %s:%d...\n", c.Gateway.Host, c.Gateway.Port)
	group.Start()
}