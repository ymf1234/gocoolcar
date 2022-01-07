package main

import (
	"context"
	"coolcar/auth/api/auth"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/dao"
	"coolcar/auth/token"
	"coolcar/auth/wechat"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	c := context.Background()
	mongoClient, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://114.55.95.101:27017/"))
	if err != nil {
		loger.Fatal("connot connect mongodb", zap.Error(err))
	}

	pkFile, err := os.Open("auth/priv.key")

	if err != nil {
		loger.Fatal("cannot open private key", zap.Error(err))
	}
	pkBytes, err := ioutil.ReadAll(pkFile)
	if err != nil {
		loger.Fatal("cannot read private key", zap.Error(err))
	}

	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(pkBytes)
	if err != nil {
		loger.Fatal("cannot parse private key", zap.Error(err))
	}

	openIDResolver := &wechat.Service{
		AppID:     "wx282842f436b0ceee",
		AppSecret: "248f23a177d11e8cf009c0dfbd511c01",
	}

	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &auth.Service{
		UnimplementedAuthServiceServer: authpb.UnimplementedAuthServiceServer{},
		Logger:                         loger,
		Mongo:                          dao.NewMongo(mongoClient.Database("coolcar")),
		OpenIDResolver:                 openIDResolver,
		TokenExpire:                    2 * time.Hour,
		TokenGenerator:                 token.NewJWTTokenGen("coolcar/auth", privKey),
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
