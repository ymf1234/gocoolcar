package main

import (
	"coolcar/auth/api/auth"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/wechat"
	"log"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	loger, err := newZapLogger()

	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}
	lis, err := net.Listen("tcp", ":8081")

	if err != nil {
		loger.Fatal("cannot listen", zap.Error(err))
	}

	openIDResolver := &wechat.Service{
		AppID:     "wx282842f436b0ceee",
		AppSecret: "248f23a177d11e8cf009c0dfbd511c01",
	}

	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &auth.Service{
		UnimplementedAuthServiceServer: authpb.UnimplementedAuthServiceServer{},
		Logger:                         loger,
		OpenIDResolver:                 openIDResolver,
	})

	err = s.Serve(lis)
	loger.Fatal("cannot server", zap.Error(err))
}

// 自定义日志
func newZapLogger() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.TimeKey = ""
	return cfg.Build()
}
