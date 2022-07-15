package main

import (
	"context"
	"log"
	"net"
	"testing"

	pb "github.com/prasang/grpc/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

var lis *bufconn.Listener

const bufSize = 1024 * 1024

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestHelloGrpcService(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)
	name := "Prasang"
	request := &pb.HelloRequest{
		Name: "something",
	}

	resp, err := client.Hello(ctx, request)
	if err != nil {
		t.Fatalf("Hello service failed: %v", err)
	}
	if resp.Greeting != "Hello "+name {
		t.Fatalf("received wrong response: %+v", resp)
	}
	log.Printf("Recived response: %+v", resp)
}
