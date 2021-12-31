package main

import (
	"context"
	trippb "coolcar/proto/gen/go"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.Lshortfile)
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("connot connect server: %v", err)
	}

	tsClient := trippb.NewTripServiceClient(conn)

	r, err := tsClient.GetTrip(context.Background(), &trippb.GetTripRequest{
		Id: "trip456",
	})

	if err != nil {
		log.Fatalf("connot call GetTrip: %v", err)
	}

	fmt.Println(r)
}
