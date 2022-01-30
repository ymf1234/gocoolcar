package main

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/trip"
	"coolcar/rental/trip/client/car"
	"coolcar/rental/trip/client/poi"
	"coolcar/rental/trip/client/profile"
	"coolcar/rental/trip/dao"
	"coolcar/shared/server"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
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

	c := context.Background()
	mongoClient, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://114.55.95.101:27017/"))
	if err != nil {
		loger.Fatal("connot connect mongodb", zap.Error(err))
	}

	loger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name:              "rental",
		Addr:              ":8082",
		AuthPublicKeyFile: "shared/auth/pub.key",
		Logger:            loger,
		RegisterFunc: func(s *grpc.Server) {
			rentalpb.RegisterTripServiceServer(s, &trip.Service{
				CarManager:                     &car.Manager{},
				ProfileManager:                 &profile.Manager{},
				POIManager:                     &poi.Manager{},
				Mongo:                          dao.NewMongo(mongoClient.Database("coolcar")),
				Logger:                         loger,
				UnimplementedTripServiceServer: rentalpb.UnimplementedTripServiceServer{},
			})
		},
	}))
}
