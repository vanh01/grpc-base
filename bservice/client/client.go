package client

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	grpcbase "github.com/vanh01/grpc-base"
	pb "github.com/vanh01/grpc-base/protos/a"
	"github.com/vanh01/grpc-base/protos/basemsg"
)

var GrpcClientInstance *grpcClient = newGrpcClient()

type grpcClient struct {
	AClient pb.AClient
}

func newGrpcClient() *grpcClient {
	return &grpcClient{
		AClient: newGrpcAClient(),
	}
}

func newGrpcAClient() pb.AClient {
	target := fmt.Sprintf("%s:%d", "localhost", 8082)
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.NewClient(target, opts...)
	if err != nil {
		log.Printf("failed to connect to the server: %v", err)
		return nil
	}
	if conn == nil {
		return nil
	}
	log.Println("successful to connect to the server")
	return pb.NewAClient(conn)
}

func SendAMessage[T, K any](request T) (K, error) {
	r := grpcbase.GetRequest[T, K](request)
	result, err := GrpcClientInstance.AClient.SendMessage(context.Background(), &basemsg.GEventMessage{
		Data:           r.Data,
		DataType:       r.DataType,
		ReturnDataType: r.ReturnDataType,
		ConsumerType:   r.ConsumerType,
	})

	var data K
	if err != nil {
		return data, err
	}

	err = json.Unmarshal([]byte(result.Data), &data)

	return data, err
}
