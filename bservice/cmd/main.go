package main

import (
	"fmt"
	"log"

	"github.com/vanh01/grpc-base/bservice/client"
	"github.com/vanh01/grpc-base/domain"
)

func main() {
	// grpcServer := grpc.NewBService()
	// go grpcServer.Run()
	request := domain.A1Request{
		Id:   1,
		Name: "23",
	}
	fmt.Println("Request: ", request)

	result, err := client.SendAMessage[domain.A1Request, domain.A1Result](request)

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(result)

	// Test request 2
	request2 := domain.A2Request{
		Id:   1,
		Name: "23",
	}
	fmt.Println("Request: ", request)

	result2, err := client.SendAMessage[domain.A2Request, domain.A2Result](request2)

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(result2)
}
