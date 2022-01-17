package server

import (
	"coolcar/shared/auth"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GRPCConfig struct {
	Name              string
	Addr              string
	AuthPublicKeyFile string
	RegisterFunc      func(*grpc.Server)
	Logger            *zap.Logger
}

func RunGRPCServer(c *GRPCConfig) error {
	nameFile := zap.String("name", c.Name)
	lis, err := net.Listen("tcp", ":8082")

	if err != nil {
		c.Logger.Fatal("cannot listen", zap.Error(err))
	}

	var opts []grpc.ServerOption
	if c.AuthPublicKeyFile != "" {
		in, err := auth.Interceptor(c.AuthPublicKeyFile)

		if err != nil {
			c.Logger.Fatal("cannot create auth interceptor", nameFile, zap.Error(err))
		}

		opts = append(opts, grpc.UnaryInterceptor(in))
	}

	s := grpc.NewServer(opts...)
	c.RegisterFunc(s)
	// rentalpb.RegisterTripServiceServer(s, &trip.Service{
	// 	Logger:                         loger,
	// 	UnimplementedTripServiceServer: rentalpb.UnimplementedTripServiceServer{},
	// })

	return s.Serve(lis)
}
