package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/prasang/grpc-server/helloworld"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("hello grpc client")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := pb.NewHelloServiceClient(cc)
	request := &pb.HelloRequest{
		Name: "Prasang Mani Manav",
	}

	resp, _ := client.Hello(context.Background(), request)

	fmt.Println(resp)
}
