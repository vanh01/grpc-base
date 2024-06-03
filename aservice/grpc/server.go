package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	grpcbase "github.com/vanh01/grpc-base"
	pb "github.com/vanh01/grpc-base/protos/a"
	"github.com/vanh01/grpc-base/protos/basemsg"
)

type AServer struct {
	pb.UnimplementedAServer
}

func NewAService() AServer {
	return AServer{}
}

func (a AServer) SendMessage(c context.Context, r *basemsg.GEventMessage) (*basemsg.GEventResult, error) {
	res := grpcbase.SendMessage(grpcbase.GEventMessage{
		DataType:       r.DataType,
		Data:           r.Data,
		ReturnDataType: r.ReturnDataType,
		ConsumerType:   r.ConsumerType,
	})
	return &basemsg.GEventResult{
		Data:     res.Data,
		DataType: res.DataType,
	}, nil
}

func (a *AServer) Run() {
	address := fmt.Sprintf("localhost:8082")
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterAServer(grpcServer, a)

	log.Printf("A service is listening on %s\n", address)
	log.Fatalln(grpcServer.Serve(lis))
}
