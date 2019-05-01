package main

import (
	"fmt"
	"log"
	"net"

	"github.com/178inaba/grpc-example/interceptor"
	pb "github.com/178inaba/grpc-example/uploader/proto"
	"google.golang.org/grpc"
)

func main() {
	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v.", err)
	}

	srv := grpc.NewServer(grpc.UnaryInterceptor(interceptor.LoggingUnaryForServer))
	pb.RegisterFileServiceServer(srv, &fileService{})
	log.Printf("Start server on port: %d.", port)
	if err := srv.Serve(lis); err != nil {
		log.Printf("Failed to serve: %v.", err)
	}
}
