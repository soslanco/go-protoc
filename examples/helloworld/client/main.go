package main

import (
	"context"
	"log"
	"time"

	pb "github.com/soslanco/go-protoc/examples/helloworld/helloworld"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	grpcClient := pb.NewHWClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := grpcClient.HelloWorld(ctx, &pb.HelloWorldRequest{Name: "gRPC client"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(r.Message)
}
