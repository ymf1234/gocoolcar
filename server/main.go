package main

import (
	"context"
	trippb "coolcar/proto/gen/go"
	trip "coolcar/tripservice"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	log.SetFlags(log.Lshortfile)
	go startGRPCGateway()
	lis, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Printf("failed to listen:%v", err)
	}

	s := grpc.NewServer()
	trippb.RegisterTripServiceServer(s, &trip.Service{})

	log.Fatal(s.Serve(lis))
}

func startGRPCGateway() {
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

	err := trippb.RegisterTripServiceHandlerFromEndpoint(
		c,
		mux,
		"127.0.0.1:8081", []grpc.DialOption{grpc.WithInsecure()},
	)

	if err != nil {
		log.Fatalf("connot start grpc gateway: %v", err)
	}

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("connot listen and server: %v", err)
	}
}
