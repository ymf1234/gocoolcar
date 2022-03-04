package main

import (
	"context"
	"coolcar/blob/blob"
	"coolcar/blob/dao"
	"coolcar/shared/server"
	"log"

	blobpb "coolcar/blob/api/gen/v1"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, err := server.NewZapLogger()

	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}

	c := context.Background()

	mongoClient, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://114.55.95.101:27017/"))

	if err != nil {
		logger.Fatal("connot connect mongodb", zap.Error(err))
	}

	db := mongoClient.Database("coolcar")

	logger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name:              "blob",
		Addr:              ":8083",
		AuthPublicKeyFile: "shared/auth/pub.key",
		Logger:            logger,
		RegisterFunc: func(s *grpc.Server) {
			blobpb.RegisterBlobServiceServer(s, &blob.Service{
				Mongo:                          dao.NewMongo(db),
				Logger:                         logger,
				UnimplementedBlobServiceServer: blobpb.UnimplementedBlobServiceServer{},
			})
		},
	}))

}
