package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Watson-Sei/test-grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.HelloServiceServer
}

func (s *Server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Print("Hello " + req.Name)
	return &pb.HelloResponse{
		Message: "Hello " + req.Name,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()

	pb.RegisterHelloServiceServer(server, &Server{})
	reflection.Register(server)

	server.Serve(lis)

}
