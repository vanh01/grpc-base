package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/vanh01/grpc-base/protos/b"
	basemsg "github.com/vanh01/grpc-base/protos/basemsg"
)

type BServer struct {
	pb.UnimplementedBServer
}

func NewBService() BServer {
	return BServer{}
}

func (b BServer) SendMessage(c context.Context, r *basemsg.GEventMessage) (*basemsg.GEventResult, error) {
	return nil, nil
}

func (b *BServer) Run() {
	address := fmt.Sprintf("localhost:8081")
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterBServer(grpcServer, b)

	log.Printf("B service is listening on %s\n", address)
	log.Fatalln(grpcServer.Serve(lis))
}
