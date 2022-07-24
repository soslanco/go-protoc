package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	pb "github.com/soslanco/go-protoc/examples/helloworld/helloworld"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type hwserver struct {
	pb.UnimplementedHWServer
}

func (hw *hwserver) HelloWorld(ctx context.Context, req *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	message := "Hello " + req.Name + "!"
	log.Println(message)
	return &pb.HelloWorldResponse{Message: message}, nil
}

func (hw *hwserver) HelloWorldPrefix(ctx context.Context, req *pb.HelloWorldPrefixRequest) (*pb.HelloWorldPrefixResponse, error) {
	message := req.Prefix.Prefix + " " + req.Name + "!"
	log.Println(message)
	return &pb.HelloWorldPrefixResponse{Message: message}, nil
}

func main() {
	var wg sync.WaitGroup

	// Start GRPC server
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	grpcSrv := grpc.NewServer()
	pb.RegisterHWServer(grpcSrv, &hwserver{})
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := grpcSrv.Serve(l); err != nil {
			log.Print(err)
		}
	}()

	// Start HTTP reverse-proxy server
	ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = pb.RegisterHWHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}
	httpSrv := &http.Server{Addr: ":8080", Handler: mux}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := httpSrv.ListenAndServe(); err != nil {
			log.Print(err)
		}
	}()

	// Wait for OS signal
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		log.Print("wait for OS signal")
		<-sig
		log.Print("signal received")

		log.Print("Stopping HTTP reverse-proxy server (GRPC)...")
		cancel()
		log.Print("Stopped HTTP reverse-proxy server (GRPC)")

		log.Print("Stopping HTTP reverse-proxy server (HTTP)...")
		httpSrv.Close()
		log.Print("Stopped HTTP reverse-proxy server (HTTP)")

		log.Print("Stopping GRPC server...")
		grpcSrv.Stop()
		log.Print("Stopped GRPC server")
	}()

	wg.Wait()
}
