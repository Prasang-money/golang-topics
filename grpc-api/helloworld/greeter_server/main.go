package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/prasang/grpc/helloworld"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) Hello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	name := request.Name

	response := &pb.HelloResponse{
		Greeting: "Hello " + name,
	}
	return response, nil
}
func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("server is listening on %v ...", port)
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serevr %v", err)
	}
}
