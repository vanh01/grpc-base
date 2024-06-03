package main

import (
	"github.com/vanh01/grpc-base/aservice/consumer"
	"github.com/vanh01/grpc-base/aservice/grpc"
)

func main() {
	consumer.Init()
	grpcServer := grpc.NewAService()
	grpcServer.Run()
}
