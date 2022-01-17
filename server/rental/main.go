package main

import (
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/trip"
	"coolcar/shared/server"
	"log"

	"google.golang.org/grpc"
)

func main() {
	loger, err := server.NewZapLogger()

	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}

	// err = server.RunGRPCServer(&server.GRPCConfig{
	// 	Name:              "rental",
	// 	Addr:              ":8082",
	// 	AuthPublicKeyFile: "shared/auth/pub.key",
	// 	Logger:            loger,
	// 	RegisterFunc: func(s *grpc.Server) {
	// 		rentalpb.RegisterTripServiceServer(s, &trip.Service{
	// 			Logger:                         loger,
	// 			UnimplementedTripServiceServer: rentalpb.UnimplementedTripServiceServer{},
	// 		})
	// 	},
	// })

	// loger.Fatal("cannot server", zap.Error(err))

	loger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name:              "rental",
		Addr:              ":8082",
		AuthPublicKeyFile: "shared/auth/pub.key",
		Logger:            loger,
		RegisterFunc: func(s *grpc.Server) {
			rentalpb.RegisterTripServiceServer(s, &trip.Service{
				Logger:                         loger,
				UnimplementedTripServiceServer: rentalpb.UnimplementedTripServiceServer{},
			})
		},
	}))
}
