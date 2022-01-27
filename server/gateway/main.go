package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/server"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	lg, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create zap logger: %v", err)
	}

	c := context.Background()

	c, cancel := context.WithCancel(c)

	defer cancel()

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard,
		&runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseEnumNumbers: true, // 枚举字段的值使用数字
				UseProtoNames:  true,
				// 传给 clients 的 json key 使用下划线 `_`
				// AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
				// 这里说明应使用 access_token
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true, // 忽略 client 发送的不存在的 poroto 字段
			},
		},
	))

	serverConfig := []struct {
		name         string
		addr         string
		registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
	}{
		{
			name:         "auth",
			addr:         "localhost:8081",
			registerFunc: authpb.RegisterAuthServiceHandlerFromEndpoint,
		},
		{
			name:         "rental",
			addr:         "localhost:8082",
			registerFunc: rentalpb.RegisterTripServiceHandlerFromEndpoint,
		},
	}

	for _, s := range serverConfig {
		// auth
		err := s.registerFunc(
			c,
			mux,
			s.addr, []grpc.DialOption{grpc.WithInsecure()},
		)

		if err != nil {
			lg.Sugar().Fatalf("cannot register %s: %v", s.name, err)
		}
	}

	addr := ":8080"
	lg.Sugar().Infof("grpc gateway started at %s", addr)
	lg.Sugar().Fatal(http.ListenAndServe(addr, mux))

}
