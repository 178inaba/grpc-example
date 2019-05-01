package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/178inaba/grpc-example/echo/proto"
	"github.com/178inaba/grpc-example/interceptor"
	"google.golang.org/grpc"
)

func main() {
	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v.", err)
	}

	srv := grpc.NewServer(grpc.UnaryInterceptor(interceptor.LoggingUnaryForServer))
	pb.RegisterEchoServiceServer(srv, &echoService{})
	log.Printf("Start server on port: %d.", port)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v.", err)
	}
}
