package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Watson-Sei/test-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := pb.NewHelloServiceClient(cc)
	request := &pb.HelloRequest{Name: "Watson Sei"}

	res, err := client.SayHello(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive response => %s ", res.Message)
}
