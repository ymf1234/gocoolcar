package main

import (
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/trip"
	"coolcar/shared/auth"
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
	lis, err := net.Listen("tcp", ":8082")

	if err != nil {
		loger.Fatal("cannot listen", zap.Error(err))
	}

	in, err := auth.Interceptor("shared/auth/pub.key")

	if err != nil {
		loger.Fatal("cannot create auth interceptor")
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(in))
	rentalpb.RegisterTripServiceServer(s, &trip.Service{
		Logger:                         loger,
		UnimplementedTripServiceServer: rentalpb.UnimplementedTripServiceServer{},
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
