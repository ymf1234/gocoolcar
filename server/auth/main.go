package main

import (
	"log"
	"net"

	"go.uber.org/zap"
)

func main() {
	loger, error = zap.NewDevelopment()

	if error != nil {
		log.Fatalf(error)
	}
	lis, err := net.Listen("tcp", ":8081")

	if err != nil {

	}
}
